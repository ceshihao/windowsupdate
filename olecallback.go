//go:build windows

/*
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
// a non-NULL IUnknown* callback argument. Passing NULL (VT_NULL) makes them fail
// with DISP_E_TYPEMISMATCH (0x80020005). newNoopDispatch returns a minimal
// IDispatch whose Invoke does nothing (returns S_OK): completion is obtained
// through the blocking EndXxx methods and progress through IXxxJob.GetProgress().
//
// The handler signatures are 100% uintptr because that is required by
// syscall.NewCallback.

// noopCallbackVtbl is the COM virtual function table layout for IDispatch.
// The order of fields MUST match the IUnknown + IDispatch v-table layout.
type noopCallbackVtbl struct {
	pQueryInterface   uintptr
	pAddRef           uintptr
	pRelease          uintptr
	pGetTypeInfoCount uintptr
	pGetTypeInfo      uintptr
	pGetIDsOfNames    uintptr
	pInvoke           uintptr
}

// noopCallback is a stateless dummy IDispatch implementation. lpVtbl MUST be
// the first field because the COM interface pointer points directly to it.
type noopCallback struct {
	lpVtbl *noopCallbackVtbl
	ref    int32
}

// HRESULT values as uintptr (only the low 32 bits are significant).
const (
	hrSOK          = uintptr(0x00000000)
	hrEPointer     = uintptr(0x80004003)
	hrENoInterface = uintptr(0x80004002)
	hrENotImpl     = uintptr(0x80004001)
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
	if ole.IsEqualGUID(guid, ole.IID_IUnknown) || ole.IsEqualGUID(guid, ole.IID_IDispatch) {
		atomic.AddInt32(&globalNoop.ref, 1)
		*out = this
		return hrSOK
	}
	*out = 0
	return hrENoInterface
}

func ncAddRef(this uintptr) uintptr {
	return uintptr(uint32(atomic.AddInt32(&globalNoop.ref, 1)))
}

func ncRelease(this uintptr) uintptr {
	// Singleton object: it is never actually freed even if the count reaches
	// zero. We still maintain the counter so the value returned to the COM
	// caller is meaningful.
	return uintptr(uint32(atomic.AddInt32(&globalNoop.ref, -1)))
}

func ncGetTypeInfoCount(this, pctinfo uintptr) uintptr {
	if pctinfo != 0 {
		*(*uint32)(unsafe.Pointer(pctinfo)) = 0
	}
	return hrSOK
}

func ncGetTypeInfo(this, iTInfo, lcid, ppTInfo uintptr) uintptr {
	return hrENotImpl
}

func ncGetIDsOfNames(this, riid, rgszNames, cNames, lcid, rgDispId uintptr) uintptr {
	return hrENotImpl
}

// ncInvoke : no-op body. WUA calls DISPID 0 on progress/completion; we ignore it
// and return S_OK. Completion is detected through EndXxx (blocking).
func ncInvoke(this, dispIdMember, riid, lcid, wFlags, pDispParams, pVarResult, pExcepInfo, puArgErr uintptr) uintptr {
	if pVarResult != 0 {
		v := (*ole.VARIANT)(unsafe.Pointer(pVarResult))
		v.VT = ole.VT_EMPTY
	}
	return hrSOK
}

var (
	noopOnce   sync.Once
	globalNoop *noopCallback
)

// newNoopDispatch returns a pointer to a global singleton IDispatch usable as
// a WUA callback. Because the callback is completely stateless, a single
// instance can be safely shared across all async calls. This avoids the
// unbounded memory leak that would result from allocating a new callback on
// every invocation and pinning it in a global slice.
func newNoopDispatch() *ole.IDispatch {
	noopOnce.Do(func() {
		vtbl := &noopCallbackVtbl{
			pQueryInterface:   syscall.NewCallback(ncQueryInterface),
			pAddRef:           syscall.NewCallback(ncAddRef),
			pRelease:          syscall.NewCallback(ncRelease),
			pGetTypeInfoCount: syscall.NewCallback(ncGetTypeInfoCount),
			pGetTypeInfo:      syscall.NewCallback(ncGetTypeInfo),
			pGetIDsOfNames:    syscall.NewCallback(ncGetIDsOfNames),
			pInvoke:           syscall.NewCallback(ncInvoke),
		}
		globalNoop = &noopCallback{lpVtbl: vtbl, ref: 1}
	})
	return (*ole.IDispatch)(unsafe.Pointer(globalNoop))
}
