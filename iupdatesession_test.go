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

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func TestIUpdateSession_StructureFields(t *testing.T) {
	session := &IUpdateSession{
		ClientApplicationID: "my-app",
		ReadOnly:            true,
		WebProxy:            nil,
	}
	if session.ClientApplicationID != "my-app" {
		t.Errorf("ClientApplicationID not set correctly, got %s", session.ClientApplicationID)
	}
	if !session.ReadOnly {
		t.Errorf("ReadOnly not set correctly, got %v", session.ReadOnly)
	}
	if session.WebProxy != nil {
		t.Errorf("WebProxy should be nil, got %v", session.WebProxy)
	}
}

// COM tests
func TestNewUpdateSession(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	session, err := NewUpdateSession()
	if err != nil {
		t.Fatalf("NewUpdateSession failed: %v", err)
	}
	if session == nil {
		t.Fatal("NewUpdateSession returned nil")
	}
	if session.disp == nil {
		t.Fatal("session.disp is nil")
	}
	// ClientApplicationID can be empty by default, it's an optional property
	// that applications can set to identify themselves
}

func TestToIUpdateSession(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("Microsoft.Update.Session")
	if err != nil {
		t.Fatalf("CreateObject failed: %v", err)
	}
	defer unknown.Release()

	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		t.Fatalf("QueryInterface failed: %v", err)
	}
	defer disp.Release()

	session, err := toIUpdateSession(disp)
	if err != nil {
		t.Fatalf("toIUpdateSession failed: %v", err)
	}
	if session == nil {
		t.Fatal("toIUpdateSession returned nil")
	}
	// ClientApplicationID can be empty by default, it's an optional property
	// that applications can set to identify themselves
}

func TestIUpdateSession_CreateUpdateSearcher(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	session, err := NewUpdateSession()
	if err != nil {
		t.Fatalf("NewUpdateSession failed: %v", err)
	}

	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		t.Fatalf("CreateUpdateSearcher failed: %v", err)
	}
	if searcher == nil {
		t.Fatal("CreateUpdateSearcher returned nil")
	}
	if searcher.disp == nil {
		t.Fatal("searcher.disp is nil")
	}
}

func TestIUpdateSession_CreateUpdateDownloader(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	session, err := NewUpdateSession()
	if err != nil {
		t.Fatalf("NewUpdateSession failed: %v", err)
	}

	downloader, err := session.CreateUpdateDownloader()
	if err != nil {
		t.Fatalf("CreateUpdateDownloader failed: %v", err)
	}
	if downloader == nil {
		t.Fatal("CreateUpdateDownloader returned nil")
	}
	if downloader.disp == nil {
		t.Fatal("downloader.disp is nil")
	}
}

func TestIUpdateSession_CreateUpdateInstaller(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	session, err := NewUpdateSession()
	if err != nil {
		t.Fatalf("NewUpdateSession failed: %v", err)
	}

	installer, err := session.CreateUpdateInstaller()
	if err != nil {
		t.Fatalf("CreateUpdateInstaller failed: %v", err)
	}
	if installer == nil {
		t.Fatal("CreateUpdateInstaller returned nil")
	}
	if installer.disp == nil {
		t.Fatal("installer.disp is nil")
	}
}
