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
)

// Test helper functions and methods that don't require COM objects

func TestIUpdate_AllFieldsInitialization(t *testing.T) {
	now := time.Now()
	bundledUpdates := []*IUpdateIdentity{
		{RevisionNumber: 1, UpdateID: "update-1"},
		{RevisionNumber: 2, UpdateID: "update-2"},
	}
	categories := []*ICategory{
		{CategoryID: "cat-1", Name: "Category 1"},
	}
	downloadContents := []*IUpdateDownloadContent{
		{DownloadUrl: "https://example.com/update1.cab"},
	}
	identity := &IUpdateIdentity{RevisionNumber: 10, UpdateID: "main-update"}
	image := &IImageInformation{Width: 100, Height: 100}
	installBehavior := &IInstallationBehavior{Impact: InstallationImpactIiNormal}

	update := &IUpdate{
		AutoSelectOnWebSites:            true,
		BundledUpdates:                  bundledUpdates,
		CanRequireSource:                false,
		Categories:                      categories,
		Deadline:                        &now,
		DeltaCompressedContentAvailable: true,
		DeltaCompressedContentPreferred: false,
		DeploymentAction:                DeploymentActionDaInstallation,
		Description:                     "Test Description",
		DownloadContents:                downloadContents,
		DownloadPriority:                DownloadPriorityDpHigh,
		EulaAccepted:                    true,
		EulaText:                        "EULA Text",
		HandlerID:                       "handler-123",
		Identity:                        identity,
		Image:                           image,
		InstallationBehavior:            installBehavior,
		IsBeta:                          false,
		IsDownloaded:                    true,
		IsHidden:                        false,
		IsInstalled:                     false,
		IsMandatory:                     true,
		IsUninstallable:                 true,
		KBArticleIDs:                    []string{"KB123", "KB456"},
		Languages:                       []string{"en-US", "zh-CN"},
		LastDeploymentChangeTime:        &now,
		MaxDownloadSize:                 1024000,
		MinDownloadSize:                 512000,
		MoreInfoUrls:                    []string{"https://info1", "https://info2"},
		MsrcSeverity:                    "Critical",
		RecommendedCpuSpeed:             2000,
		RecommendedHardDiskSpace:        10000,
		RecommendedMemory:               4096,
		ReleaseNotes:                    "Release notes",
		SecurityBulletinIDs:             []string{"MS22-001"},
		SupersededUpdateIDs:             []string{"old-update-1"},
		SupportUrl:                      "https://support",
		Title:                           "Test Update",
		UninstallationBehavior:          installBehavior,
		UninstallationNotes:             "Uninstall notes",
		UninstallationSteps:             []string{"step1", "step2"},
		CveIDs:                          []string{"CVE-2022-0001"},
		IsPresent:                       true,
		RebootRequired:                  false,
		BrowseOnly:                      false,
		PerUser:                         true,
		AutoDownload:                    AutoDownloadModeAllowAutoDownload,
		AutoSelection:                   AutoSelectionModeAutoSelectIfDownloaded,
	}

	// Verify all fields
	if !update.AutoSelectOnWebSites {
		t.Error("AutoSelectOnWebSites not set")
	}
	if len(update.BundledUpdates) != 2 {
		t.Error("BundledUpdates not set correctly")
	}
	if len(update.Categories) != 1 {
		t.Error("Categories not set correctly")
	}
	if update.Deadline == nil {
		t.Error("Deadline should not be nil")
	}
	if len(update.DownloadContents) != 1 {
		t.Error("DownloadContents not set correctly")
	}
	if update.Identity == nil || update.Identity.RevisionNumber != 10 {
		t.Error("Identity not set correctly")
	}
	if len(update.KBArticleIDs) != 2 {
		t.Error("KBArticleIDs not set correctly")
	}
	if len(update.Languages) != 2 {
		t.Error("Languages not set correctly")
	}
	if update.MaxDownloadSize != 1024000 {
		t.Error("MaxDownloadSize not set correctly")
	}
	if len(update.SecurityBulletinIDs) != 1 {
		t.Error("SecurityBulletinIDs not set correctly")
	}
	if len(update.CveIDs) != 1 {
		t.Error("CveIDs not set correctly")
	}
	if !update.PerUser {
		t.Error("PerUser not set correctly")
	}
}

func TestIUpdateCollection_EmptySliceConversion(t *testing.T) {
	uc := &IUpdateCollection{
		Count:    0,
		ReadOnly: true,
	}

	// This tests the empty collection path in ToSlice
	// Even though we can't call ToSlice without a real dispatch,
	// we can test the struct itself
	if uc.Count != 0 {
		t.Error("Count should be 0")
	}
	if !uc.ReadOnly {
		t.Error("ReadOnly should be true")
	}
}

func TestEnumValues(t *testing.T) {
	// Test all enum constants are defined correctly
	tests := []struct {
		name     string
		value    int32
		expected int32
	}{
		// OperationResultCode
		{"OrcNotStarted", OperationResultCodeOrcNotStarted, 0},
		{"OrcInProgress", OperationResultCodeOrcInProgress, 1},
		{"OrcSucceeded", OperationResultCodeOrcSucceeded, 2},
		{"OrcSucceededWithErrors", OperationResultCodeOrcSucceededWithErrors, 3},
		{"OrcFailed", OperationResultCodeOrcFailed, 4},
		{"OrcAborted", OperationResultCodeOrcAborted, 5},

		// DeploymentAction
		{"DaNone", DeploymentActionDaNone, 0},
		{"DaInstallation", DeploymentActionDaInstallation, 2},

		// DownloadPriority
		{"DpLow", DownloadPriorityDpLow, 1},
		{"DpNormal", DownloadPriorityDpNormal, 2},
		{"DpHigh", DownloadPriorityDpHigh, 3},

		// ServerSelection
		{"SsDefault", ServerSelectionSsDefault, 0},
		{"SsWindowsUpdate", ServerSelectionSsWindowsUpdate, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.value, tt.expected)
			}
		})
	}
}

func TestComplexStructureNesting(t *testing.T) {
	// Test nested structures
	image := &IImageInformation{
		AltText: "Alt",
		Height:  50,
		Width:   100,
		Source:  "http://example.com/img.png",
	}

	behavior := &IInstallationBehavior{
		CanRequestUserInput:         true,
		Impact:                      InstallationImpactIiMinor,
		RebootBehavior:              InstallationRebootBehaviorIrbNeverReboots,
		RequiresNetworkConnectivity: false,
	}

	category := &ICategory{
		CategoryID:  "cat-123",
		Description: "Test Category",
		Name:        "Category Name",
		Order:       5,
		Type:        "Software",
		Image:       image,
		Children:    []*ICategory{},
	}

	if category.Image == nil {
		t.Error("Category Image should not be nil")
	}
	if category.Image.Width != 100 {
		t.Error("Nested image width incorrect")
	}
	if behavior.Impact != InstallationImpactIiMinor {
		t.Error("Installation behavior impact incorrect")
	}
}

func TestDownloadProgressFields(t *testing.T) {
	progress := &IDownloadProgress{
		CurrentUpdateBytesDownloaded: 1024,
		CurrentUpdateBytesToDownload: 2048,
		CurrentUpdateDownloadPhase:   DownloadPhaseDownloading,
		CurrentUpdateIndex:           0,
		CurrentUpdatePercentComplete: 50,
		PercentComplete:              25,
		TotalBytesDownloaded:         5120,
		TotalBytesToDownload:         20480,
	}

	if progress.CurrentUpdatePercentComplete != 50 {
		t.Error("CurrentUpdatePercentComplete incorrect")
	}
	if progress.PercentComplete != 25 {
		t.Error("PercentComplete incorrect")
	}
	if progress.CurrentUpdateDownloadPhase != DownloadPhaseDownloading {
		t.Error("Download phase incorrect")
	}

	// Test calculation
	downloaded := progress.TotalBytesDownloaded
	total := progress.TotalBytesToDownload
	if downloaded == 0 || total == 0 {
		t.Error("Bytes should not be zero")
	}
	percent := float64(downloaded) / float64(total) * 100
	if percent < 0 || percent > 100 {
		t.Error("Calculated percent out of range")
	}
}

func TestInstallationProgressFields(t *testing.T) {
	progress := &IInstallationProgress{
		CurrentUpdatePercentComplete: 75,
		PercentComplete:              50,
	}

	if progress.CurrentUpdatePercentComplete != 75 {
		t.Error("CurrentUpdatePercentComplete incorrect")
	}
	if progress.PercentComplete != 50 {
		t.Error("PercentComplete incorrect")
	}

	// Verify progress values are in valid range
	if progress.CurrentUpdatePercentComplete < 0 || progress.CurrentUpdatePercentComplete > 100 {
		t.Error("CurrentUpdatePercentComplete out of range")
	}
	if progress.PercentComplete < 0 || progress.PercentComplete > 100 {
		t.Error("PercentComplete out of range")
	}
}

func TestJobStructures(t *testing.T) {
	searchJob := &ISearchJob{
		AsyncState:  nil,
		IsCompleted: true,
	}
	if !searchJob.IsCompleted {
		t.Error("SearchJob IsCompleted should be true")
	}

	downloadJob := &IDownloadJob{
		AsyncState:  nil,
		IsCompleted: false,
	}
	if downloadJob.IsCompleted {
		t.Error("DownloadJob IsCompleted should be false")
	}

	installJob := &IInstallationJob{
		AsyncState:  nil,
		IsCompleted: true,
	}
	if !installJob.IsCompleted {
		t.Error("InstallationJob IsCompleted should be true")
	}
}

func TestResultStructures(t *testing.T) {
	downloadResult := &IDownloadResult{
		HResult:    0,
		ResultCode: OperationResultCodeOrcSucceeded,
	}
	if downloadResult.ResultCode != OperationResultCodeOrcSucceeded {
		t.Error("DownloadResult ResultCode incorrect")
	}

	installResult := &IInstallationResult{
		HResult:        0,
		RebootRequired: true,
		ResultCode:     OperationResultCodeOrcSucceeded,
	}
	if !installResult.RebootRequired {
		t.Error("InstallationResult RebootRequired should be true")
	}

	updateDownloadResult := &IUpdateDownloadResult{
		HResult:    0,
		ResultCode: OperationResultCodeOrcSucceeded,
	}
	if updateDownloadResult.HResult != 0 {
		t.Error("UpdateDownloadResult HResult incorrect")
	}

	updateInstallResult := &IUpdateInstallationResult{
		HResult:        0,
		RebootRequired: false,
		ResultCode:     OperationResultCodeOrcSucceeded,
	}
	if updateInstallResult.RebootRequired {
		t.Error("UpdateInstallationResult RebootRequired should be false")
	}
}

func TestSearcherAndInstallerStructures(t *testing.T) {
	searcher := &IUpdateSearcher{
		CanAutomaticallyUpgradeService:      false,
		ClientApplicationID:                 "test-searcher",
		IncludePotentiallySupersededUpdates: true,
		Online:                              true,
		ServerSelection:                     ServerSelectionSsManagedServer,
		ServiceID:                           "service-456",
	}
	if searcher.ServerSelection != ServerSelectionSsManagedServer {
		t.Error("Searcher ServerSelection incorrect")
	}

	downloader := &IUpdateDownloader{
		ClientApplicationID: "test-downloader",
		IsForced:            true,
		Priority:            DownloadPriorityDpHigh,
	}
	if !downloader.IsForced {
		t.Error("Downloader IsForced should be true")
	}
	if downloader.Priority != DownloadPriorityDpHigh {
		t.Error("Downloader Priority incorrect")
	}

	installer := &IUpdateInstaller{
		AllowSourcePrompts:  false,
		ClientApplicationID: "test-installer",
		ForceQuiet:          true,
		IsForced:            false,
	}
	if !installer.ForceQuiet {
		t.Error("Installer ForceQuiet should be true")
	}
}

func TestWebProxyConfiguration(t *testing.T) {
	proxy := &IWebProxy{
		Address:            "http://proxy.example.com:8080",
		AutoDetect:         false,
		BypassList:         []string{"localhost", "127.0.0.1", "*.local"},
		BypassProxyOnLocal: true,
		ReadOnly:           false,
		UserName:           "proxyuser",
	}

	if proxy.Address != "http://proxy.example.com:8080" {
		t.Error("Proxy Address incorrect")
	}
	if proxy.AutoDetect {
		t.Error("Proxy AutoDetect should be false")
	}
	if len(proxy.BypassList) != 3 {
		t.Error("Proxy BypassList length incorrect")
	}
	if !proxy.BypassProxyOnLocal {
		t.Error("Proxy BypassProxyOnLocal should be true")
	}
}

func TestUpdateHistoryEntryComplete(t *testing.T) {
	now := time.Now()
	identity := &IUpdateIdentity{
		RevisionNumber: 5,
		UpdateID:       "history-update-id",
	}

	entry := &IUpdateHistoryEntry{
		ClientApplicationID: "history-client",
		Date:                &now,
		Description:         "History Description",
		HResult:             0,
		Operation:           UpdateOperationUoUninstallation,
		ResultCode:          OperationResultCodeOrcFailed,
		ServerSelection:     ServerSelectionSsManagedServer,
		ServiceID:           "history-service",
		SupportUrl:          "https://support.example.com",
		Title:               "History Title",
		UninstallationNotes: "Uninstall notes",
		UninstallationSteps: []string{"stop service", "remove files", "clean registry"},
		UnmappedResultCode:  123,
		UpdateIdentity:      identity,
	}

	if entry.Operation != UpdateOperationUoUninstallation {
		t.Error("History Operation incorrect")
	}
	if entry.ResultCode != OperationResultCodeOrcFailed {
		t.Error("History ResultCode incorrect")
	}
	if len(entry.UninstallationSteps) != 3 {
		t.Error("History UninstallationSteps length incorrect")
	}
	if entry.UpdateIdentity == nil || entry.UpdateIdentity.RevisionNumber != 5 {
		t.Error("History UpdateIdentity incorrect")
	}
}
