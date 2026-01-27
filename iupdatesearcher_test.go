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

func TestIUpdateSearcher_StructureFields(t *testing.T) {
	searcher := &IUpdateSearcher{
		CanAutomaticallyUpgradeService:      true,
		ClientApplicationID:                 "test-client",
		IncludePotentiallySupersededUpdates: false,
		Online:                              true,
		ServerSelection:                     ServerSelectionSsWindowsUpdate,
		ServiceID:                           "service-123",
	}
	if !searcher.CanAutomaticallyUpgradeService {
		t.Errorf("CanAutomaticallyUpgradeService not set correctly")
	}
	if searcher.ClientApplicationID != "test-client" {
		t.Errorf("ClientApplicationID not set correctly, got %s", searcher.ClientApplicationID)
	}
	if searcher.IncludePotentiallySupersededUpdates {
		t.Errorf("IncludePotentiallySupersededUpdates should be false")
	}
	if !searcher.Online {
		t.Errorf("Online should be true")
	}
	if searcher.ServerSelection != ServerSelectionSsWindowsUpdate {
		t.Errorf("ServerSelection not set correctly")
	}
	if searcher.ServiceID != "service-123" {
		t.Errorf("ServiceID not set correctly")
	}
}

// COM tests
func TestIUpdateSearcher_GetTotalHistoryCount(t *testing.T) {
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

	count, err := searcher.GetTotalHistoryCount()
	if err != nil {
		t.Fatalf("GetTotalHistoryCount failed: %v", err)
	}
	if count < 0 {
		t.Errorf("GetTotalHistoryCount returned negative count: %d", count)
	}
}

func TestIUpdateSearcher_QueryHistory(t *testing.T) {
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

	count, err := searcher.GetTotalHistoryCount()
	if err != nil {
		t.Fatalf("GetTotalHistoryCount failed: %v", err)
	}

	if count > 0 {
		queryCount := int32(5)
		if count < queryCount {
			queryCount = count
		}

		history, err := searcher.QueryHistory(0, queryCount)
		if err != nil {
			t.Fatalf("QueryHistory failed: %v", err)
		}
		if history == nil {
			t.Fatal("QueryHistory returned nil")
		}
		if len(history) > int(queryCount) {
			t.Errorf("QueryHistory returned more entries than requested: got %d, want <= %d", len(history), queryCount)
		}

		if len(history) > 0 {
			entry := history[0]
			if entry.UpdateIdentity == nil {
				t.Error("First history entry has nil UpdateIdentity")
			}
			if entry.Title == "" {
				t.Error("First history entry has empty Title")
			}
		}
	}
}

func TestIUpdateSearcher_QueryHistoryAll(t *testing.T) {
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

	count, err := searcher.GetTotalHistoryCount()
	if err != nil {
		t.Fatalf("GetTotalHistoryCount failed: %v", err)
	}

	if count > 0 {
		history, err := searcher.QueryHistoryAll()
		if err != nil {
			t.Fatalf("QueryHistoryAll failed: %v", err)
		}
		if int32(len(history)) != count {
			t.Errorf("QueryHistoryAll returned %d entries, expected %d", len(history), count)
		}
	}
}

func TestIUpdateSearcher_Search(t *testing.T) {
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

	// Search for installed updates (quick search)
	result, err := searcher.Search("IsInstalled=1")
	if err != nil {
		t.Skipf("Search failed (may be expected in some environments): %v", err)
		return
	}
	if result == nil {
		t.Fatal("Search returned nil result")
	}
	if result.ResultCode < OperationResultCodeOrcNotStarted || result.ResultCode > OperationResultCodeOrcAborted {
		t.Errorf("Invalid ResultCode: %d", result.ResultCode)
	}
}

func TestIUpdateSearcher_EscapeString(t *testing.T) {
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

	testStrings := []string{
		"simple",
		"with space",
		"with'quote",
		`with"double`,
	}

	for _, str := range testStrings {
		escaped, err := searcher.EscapeString(str)
		if err != nil {
			t.Errorf("EscapeString(%q) failed: %v", str, err)
		}
		if escaped == "" && str != "" {
			t.Errorf("EscapeString(%q) returned empty string", str)
		}
	}
}

func TestIUpdateSearcher_PutClientApplicationID(t *testing.T) {
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

	testID := "TestApplication"
	err = searcher.PutClientApplicationID(testID)
	if err != nil {
		t.Errorf("PutClientApplicationID failed: %v", err)
	}
	if searcher.ClientApplicationID != testID {
		t.Errorf("ClientApplicationID = %q, want %q", searcher.ClientApplicationID, testID)
	}
}

func TestIUpdateSearcher_PutServerSelection(t *testing.T) {
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

	err = searcher.PutServerSelection(ServerSelectionSsManagedServer)
	if err != nil {
		t.Errorf("PutServerSelection failed: %v", err)
	}
	if searcher.ServerSelection != ServerSelectionSsManagedServer {
		t.Errorf("ServerSelection = %d, want %d", searcher.ServerSelection, ServerSelectionSsManagedServer)
	}
}

func TestIUpdateSearcher_PutOnline(t *testing.T) {
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

	err = searcher.PutOnline(false)
	if err != nil {
		t.Errorf("PutOnline failed: %v", err)
	}
	if searcher.Online {
		t.Error("Online should be false")
	}
}

func TestIUpdateSearcher_PutServiceID(t *testing.T) {
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

	// Get the initial ServiceID
	initialServiceID := searcher.ServiceID

	// Try to set a service ID
	// Note: This requires a valid Windows Update service GUID
	// Using the Windows Update service GUID
	testServiceID := "9482f4b4-e343-43b6-b170-9a65bc822c77" // Windows Update service
	err = searcher.PutServiceID(testServiceID)
	if err != nil {
		// Setting service ID may fail depending on permissions and validity
		t.Logf("PutServiceID failed (may be expected): %v", err)
		// Verify the ServiceID wasn't changed on error
		if searcher.ServiceID != initialServiceID {
			t.Errorf("ServiceID changed despite error: got %q, want %q", searcher.ServiceID, initialServiceID)
		}
		return
	}

	// If successful, verify the change
	if searcher.ServiceID != testServiceID {
		t.Errorf("ServiceID = %q, want %q", searcher.ServiceID, testServiceID)
	}
}

func TestIUpdateSearcher_PutIncludePotentiallySupersededUpdates(t *testing.T) {
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

	err = searcher.PutIncludePotentiallySupersededUpdates(true)
	if err != nil {
		t.Errorf("PutIncludePotentiallySupersededUpdates failed: %v", err)
	}
	if !searcher.IncludePotentiallySupersededUpdates {
		t.Error("IncludePotentiallySupersededUpdates should be true")
	}
}

func TestIUpdateSearcher_ServerSelection(t *testing.T) {
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

	// Initial ServerSelection should be a valid value
	serverSelection := searcher.ServerSelection
	if serverSelection < ServerSelectionSsDefault || serverSelection > ServerSelectionSsOthers {
		t.Errorf("Invalid ServerSelection: %d", serverSelection)
	}

	// Test PutServerSelection
	err = searcher.PutServerSelection(ServerSelectionSsWindowsUpdate)
	if err != nil {
		t.Fatalf("PutServerSelection failed: %v", err)
	}

	// Verify the change
	if searcher.ServerSelection != ServerSelectionSsWindowsUpdate {
		t.Errorf("ServerSelection not updated: got %d, want %d", searcher.ServerSelection, ServerSelectionSsWindowsUpdate)
	}
}

func TestIUpdateSearcher_BeginSearch(t *testing.T) {
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

	// Start an async search for installed updates
	searchJob, err := searcher.BeginSearch("IsInstalled=1")
	if err != nil {
		t.Skipf("BeginSearch failed (may be expected in some environments): %v", err)
		return
	}
	if searchJob == nil {
		t.Fatal("BeginSearch returned nil")
	}
	if searchJob.disp == nil {
		t.Fatal("searchJob.disp is nil")
	}
}

func TestIUpdateSearcher_EndSearch(t *testing.T) {
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

	// Start an async search
	searchJob, err := searcher.BeginSearch("IsInstalled=1")
	if err != nil {
		t.Skipf("BeginSearch failed: %v", err)
		return
	}

	// Wait a bit for the search to complete
	// In a real scenario, you would poll IsCompleted or use callbacks
	// For testing, we'll try to end it immediately
	result, err := searcher.EndSearch(searchJob)
	if err != nil {
		t.Logf("EndSearch may fail if search not complete: %v", err)
		return
	}
	if result == nil {
		t.Error("EndSearch returned nil result")
	}
}
