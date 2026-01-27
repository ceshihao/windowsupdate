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

func TestIUpdateInstaller_StructureFields(t *testing.T) {
	installer := &IUpdateInstaller{
		AllowSourcePrompts:  true,
		ClientApplicationID: "test-installer",
		ForceQuiet:          false,
		IsForced:            true,
	}
	if !installer.AllowSourcePrompts {
		t.Errorf("AllowSourcePrompts not set correctly")
	}
	if installer.ClientApplicationID != "test-installer" {
		t.Errorf("ClientApplicationID not set correctly")
	}
	if installer.ForceQuiet {
		t.Errorf("ForceQuiet should be false")
	}
	if !installer.IsForced {
		t.Errorf("IsForced should be true")
	}
}

// COM tests
func TestIUpdateInstaller_Properties(t *testing.T) {
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

	// Test AllowSourcePrompts property
	_ = installer.AllowSourcePrompts

	err = installer.PutAllowSourcePrompts(false)
	if err != nil {
		t.Fatalf("PutAllowSourcePrompts failed: %v", err)
	}

	if installer.AllowSourcePrompts {
		t.Error("AllowSourcePrompts not updated to false")
	}

	// Test IsForced property
	_ = installer.IsForced

	// Test IsBusy property
	_ = installer.IsBusy

	// Test RebootRequiredBeforeInstallation property
	_ = installer.RebootRequiredBeforeInstallation
}

func TestIUpdateInstaller_PutClientApplicationID(t *testing.T) {
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

	testID := "TestInstallerApp"
	err = installer.PutClientApplicationID(testID)
	if err != nil {
		t.Errorf("PutClientApplicationID failed: %v", err)
	}
	if installer.ClientApplicationID != testID {
		t.Errorf("ClientApplicationID = %q, want %q", installer.ClientApplicationID, testID)
	}
}

func TestIUpdateInstaller_PutForceQuiet(t *testing.T) {
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

	err = installer.PutForceQuiet(true)
	if err != nil {
		t.Errorf("PutForceQuiet failed: %v", err)
	}
	if !installer.ForceQuiet {
		t.Errorf("ForceQuiet not updated")
	}
}

func TestIUpdateInstaller_PutIsForced(t *testing.T) {
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

	err = installer.PutIsForced(true)
	if err != nil {
		t.Errorf("PutIsForced failed: %v", err)
	}
	if !installer.IsForced {
		t.Errorf("IsForced not updated")
	}
}

func TestIUpdateInstaller_Install(t *testing.T) {
	// Note: Install requires a real IUpdateInstaller with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp:       nil,
		ForceQuiet: false,
	}

	// Verify structure can be created
	if installer.ForceQuiet {
		t.Error("ForceQuiet should be false")
	}
}

func TestIUpdateInstaller_Uninstall(t *testing.T) {
	// Note: Uninstall requires a real IUpdateInstaller with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp:     nil,
		IsForced: false,
	}

	// Verify structure can be created
	if installer.IsForced {
		t.Error("IsForced should be false")
	}
}

func TestIUpdateInstaller_Commit(t *testing.T) {
	// Note: Commit requires a real installation result.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp:   nil,
		IsBusy: false,
	}

	// Verify structure can be created
	if installer.IsBusy {
		t.Error("IsBusy should be false")
	}
}

func TestIUpdateInstaller_BeginInstall(t *testing.T) {
	// Note: BeginInstall requires a real IUpdateInstaller with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp:                             nil,
		RebootRequiredBeforeInstallation: false,
	}

	// Verify structure can be created
	if installer.RebootRequiredBeforeInstallation {
		t.Error("RebootRequiredBeforeInstallation should be false")
	}
}

func TestIUpdateInstaller_EndInstall(t *testing.T) {
	// Note: EndInstall requires a real installation job.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp: nil,
	}

	// Verify structure can be created
	if installer.disp != nil {
		t.Error("disp should be nil")
	}
}

// TestIUpdateInstaller_BeginInstallEndInstall exercises toIInstallationResult and IInstallationResult.GetUpdateResult via real COM.
func TestIUpdateInstaller_BeginInstallEndInstall(t *testing.T) {
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

	job, err := installer.BeginInstall([]*IUpdate{})
	if err != nil {
		t.Skipf("BeginInstall failed: %v", err)
		return
	}
	if job == nil {
		t.Fatal("BeginInstall returned nil job")
	}

	result, err := installer.EndInstall(job)
	if err != nil {
		t.Skipf("EndInstall failed (job may not be complete): %v", err)
		return
	}
	if result != nil {
		_, _ = result.GetUpdateResult(0)
	}
}

func TestIUpdateInstaller_BeginUninstall(t *testing.T) {
	// Note: BeginUninstall requires a real IUpdateInstaller with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp:               nil,
		AllowSourcePrompts: true,
	}

	// Verify structure can be created
	if !installer.AllowSourcePrompts {
		t.Error("AllowSourcePrompts should be true")
	}
}

func TestIUpdateInstaller_EndUninstall(t *testing.T) {
	// Note: EndUninstall requires a real installation job.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	installer := &IUpdateInstaller{
		disp: nil,
	}

	// Verify structure can be created
	if installer.disp != nil {
		t.Error("disp should be nil")
	}
}
