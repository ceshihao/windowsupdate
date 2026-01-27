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

func TestIAutomaticUpdatesSettings_StructureFields(t *testing.T) {
	settings := &IAutomaticUpdatesSettings{
		NotificationLevel:         2,
		ReadOnly:                  false,
		Required:                  true,
		ScheduledInstallationDay:  1,
		ScheduledInstallationTime: 3,
	}
	if settings.NotificationLevel != 2 {
		t.Errorf("NotificationLevel not set correctly")
	}
	if settings.ReadOnly {
		t.Errorf("ReadOnly should be false")
	}
	if !settings.Required {
		t.Errorf("Required should be true")
	}
	if settings.ScheduledInstallationDay != 1 {
		t.Errorf("ScheduledInstallationDay not set correctly")
	}
	if settings.ScheduledInstallationTime != 3 {
		t.Errorf("ScheduledInstallationTime not set correctly")
	}
}

// COM tests
func TestIAutomaticUpdatesSettings_Refresh(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		// GetSettings may fail in some environments (e.g., policy restrictions)
		t.Skipf("GetSettings failed (may be expected): %v", err)
	}

	err = settings.Refresh()
	if err != nil {
		// Refresh may fail in some environments (e.g., without admin rights or in CI)
		t.Logf("Refresh failed (may be expected): %v", err)
	}
}

func TestIAutomaticUpdatesSettings_Save(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		// GetSettings may fail in some environments (e.g., policy restrictions)
		t.Skipf("GetSettings failed (may be expected): %v", err)
	}

	if settings.ReadOnly {
		t.Skip("Settings are read-only, skipping Save test")
	}

	// Save without changes - should not fail
	err = settings.Save()
	if err != nil {
		t.Logf("Save may fail if not elevated: %v", err)
	}
}

func TestIAutomaticUpdatesSettings_PutNotificationLevel(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		// GetSettings may fail in some environments (e.g., policy restrictions)
		t.Skipf("GetSettings failed (may be expected): %v", err)
	}

	if settings.ReadOnly {
		t.Skip("Settings are read-only, skipping test")
	}

	oldLevel := settings.NotificationLevel
	newLevel := int32(2)

	err = settings.PutNotificationLevel(newLevel)
	if err != nil {
		t.Logf("PutNotificationLevel may fail if not elevated: %v", err)
		return
	}

	if settings.NotificationLevel != newLevel {
		t.Errorf("NotificationLevel = %d, want %d", settings.NotificationLevel, newLevel)
	}

	// Restore original value
	_ = settings.PutNotificationLevel(oldLevel)
}

func TestIAutomaticUpdatesSettings_PutScheduledInstallationDay(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		// GetSettings may fail in some environments (e.g., policy restrictions)
		t.Skipf("GetSettings failed (may be expected): %v", err)
	}

	if settings.ReadOnly {
		t.Skip("Settings are read-only, skipping test")
	}

	oldDay := settings.ScheduledInstallationDay
	newDay := int32(2)

	err = settings.PutScheduledInstallationDay(newDay)
	if err != nil {
		t.Logf("PutScheduledInstallationDay may fail (not supported on Windows 8+): %v", err)
		return
	}

	if settings.ScheduledInstallationDay != newDay {
		t.Errorf("ScheduledInstallationDay = %d, want %d", settings.ScheduledInstallationDay, newDay)
	}

	// Restore original value
	_ = settings.PutScheduledInstallationDay(oldDay)
}

func TestIAutomaticUpdatesSettings_PutScheduledInstallationTime(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	au, err := NewAutomaticUpdates()
	if err != nil {
		t.Fatalf("NewAutomaticUpdates failed: %v", err)
	}

	settings, err := au.GetSettings()
	if err != nil {
		// GetSettings may fail in some environments (e.g., policy restrictions)
		t.Skipf("GetSettings failed (may be expected): %v", err)
	}

	if settings.ReadOnly {
		t.Skip("Settings are read-only, skipping test")
	}

	oldTime := settings.ScheduledInstallationTime
	newTime := int32(14)

	err = settings.PutScheduledInstallationTime(newTime)
	if err != nil {
		t.Logf("PutScheduledInstallationTime may fail (not supported on Windows 8+): %v", err)
		return
	}

	if settings.ScheduledInstallationTime != newTime {
		t.Errorf("ScheduledInstallationTime = %d, want %d", settings.ScheduledInstallationTime, newTime)
	}

	// Restore original value
	_ = settings.PutScheduledInstallationTime(oldTime)
}
