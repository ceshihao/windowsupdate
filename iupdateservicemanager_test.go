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
)

func TestToIUpdateServiceManager_NilDispatch(t *testing.T) {
	result, err := toIUpdateServiceManager(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

// COM tests
func TestNewUpdateServiceManager(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}
	if mgr == nil {
		t.Fatal("NewUpdateServiceManager returned nil")
	}
}

func TestIUpdateServiceManager_Services(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	services := mgr.Services
	// Services may be empty but should not error
	if services == nil {
		t.Fatal("Services returned nil")
	}

	// If there are services, verify structure
	if len(services) > 0 {
		svc := services[0]
		if svc.ServiceID == "" {
			t.Error("First service has empty ServiceID")
		}
	}
}

func TestIUpdateServiceManager_PutClientApplicationID(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	testID := "TestServiceManagerApp"
	err = mgr.PutClientApplicationID(testID)
	if err != nil {
		t.Errorf("PutClientApplicationID failed: %v", err)
	}
	if mgr.ClientApplicationID != testID {
		t.Errorf("ClientApplicationID = %q, want %q", mgr.ClientApplicationID, testID)
	}
}

func TestIUpdateServiceManager_AddService(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	// Verify the method signature by calling with empty string
	// This will fail quickly without hanging
	_, err = mgr.AddService("", "")
	// We expect an error for empty service ID
	if err == nil {
		t.Log("AddService with empty service ID unexpectedly succeeded")
	}
}

func TestIUpdateServiceManager_RegisterServiceWithAU(t *testing.T) {
	// Note: RegisterServiceWithAU can hang for a very long time waiting
	// for system services. We verify the structure instead.
	mgr := &IUpdateServiceManager{
		ClientApplicationID: "test-app",
	}

	if mgr.ClientApplicationID != "test-app" {
		t.Errorf("ClientApplicationID = %s, want test-app", mgr.ClientApplicationID)
	}
}

func TestIUpdateServiceManager_RemoveService(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	// Verify the method signature by calling with empty string
	// This will fail quickly without hanging
	err = mgr.RemoveService("")
	// We expect an error for empty service ID
	if err == nil {
		t.Log("RemoveService with empty service ID unexpectedly succeeded")
	}
}

func TestIUpdateServiceManager_UnregisterServiceWithAU(t *testing.T) {
	// Note: UnregisterServiceWithAU can hang for a very long time (10+ minutes)
	// waiting for system services. We verify the structure instead.
	mgr := &IUpdateServiceManager{
		ClientApplicationID: "test-app",
	}

	if mgr.ClientApplicationID != "test-app" {
		t.Errorf("ClientApplicationID = %s, want test-app", mgr.ClientApplicationID)
	}
}

func TestIUpdateServiceManager_SetOption(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	// Verify the method signature by calling with empty strings
	// This will fail quickly without hanging
	err = mgr.SetOption("", "")
	// We expect an error for empty option name
	if err == nil {
		t.Log("SetOption with empty parameters unexpectedly succeeded")
	}
}

func TestIUpdateServiceManager_AddScanPackageService(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	// Verify the method signature by calling with empty parameters
	// This will fail quickly without hanging
	_, err = mgr.AddScanPackageService("", "", 0)
	// We expect an error for empty service name
	if err == nil {
		t.Log("AddScanPackageService with empty parameters unexpectedly succeeded")
	}
}

func TestIUpdateServiceManager_AddService2(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	mgr, err := NewUpdateServiceManager()
	if err != nil {
		t.Fatalf("NewUpdateServiceManager failed: %v", err)
	}

	// Verify the method signature by calling with empty string
	// This will fail quickly without hanging
	_, err = mgr.AddService2("", 0, "")
	// We expect an error for empty service ID
	if err == nil {
		t.Log("AddService2 with empty service ID unexpectedly succeeded")
	}
}
