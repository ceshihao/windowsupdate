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

func TestToIWindowsDriverUpdateEntry_NilDispatch(t *testing.T) {
	result, err := toIWindowsDriverUpdateEntry(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestToIWindowsDriverUpdateEntries_NilDispatch(t *testing.T) {
	result, err := toIWindowsDriverUpdateEntries(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIWindowsDriverUpdate_StructureFields(t *testing.T) {
	wdu := &IWindowsDriverUpdate{
		DeviceProblemNumber: 28,
		DeviceStatus:        0,
		DriverClass:         "Display",
		DriverHardwareID:    "PCI\\VEN_1234&DEV_5678",
		DriverManufacturer:  "Test Manufacturer",
		DriverModel:         "Test Model",
		DriverProvider:      "Test Provider",
		AutoDownload2:       1,
		AutoSelection2:      2,
		RebootRequired2:     true,
		IsPresent2:          true,
		BrowseOnly2:         false,
	}

	if wdu.DeviceProblemNumber != 28 {
		t.Errorf("DeviceProblemNumber = %d, want 28", wdu.DeviceProblemNumber)
	}
	if wdu.DriverClass != "Display" {
		t.Errorf("DriverClass = %s, want Display", wdu.DriverClass)
	}
	if wdu.DriverHardwareID != "PCI\\VEN_1234&DEV_5678" {
		t.Errorf("DriverHardwareID = %s", wdu.DriverHardwareID)
	}
	if !wdu.RebootRequired2 {
		t.Errorf("RebootRequired2 = %v, want true", wdu.RebootRequired2)
	}
	if !wdu.IsPresent2 {
		t.Errorf("IsPresent2 = %v, want true", wdu.IsPresent2)
	}
	if wdu.BrowseOnly2 {
		t.Errorf("BrowseOnly2 = %v, want false", wdu.BrowseOnly2)
	}
}

func TestIWindowsDriverUpdate_AllDriverFields(t *testing.T) {
	wdu := &IWindowsDriverUpdate{
		DeviceProblemNumber: 0,
		DeviceStatus:        22,
		DriverClass:         "USB",
		DriverHardwareID:    "USB\\VID_1234&PID_5678",
		DriverManufacturer:  "Microsoft",
		DriverModel:         "USB 3.0 Controller",
		DriverProvider:      "Microsoft Corporation",
		AutoDownload2:       0,
		AutoSelection2:      1,
	}

	if wdu.DeviceStatus != 22 {
		t.Errorf("DeviceStatus = %d, want 22", wdu.DeviceStatus)
	}
	if wdu.DriverManufacturer != "Microsoft" {
		t.Errorf("DriverManufacturer = %s, want Microsoft", wdu.DriverManufacturer)
	}
	if wdu.DriverModel != "USB 3.0 Controller" {
		t.Errorf("DriverModel = %s", wdu.DriverModel)
	}
	if wdu.DriverProvider != "Microsoft Corporation" {
		t.Errorf("DriverProvider = %s", wdu.DriverProvider)
	}
}

func TestIWindowsDriverUpdateEntry_StructureFields(t *testing.T) {
	entry := &IWindowsDriverUpdateEntry{
		DeviceProblemNumber: 10,
		DeviceStatus:        1,
		DriverClass:         "Network",
		DriverHardwareID:    "USB\\VEN_ABCD&DEV_EFGH",
		DriverManufacturer:  "Entry Manufacturer",
		DriverModel:         "Entry Model",
		DriverProvider:      "Entry Provider",
	}

	if entry.DeviceProblemNumber != 10 {
		t.Errorf("DeviceProblemNumber = %d, want 10", entry.DeviceProblemNumber)
	}
	if entry.DriverClass != "Network" {
		t.Errorf("DriverClass = %s, want Network", entry.DriverClass)
	}
	if entry.DeviceStatus != 1 {
		t.Errorf("DeviceStatus = %d, want 1", entry.DeviceStatus)
	}
	if entry.DriverHardwareID != "USB\\VEN_ABCD&DEV_EFGH" {
		t.Errorf("DriverHardwareID = %s", entry.DriverHardwareID)
	}
	if entry.DriverManufacturer != "Entry Manufacturer" {
		t.Errorf("DriverManufacturer = %s", entry.DriverManufacturer)
	}
	if entry.DriverModel != "Entry Model" {
		t.Errorf("DriverModel = %s", entry.DriverModel)
	}
	if entry.DriverProvider != "Entry Provider" {
		t.Errorf("DriverProvider = %s", entry.DriverProvider)
	}
}

func TestIUpdate_ToWindowsDriverUpdate(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Create an update session and searcher to get real updates
	session, err := NewUpdateSession()
	if err != nil {
		t.Skipf("NewUpdateSession failed: %v", err)
	}

	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		t.Skipf("CreateUpdateSearcher failed: %v", err)
	}

	// Search for any available updates (including driver updates)
	// Use a simple criteria to get some updates
	result, err := searcher.Search("IsInstalled=0 and Type='Driver'")
	if err != nil {
		// If driver search fails, try a broader search
		result, err = searcher.Search("IsInstalled=0")
		if err != nil {
			t.Skipf("Search failed: %v", err)
		}
	}

	if result == nil || len(result.Updates) == 0 {
		t.Skip("No updates available to test ToWindowsDriverUpdate")
	}

	// Try to convert the first update to a driver update
	update := result.Updates[0]

	// Attempt to convert to driver update
	// This will return nil if it's not a driver update, which is fine
	driverUpdate, err := update.ToWindowsDriverUpdate()
	if err != nil {
		t.Logf("ToWindowsDriverUpdate failed (may not be a driver update): %v", err)
	}

	// If it is a driver update, verify the structure
	if driverUpdate != nil {
		t.Logf("Successfully converted to driver update")
		t.Logf("Driver Class: %s", driverUpdate.DriverClass)
		t.Logf("Driver Manufacturer: %s", driverUpdate.DriverManufacturer)

		// Verify that driver-specific fields are populated
		if driverUpdate.DriverClass == "" && driverUpdate.DriverManufacturer == "" {
			t.Error("Driver update should have at least some driver-specific fields populated")
		}
	} else {
		t.Log("Update is not a driver update (expected for non-driver updates)")
	}
}

func TestToIWindowsDriverUpdateEntry_WithRealData(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Try to get real driver updates to test entry conversion
	session, err := NewUpdateSession()
	if err != nil {
		t.Skipf("NewUpdateSession failed: %v", err)
	}

	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		t.Skipf("CreateUpdateSearcher failed: %v", err)
	}

	result, err := searcher.Search("Type='Driver'")
	if err != nil {
		t.Skipf("Search for driver updates failed: %v", err)
	}

	if result == nil || len(result.Updates) == 0 {
		t.Skip("No driver updates available to test entries")
	}

	// Try to get a driver update with entries
	maxChecks := len(result.Updates)
	if maxChecks > 5 {
		maxChecks = 5
	}

	for i := 0; i < maxChecks; i++ {
		update := result.Updates[i]

		driverUpdate, err := update.ToWindowsDriverUpdate()
		if err != nil || driverUpdate == nil {
			continue
		}

		// Check if it has WindowsDriverUpdateEntries
		if len(driverUpdate.WindowsDriverUpdateEntries) > 0 {
			t.Logf("Found driver update with %d entries", len(driverUpdate.WindowsDriverUpdateEntries))

			entry := driverUpdate.WindowsDriverUpdateEntries[0]
			if entry.DriverClass != "" {
				t.Logf("Entry DriverClass: %s", entry.DriverClass)
			}
			if entry.DriverManufacturer != "" {
				t.Logf("Entry DriverManufacturer: %s", entry.DriverManufacturer)
			}

			return // Test passed
		}
	}

	t.Skip("No driver updates with entries found")
}

func TestToIWindowsDriverUpdateEntries_Coverage(t *testing.T) {
	// This test verifies the nil handling which is already covered
	// The main logic requires real COM objects which are tested above

	result, err := toIWindowsDriverUpdateEntries(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}

	// Additional coverage note: The loop logic in toIWindowsDriverUpdateEntries
	// is tested through TestToIWindowsDriverUpdateEntry_WithRealData when
	// real driver updates with entries are available
}
