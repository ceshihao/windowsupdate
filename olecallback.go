/*
Copyright 2022 Zheng Dayu
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

//go:build windows

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

type noopCallbackVtbl struct {
	pQueryInterface   uintptr
	pAddRef           uintptr
	pRelease          uintptr
	pGetTypeInfoCount uintptr
	pGetTypeInfo      uintptr
	pGetIDsOfNames    uintptr
	pInvoke           uintptr
}

// noopCallback : lpVtbl MUST be the first field (the COM interface pointer points
// to it).
type noopCallback struct {
	lpVtbl *noopCallbackVtbl
	ref    int32
}

// HRESULT values as uintptr (only the low 32 bits are significant).
const (
	hrSOK          = uintptr(0x00000000)
	hrENoInterface = uintptr(0x80004002)
	hrENotImpl     = uintptr(0x80004001)
)

func ncQueryInterface(this, iid, ppvObject uintptr) uintptr {
	guid := (*ole.GUID)(unsafe.Pointer(iid))
	out := (*uintptr)(unsafe.Pointer(ppvObject))
	if ole.IsEqualGUID(guid, ole.IID_IUnknown) || ole.IsEqualGUID(guid, ole.IID_IDispatch) {
		p := (*noopCallback)(unsafe.Pointer(this))
		p.ref++
		if out != nil {
			*out = this
		}
		return hrSOK
	}
	if out != nil {
		*out = 0
	}
	return hrENoInterface
}

func ncAddRef(this uintptr) uintptr {
	p := (*noopCallback)(unsafe.Pointer(this))
	p.ref++
	return uintptr(uint32(p.ref))
}

func ncRelease(this uintptr) uintptr {
	p := (*noopCallback)(unsafe.Pointer(this))
	p.ref--
	return uintptr(uint32(p.ref))
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
	return hrSOK
}

var (
	noopVtbl    *noopCallbackVtbl
	noopOnce    sync.Once
	keepAliveMu sync.Mutex
	keepAlive   []*noopCallback // pin the objects so the GC does not collect them while WUA holds them
)

func getNoopVtbl() *noopCallbackVtbl {
	noopOnce.Do(func() {
		noopVtbl = &noopCallbackVtbl{
			pQueryInterface:   syscall.NewCallback(ncQueryInterface),
			pAddRef:           syscall.NewCallback(ncAddRef),
			pRelease:          syscall.NewCallback(ncRelease),
			pGetTypeInfoCount: syscall.NewCallback(ncGetTypeInfoCount),
			pGetTypeInfo:      syscall.NewCallback(ncGetTypeInfo),
			pGetIDsOfNames:    syscall.NewCallback(ncGetIDsOfNames),
			pInvoke:           syscall.NewCallback(ncInvoke),
		}
	})
	return noopVtbl
}

// newNoopDispatch creates a minimal IDispatch usable as a WUA callback.
// The object is pinned (keepAlive) so it is not collected while WUA holds it.
func newNoopDispatch() *ole.IDispatch {
	cb := &noopCallback{lpVtbl: getNoopVtbl(), ref: 1}
	keepAliveMu.Lock()
	keepAlive = append(keepAlive, cb)
	keepAliveMu.Unlock()
	return (*ole.IDispatch)(unsafe.Pointer(cb))
}
