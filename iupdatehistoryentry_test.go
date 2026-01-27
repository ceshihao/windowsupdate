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
	"time"
)

func TestIUpdateHistoryEntry_StructureFields(t *testing.T) {
	now := time.Now()
	entry := &IUpdateHistoryEntry{
		ClientApplicationID: "test-app",
		Date:                &now,
		Description:         "Test update",
		HResult:             0,
		Operation:           UpdateOperationUoInstallation,
		ResultCode:          OperationResultCodeOrcSucceeded,
		ServerSelection:     ServerSelectionSsWindowsUpdate,
		ServiceID:           "service-id",
		SupportUrl:          "https://example.com",
		Title:               "Test Update",
		UninstallationNotes: "Notes",
		UninstallationSteps: []string{"step1", "step2"},
		UnmappedResultCode:  0,
	}
	if entry.ClientApplicationID != "test-app" {
		t.Errorf("ClientApplicationID not set correctly, got %s", entry.ClientApplicationID)
	}
	if entry.Operation != UpdateOperationUoInstallation {
		t.Errorf("Operation not set correctly")
	}
	if entry.Title != "Test Update" {
		t.Errorf("Title not set correctly, got %s", entry.Title)
	}
	if entry.Date == nil {
		t.Errorf("Date should not be nil")
	}
	if len(entry.UninstallationSteps) != 2 {
		t.Errorf("UninstallationSteps length incorrect, got %d", len(entry.UninstallationSteps))
	}
}
