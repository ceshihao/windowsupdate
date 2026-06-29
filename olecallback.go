//go:build windows

/*
Copyright 2022 Zheng Dayu
Copyright 2026 Zheng Dayu
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package windowsupdate

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// The asynchronous WUA methods (BeginSearch, BeginDownload, BeginInstall) require
// a non-NULL callback argument. Passing NULL (VT_NULL) makes them fail with
// DISP_E_TYPEMISMATCH (0x80020005).
//
// IMPORTANT: those parameters are NOT IDispatch. They are custom IUnknown-derived
// interfaces (ISearchCompletedCallback, IDownloadProgressChangedCallback,
// IDownloadCompletedCallback, IInstallationProgressChangedCallback,
// IInstallationCompletedCallback). Each declares exactly one method,
// Invoke(IXxxJob*, IXxxCallbackArgs*), located at vtable slot 3 (right after
// IUnknown's QueryInterface/AddRef/Release). They all share the same layout, so
// one 4-entry vtable serves as a universal no-op callback.
//
// A previous version built an IDispatch vtable (7 entries). When WUA invoked the
// completion callback it called slot 3 — which in an IDispatch layout is
// GetTypeInfoCount(this, pctinfo) — passing the job pointer as pctinfo. The body
// wrote 0 through that pointer, corrupting the job's vtable pointer and crashing
// the process (the service then restarts).
//
// AGILITY: WUA runs the asynchronous operation on its own worker thread and
// invokes the callback from there. Our STA session lives on a different COM
// apartment, so COM must marshal the callback across apartments. A raw Go vtable
// object has no marshaler, which made BeginXxx fail with DISP_E_EXCEPTION
// ("Une exception s'est produite"). To fix that we aggregate the COM
// Free-Threaded Marshaler (FTM): QueryInterface(IID_IMarshal) is delegated to the
// FTM, which makes the object agile (callable directly from any apartment, no
// proxy). BeginXxx then succeeds and progress is read by polling IXxxJob.GetProgress().
// The handler signatures are 100% uintptr because that is required by
// syscall.NewCallback.

type noopCallbackVtbl struct {
	pQueryInterface uintptr
	pAddRef         uintptr
	pRelease        uintptr
	pInvoke         uintptr // slot 3: Invoke for every WUA *Completed/*ProgressChanged callback
}

// noopCallback : lpVtbl MUST be the first field (the COM interface pointer points
// to it).
type noopCallback struct {
	lpVtbl *noopCallbackVtbl
	ref    int32
	ftm    uintptr // IUnknown* of the aggregated free-threaded marshaler (0 if unavailable)
}

// HRESULT values as uintptr (only the low 32 bits are significant).
const (
	hrSOK          = uintptr(0x00000000)
	hrEPointer     = uintptr(0x80004003)
	hrENoInterface = uintptr(0x80004002)
)

// WUA callback interface IIDs and IID_IMarshal. The callback IIDs are a
// belt-and-suspenders allowlist: even if WUA only ever queries IUnknown (the
// declared parameter type) and uses the pointer directly, accepting the specific
// IIDs is harmless because they all share our vtable layout.
var (
	iidIMarshal          = ole.NewGUID("{00000003-0000-0000-C000-000000000046}")
	iidSearchCompleted   = ole.NewGUID("{88AEE058-D4B0-4725-A2F1-814A67AE964C}")
	iidDownloadProgress  = ole.NewGUID("{8C3F1CDD-6173-4591-AEBD-A56A53CA77C1}")
	iidDownloadCompleted = ole.NewGUID("{77254866-9F5B-4C8E-B9E2-C77A8530D64B}")
	iidInstallProgress   = ole.NewGUID("{E01402D5-F8DA-43BA-A012-38894BD048F1}")
	iidInstallCompleted  = ole.NewGUID("{45F4F6F3-D602-4F98-9A8A-3EFA152AD2D3}")
)

func ncQueryInterface(this, iid, ppvObject uintptr) uintptr {
	if ppvObject == 0 {
		return hrEPointer
	}
	out := (*uintptr)(unsafe.Pointer(ppvObject))
	if iid == 0 {
		*out = 0
		return hrENoInterface
	}
	guid := (*ole.GUID)(unsafe.Pointer(iid))
	p := (*noopCallback)(unsafe.Pointer(this))

	// Delegate IMarshal to the free-threaded marshaler so the object is agile and
	// WUA can invoke it from its own apartment without a (broken) proxy.
	if p.ftm != 0 && ole.IsEqualGUID(guid, iidIMarshal) {
		return comQueryInterface(p.ftm, iidIMarshal, out)
	}

	if ole.IsEqualGUID(guid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(guid, iidSearchCompleted) ||
		ole.IsEqualGUID(guid, iidDownloadProgress) ||
		ole.IsEqualGUID(guid, iidDownloadCompleted) ||
		ole.IsEqualGUID(guid, iidInstallProgress) ||
		ole.IsEqualGUID(guid, iidInstallCompleted) {
		atomic.AddInt32(&p.ref, 1)
		*out = this
		return hrSOK
	}

	*out = 0
	return hrENoInterface
}

func ncAddRef(this uintptr) uintptr {
	p := (*noopCallback)(unsafe.Pointer(this))
	return uintptr(uint32(atomic.AddInt32(&p.ref, 1)))
}

func ncRelease(this uintptr) uintptr {
	p := (*noopCallback)(unsafe.Pointer(this))
	return uintptr(uint32(atomic.AddInt32(&p.ref, -1)))
}

// ncInvoke is the slot-3 method for every WUA callback interface, e.g.
// ISearchCompletedCallback::Invoke(ISearchJob*, ISearchCompletedCallbackArgs*).
// We ignore the arguments and return S_OK; completion is detected through the
// blocking EndXxx methods and progress through IXxxJob.GetProgress().
func ncInvoke(this, job, args uintptr) uintptr {
	return hrSOK
}

// comQueryInterface calls IUnknown::QueryInterface (vtable slot 0) on a raw COM
// object pointer, used to fetch IMarshal from the aggregated FTM.
func comQueryInterface(unk uintptr, iid *ole.GUID, out *uintptr) uintptr {
	vtbl := *(*uintptr)(unsafe.Pointer(unk)) // first field is the vtable pointer
	pQI := *(*uintptr)(unsafe.Pointer(vtbl)) // slot 0 = QueryInterface
	ret, _, _ := syscall.SyscallN(pQI, unk, uintptr(unsafe.Pointer(iid)), uintptr(unsafe.Pointer(out)))
	return ret
}

var (
	noopVtbl     *noopCallbackVtbl
	noopOnce     sync.Once
	globalNoopCb *noopCallback
	globalNoopMu sync.Once

	modole32                          = syscall.NewLazyDLL("ole32.dll")
	procCoCreateFreeThreadedMarshaler = modole32.NewProc("CoCreateFreeThreadedMarshaler")
)

func getNoopVtbl() *noopCallbackVtbl {
	noopOnce.Do(func() {
		noopVtbl = &noopCallbackVtbl{
			pQueryInterface: syscall.NewCallback(ncQueryInterface),
			pAddRef:         syscall.NewCallback(ncAddRef),
			pRelease:        syscall.NewCallback(ncRelease),
			pInvoke:         syscall.NewCallback(ncInvoke),
		}
	})
	return noopVtbl
}

// newNoopCallback returns a shared singleton usable as a WUA async callback.
// Because the callback is completely stateless (ncInvoke is a no-op), a single
// instance can be safely shared across all async calls. This avoids the
// unbounded memory growth that would result from allocating a new callback on
// every invocation. COM must already be initialized on the calling thread.
//
// The return type is *ole.IDispatch (not *ole.IUnknown) because go-ole's
// oleutil.CallMethod only handles *IDispatch in its type switch; passing
// *IUnknown causes a panic("unknown type"). The cast is safe: go-ole just
// uses the pointer value to build a VT_DISPATCH VARIANT, and the WUA method
// will QueryInterface our object for the actual callback interface it needs.
func newNoopCallback() *ole.IDispatch {
	globalNoopMu.Do(func() {
		globalNoopCb = &noopCallback{lpVtbl: getNoopVtbl(), ref: 1}

		// Aggregate the free-threaded marshaler so the callback is agile. The
		// controlling unknown is the callback itself; the FTM delegates non-IMarshal
		// QueryInterface calls back to us. If this fails we leave ftm=0 and fall back
		// to standard marshaling (BeginXxx may then fail and the caller falls back to
		// the synchronous path).
		var ftm uintptr
		if procCoCreateFreeThreadedMarshaler.Find() == nil {
			ret, _, _ := procCoCreateFreeThreadedMarshaler.Call(
				uintptr(unsafe.Pointer(globalNoopCb)),
				uintptr(unsafe.Pointer(&ftm)),
			)
			if ret == 0 {
				globalNoopCb.ftm = ftm
			}
		}
	})
	return (*ole.IDispatch)(unsafe.Pointer(globalNoopCb))
}
