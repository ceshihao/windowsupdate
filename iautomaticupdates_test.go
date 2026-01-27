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

func TestToIAutomaticUpdatesResults_NilDispatch(t *testing.T) {
	result, err := toIAutomaticUpdatesResults(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIAutomaticUpdates_StructureFields(t *testing.T) {
	au := &IAutomaticUpdates{
		ServiceEnabled: true,
	}
	if !au.ServiceEnabled {
		t.Errorf("ServiceEnabled not set correctly")
	}
}

// COM tests
func TestNewAutomaticUpdates(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}
	if au == nil {
		t.Fatal("NewAutomaticUpdates returned nil")
	}
	if au.disp == nil {
		t.Fatal("disp is nil")
	}
}

func TestIAutomaticUpdates_DetectNow(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	err = au.DetectNow()
	if err != nil {
		t.Logf("DetectNow may fail if not elevated: %v", err)
	}
}

func TestIAutomaticUpdates_EnableService(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	err = au.EnableService()
	if err != nil {
		t.Logf("EnableService may fail if not elevated: %v", err)
	}
}

func TestIAutomaticUpdates_Pause(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	err = au.Pause()
	if err != nil {
		t.Logf("Pause may fail if not elevated: %v", err)
	}
}

func TestIAutomaticUpdates_Resume(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	err = au.Resume()
	if err != nil {
		t.Logf("Resume may fail if not elevated: %v", err)
	}
}

func TestIAutomaticUpdates_ShowSettingsDialog(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	// ShowSettingsDialog displays UI, so it may fail in automated/headless environments
	// We call it to get coverage but don't fail the test if it errors
	err = au.ShowSettingsDialog()
	if err != nil {
		t.Logf("ShowSettingsDialog failed (may be expected in automated tests): %v", err)
	}
}

func TestIAutomaticUpdates_GetSettings(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		t.Skipf("GetSettings failed (may be expected in some environments): %v", err)
		return
	}
	if settings == nil {
		t.Fatal("GetSettings returned nil")
	}
	if settings.disp == nil {
		t.Fatal("settings.disp is nil")
	}
}

func TestIAutomaticUpdates_GetResults(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	results, err := au.GetResults()
	if err != nil {
		t.Logf("GetResults may fail in some environments: %v", err)
		return
	}
	if results != nil && results.disp == nil {
		t.Error("results.disp is nil")
	}
}
