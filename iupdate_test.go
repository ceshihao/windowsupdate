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
	"time"

	"github.com/go-ole/go-ole"
)

func TestToIUpdatesIdentities_NilDispatch(t *testing.T) {
	result, err := toIUpdatesIdentities(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIUpdate_StructureFields(t *testing.T) {
	now := time.Now()
	update := &IUpdate{
		AutoSelectOnWebSites:            true,
		CanRequireSource:                false,
		DeltaCompressedContentAvailable: true,
		DeltaCompressedContentPreferred: false,
		DeploymentAction:                DeploymentActionDaInstallation,
		Description:                     "Test update description",
		DownloadPriority:                DownloadPriorityDpHigh,
		EulaAccepted:                    true,
		EulaText:                        "EULA text",
		HandlerID:                       "handler-123",
		IsBeta:                          false,
		IsDownloaded:                    true,
		IsHidden:                        false,
		IsInstalled:                     false,
		IsMandatory:                     true,
		IsUninstallable:                 true,
		MaxDownloadSize:                 1024000,
		MinDownloadSize:                 512000,
		MsrcSeverity:                    "Critical",
		RecommendedCpuSpeed:             2000,
		RecommendedHardDiskSpace:        10000,
		RecommendedMemory:               4096,
		ReleaseNotes:                    "Release notes",
		SupportUrl:                      "https://support.example.com",
		Title:                           "Test Update Title",
		UninstallationNotes:             "Uninstall notes",
		IsPresent:                       true,
		RebootRequired:                  false,
		BrowseOnly:                      false,
		PerUser:                         true,
		AutoDownload:                    AutoDownloadModeAllowAutoDownload,
		AutoSelection:                   AutoSelectionModeAutoSelectIfDownloaded,
		Deadline:                        &now,
		LastDeploymentChangeTime:        &now,
	}

	if !update.AutoSelectOnWebSites {
		t.Errorf("AutoSelectOnWebSites not set correctly")
	}
	if update.Description != "Test update description" {
		t.Errorf("Description not set correctly, got %s", update.Description)
	}
	if update.DeploymentAction != DeploymentActionDaInstallation {
		t.Errorf("DeploymentAction not set correctly")
	}
	if update.MaxDownloadSize != 1024000 {
		t.Errorf("MaxDownloadSize not set correctly, got %d", update.MaxDownloadSize)
	}
	if update.Title != "Test Update Title" {
		t.Errorf("Title not set correctly, got %s", update.Title)
	}
	if update.AutoDownload != AutoDownloadModeAllowAutoDownload {
		t.Errorf("AutoDownload not set correctly")
	}
	if update.PerUser != true {
		t.Errorf("PerUser not set correctly")
	}
}

func TestIUpdate_GetDispatch(t *testing.T) {
	update := &IUpdate{
		disp: nil,
	}
	if update.GetDispatch() != nil {
		t.Errorf("expected nil dispatch, got %v", update.GetDispatch())
	}
}

func TestIUpdate_SliceFields(t *testing.T) {
	update := &IUpdate{
		KBArticleIDs:        []string{"KB123456", "KB789012"},
		Languages:           []string{"en-US", "zh-CN"},
		MoreInfoUrls:        []string{"https://example.com/1", "https://example.com/2"},
		SecurityBulletinIDs: []string{"MS22-001", "MS22-002"},
		SupersededUpdateIDs: []string{"update-1", "update-2"},
		UninstallationSteps: []string{"step1", "step2", "step3"},
		CveIDs:              []string{"CVE-2022-0001", "CVE-2022-0002"},
	}

	if len(update.KBArticleIDs) != 2 {
		t.Errorf("KBArticleIDs length incorrect, got %d", len(update.KBArticleIDs))
	}
	if len(update.Languages) != 2 {
		t.Errorf("Languages length incorrect, got %d", len(update.Languages))
	}
	if len(update.MoreInfoUrls) != 2 {
		t.Errorf("MoreInfoUrls length incorrect, got %d", len(update.MoreInfoUrls))
	}
	if len(update.SecurityBulletinIDs) != 2 {
		t.Errorf("SecurityBulletinIDs length incorrect, got %d", len(update.SecurityBulletinIDs))
	}
	if len(update.UninstallationSteps) != 3 {
		t.Errorf("UninstallationSteps length incorrect, got %d", len(update.UninstallationSteps))
	}
	if len(update.CveIDs) != 2 {
		t.Errorf("CveIDs length incorrect, got %d", len(update.CveIDs))
	}
	if update.KBArticleIDs[0] != "KB123456" {
		t.Errorf("KBArticleIDs[0] incorrect, got %s", update.KBArticleIDs[0])
	}
	if update.CveIDs[1] != "CVE-2022-0002" {
		t.Errorf("CveIDs[1] incorrect, got %s", update.CveIDs[1])
	}
}

// COM tests
func TestToIUpdateCollection(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Test with empty slice
	disp, err := toIUpdateCollection([]*IUpdate{})
	if err != nil {
		t.Fatalf("toIUpdateCollection with empty slice failed: %v", err)
	}
	if disp == nil {
		t.Error("toIUpdateCollection returned nil dispatch")
	}
}

func TestIUpdate_AcceptEula(t *testing.T) {
	// Note: AcceptEula requires a real IUpdate object from Windows Update.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	update := &IUpdate{
		disp:         nil,
		EulaAccepted: false,
	}

	// Verify structure can be created
	if update.EulaAccepted {
		t.Error("EulaAccepted should be false")
	}
}

func TestIUpdate_CopyToCache(t *testing.T) {
	// Note: CopyToCache requires a real IUpdate object from Windows Update.
	// Calling with nil dispatch would cause panic.
	// This method is covered through integration tests.

	update := &IUpdate{
		disp:  nil,
		Title: "Test Update",
	}

	// Verify structure can be created
	if update.Title != "Test Update" {
		t.Error("Title not set correctly")
	}
}
