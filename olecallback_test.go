//go:build windows
// +build windows

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

package windowsupdate

import (
	"testing"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func TestNewNoopCallback_SharedSingleton(t *testing.T) {
	first := newNoopCallback()
	if first == nil {
		t.Fatal("newNoopCallback returned nil")
	}
	second := newNoopCallback()
	if first != second {
		t.Errorf("expected newNoopCallback to return the shared singleton, got %p and %p", first, second)
	}
}

func TestNoopCallback_QueryInterface(t *testing.T) {
	cb := getNoopVtbl()
	this := uintptr(unsafe.Pointer(cb))

	// A requested IID of IUnknown or IDispatch must succeed and return the object.
	for _, iid := range []*ole.GUID{ole.IID_IUnknown, ole.IID_IDispatch} {
		var out uintptr
		hr := ncQueryInterface(this, uintptr(unsafe.Pointer(iid)), uintptr(unsafe.Pointer(&out)))
		if hr != hrSOK {
			t.Errorf("QueryInterface(%v) = 0x%x, want S_OK", iid, hr)
		}
		if out != this {
			t.Errorf("QueryInterface(%v) out = 0x%x, want 0x%x", iid, out, this)
		}
	}

	// An unsupported IID must fail with E_NOINTERFACE and clear the out pointer.
	unsupported := &ole.GUID{Data1: 0xdeadbeef, Data2: 0x1234, Data3: 0x5678}
	out := uintptr(0xfff)
	hr := ncQueryInterface(this, uintptr(unsafe.Pointer(unsupported)), uintptr(unsafe.Pointer(&out)))
	if hr != hrENoInterface {
		t.Errorf("QueryInterface(unsupported) = 0x%x, want E_NOINTERFACE", hr)
	}
	if out != 0 {
		t.Errorf("QueryInterface(unsupported) out = 0x%x, want 0", out)
	}
}

func TestNoopCallback_AddRefRelease(t *testing.T) {
	cb := &noopCallback{lpVtbl: getNoopVtbl(), ref: 1}
	this := uintptr(unsafe.Pointer(cb))

	if got := ncAddRef(this); got != 2 {
		t.Errorf("AddRef = %d, want 2", got)
	}
	if got := ncRelease(this); got != 1 {
		t.Errorf("Release = %d, want 1", got)
	}
}
