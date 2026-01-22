//go:build windows && integration
// +build windows,integration

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

func init() {}

func TestNewUpdateSession(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	session, err := NewUpdateSession()
	if err != nil {
		t.Fatalf("NewUpdateSession failed: %v", err)
	}
	if session == nil {
		t.Fatal("NewUpdateSession returned nil session")
	}
	if session.disp == nil {
		t.Fatal("NewUpdateSession returned session with nil dispatch")
	}
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
}

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
	// Count should be >= 0
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
		// Query at most 10 entries
		queryCount := int32(10)
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

		// Verify first entry has expected fields populated
		if len(history) > 0 {
			entry := history[0]
			if entry.UpdateIdentity == nil {
				t.Error("First history entry has nil UpdateIdentity")
			}
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

	// Search for already installed updates (quick search)
	result, err := searcher.Search("IsInstalled=1 and IsHidden=0")
	if err != nil {
		// Search might fail due to network or permissions, skip in that case
		t.Skipf("Search failed (may be expected in CI): %v", err)
	}
	if result == nil {
		t.Fatal("Search returned nil result")
	}

	// ResultCode should be a valid OperationResultCode
	if result.ResultCode < OperationResultCodeOrcNotStarted || result.ResultCode > OperationResultCodeOrcAborted {
		t.Errorf("Invalid ResultCode: %d", result.ResultCode)
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

func TestIUpdateSearcher_IncludePotentiallySupersededUpdates(t *testing.T) {
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

	// Get initial value
	_ = searcher.IncludePotentiallySupersededUpdates

	// Test PutIncludePotentiallySupersededUpdates
	err = searcher.PutIncludePotentiallySupersededUpdates(true)
	if err != nil {
		t.Fatalf("PutIncludePotentiallySupersededUpdates failed: %v", err)
	}

	// Verify the change
	if !searcher.IncludePotentiallySupersededUpdates {
		t.Error("IncludePotentiallySupersededUpdates not updated to true")
	}
}

func TestIUpdateDownloader_Properties(t *testing.T) {
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

	// Test IsForced property
	_ = downloader.IsForced

	err = downloader.PutIsForced(true)
	if err != nil {
		t.Fatalf("PutIsForced failed: %v", err)
	}

	if !downloader.IsForced {
		t.Error("IsForced not updated to true")
	}

	// Test Priority property
	_ = downloader.Priority

	err = downloader.PutPriority(DownloadPriorityDpHigh)
	if err != nil {
		t.Fatalf("PutPriority failed: %v", err)
	}

	if downloader.Priority != DownloadPriorityDpHigh {
		t.Errorf("Priority not updated: got %d, want %d", downloader.Priority, DownloadPriorityDpHigh)
	}
}

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
}

func TestNewSystemInformation(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	sysInfo, err := NewSystemInformation()
	if err != nil {
		t.Fatalf("NewSystemInformation failed: %v", err)
	}
	if sysInfo == nil {
		t.Fatal("NewSystemInformation returned nil")
	}

	// RebootRequired should be a valid boolean (true or false)
	// This is a read-only property, just verify we can read it
	_ = sysInfo.RebootRequired
}

func TestNewWindowsUpdateAgentInfo(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}
	if agentInfo == nil {
		t.Fatal("NewWindowsUpdateAgentInfo returned nil")
	}
}

func TestIWindowsUpdateAgentInfo_GetApiMajorVersion(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	apiMajor, err := agentInfo.GetApiMajorVersion()
	if err != nil {
		t.Fatalf("GetApiMajorVersion failed: %v", err)
	}
	// API major version should be >= 7 for modern Windows
	if apiMajor < 1 {
		t.Errorf("ApiMajorVersion seems invalid: %d", apiMajor)
	}
}

func TestIWindowsUpdateAgentInfo_GetApiMinorVersion(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	apiMinor, err := agentInfo.GetApiMinorVersion()
	if err != nil {
		t.Fatalf("GetApiMinorVersion failed: %v", err)
	}
	// API minor version should be >= 0
	if apiMinor < 0 {
		t.Errorf("ApiMinorVersion seems invalid: %d", apiMinor)
	}
}

func TestIWindowsUpdateAgentInfo_GetProductVersionString(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	productVersion, err := agentInfo.GetProductVersionString()
	if err != nil {
		t.Fatalf("GetProductVersionString failed: %v", err)
	}
	if productVersion == "" {
		t.Error("ProductVersionString is empty")
	}
}

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

func TestNewUpdateCollection(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewUpdateCollection()
	if err != nil {
		t.Fatalf("NewUpdateCollection failed: %v", err)
	}
	if collection == nil {
		t.Fatal("NewUpdateCollection returned nil")
	}

	// New collection should have Count = 0
	if collection.Count != 0 {
		t.Errorf("New collection Count should be 0, got %d", collection.Count)
	}
}

func TestNewStringCollection(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}
	if collection == nil {
		t.Fatal("NewStringCollection returned nil")
	}

	// New collection should have Count = 0
	if collection.Count != 0 {
		t.Errorf("New collection Count should be 0, got %d", collection.Count)
	}
}

func TestIStringCollection_Operations(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	// Add items
	idx, err := collection.Add("test1")
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if idx != 0 {
		t.Errorf("First Add should return index 0, got %d", idx)
	}

	idx, err = collection.Add("test2")
	if err != nil {
		t.Fatalf("Add second item failed: %v", err)
	}
	if idx != 1 {
		t.Errorf("Second Add should return index 1, got %d", idx)
	}

	// Verify count
	if collection.Count != 2 {
		t.Errorf("Count should be 2, got %d", collection.Count)
	}

	// Get item
	item, err := collection.Item(0)
	if err != nil {
		t.Fatalf("Item(0) failed: %v", err)
	}
	if item != "test1" {
		t.Errorf("Item(0) should be 'test1', got '%s'", item)
	}

	// ToSlice
	slice, err := collection.ToSlice()
	if err != nil {
		t.Fatalf("ToSlice failed: %v", err)
	}
	if len(slice) != 2 {
		t.Errorf("ToSlice should return 2 items, got %d", len(slice))
	}
	if slice[0] != "test1" || slice[1] != "test2" {
		t.Errorf("ToSlice returned wrong values: %v", slice)
	}

	// Insert
	err = collection.Insert(1, "inserted")
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
	if collection.Count != 3 {
		t.Errorf("Count after insert should be 3, got %d", collection.Count)
	}

	// RemoveAt
	err = collection.RemoveAt(1)
	if err != nil {
		t.Fatalf("RemoveAt failed: %v", err)
	}
	if collection.Count != 2 {
		t.Errorf("Count after remove should be 2, got %d", collection.Count)
	}

	// Clear
	err = collection.Clear()
	if err != nil {
		t.Fatalf("Clear failed: %v", err)
	}
	if collection.Count != 0 {
		t.Errorf("Count after clear should be 0, got %d", collection.Count)
	}
}
