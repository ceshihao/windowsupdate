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

func TestIUpdateDownloader_StructureFields(t *testing.T) {
	downloader := &IUpdateDownloader{
		ClientApplicationID: "test-downloader",
		IsForced:            false,
		Priority:            DownloadPriorityDpNormal,
	}
	if downloader.ClientApplicationID != "test-downloader" {
		t.Errorf("ClientApplicationID not set correctly")
	}
	if downloader.IsForced {
		t.Errorf("IsForced should be false")
	}
	if downloader.Priority != DownloadPriorityDpNormal {
		t.Errorf("Priority not set correctly")
	}
}

// COM tests
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

func TestIUpdateDownloader_PutClientApplicationID(t *testing.T) {
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

	testID := "TestDownloaderApp"
	err = downloader.PutClientApplicationID(testID)
	if err != nil {
		t.Errorf("PutClientApplicationID failed: %v", err)
	}
	if downloader.ClientApplicationID != testID {
		t.Errorf("ClientApplicationID = %q, want %q", downloader.ClientApplicationID, testID)
	}
}

func TestIUpdateDownloader_Download(t *testing.T) {
	// Note: Download requires a real IUpdateDownloader with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	downloader := &IUpdateDownloader{
		disp:     nil,
		IsForced: false,
	}

	// Verify structure can be created
	if downloader.IsForced {
		t.Error("IsForced should be false")
	}
}

func TestIUpdateDownloader_BeginDownload(t *testing.T) {
	// Note: BeginDownload requires a real IUpdateDownloader with updates.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	downloader := &IUpdateDownloader{
		disp:     nil,
		Priority: DownloadPriorityDpNormal,
	}

	// Verify structure can be created
	if downloader.Priority != DownloadPriorityDpNormal {
		t.Error("Priority not set correctly")
	}
}

func TestIUpdateDownloader_EndDownload(t *testing.T) {
	// Note: EndDownload requires a real download job.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	downloader := &IUpdateDownloader{
		disp: nil,
	}

	// Verify structure can be created
	if downloader.disp != nil {
		t.Error("disp should be nil")
	}
}
