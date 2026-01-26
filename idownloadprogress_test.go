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

import "testing"

func TestIDownloadProgress_StructureFields(t *testing.T) {
	progress := &IDownloadProgress{
		CurrentUpdateBytesDownloaded: 1024,
		CurrentUpdateBytesToDownload: 2048,
		CurrentUpdatePercentComplete: 50,
		PercentComplete:              75,
		TotalBytesDownloaded:         8192,
		TotalBytesToDownload:         16384,
		CurrentUpdateDownloadPhase:   1,
		CurrentUpdateIndex:           0,
	}
	if progress.CurrentUpdateBytesDownloaded != 1024 {
		t.Errorf("CurrentUpdateBytesDownloaded not set correctly")
	}
	if progress.PercentComplete != 75 {
		t.Errorf("PercentComplete not set correctly")
	}
	if progress.CurrentUpdateDownloadPhase != 1 {
		t.Errorf("CurrentUpdateDownloadPhase not set correctly")
	}
	if progress.CurrentUpdateIndex != 0 {
		t.Errorf("CurrentUpdateIndex not set correctly")
	}
}

func TestIDownloadProgress_TotalBytesDownloaded(t *testing.T) {
	progress := &IDownloadProgress{
		TotalBytesDownloaded: 4096,
	}
	if progress.TotalBytesDownloaded != 4096 {
		t.Errorf("TotalBytesDownloaded not working correctly, got %d", progress.TotalBytesDownloaded)
	}
}

func TestIDownloadProgress_TotalBytesToDownload(t *testing.T) {
	progress := &IDownloadProgress{
		TotalBytesToDownload: 8192,
	}
	if progress.TotalBytesToDownload != 8192 {
		t.Errorf("TotalBytesToDownload not working correctly, got %d", progress.TotalBytesToDownload)
	}
}

func TestIDownloadProgress_CurrentUpdateFields(t *testing.T) {
	progress := &IDownloadProgress{
		CurrentUpdateBytesDownloaded: 512,
		CurrentUpdateBytesToDownload: 1024,
		CurrentUpdatePercentComplete: 50,
	}
	if progress.CurrentUpdateBytesDownloaded != 512 {
		t.Errorf("CurrentUpdateBytesDownloaded = %d, want 512", progress.CurrentUpdateBytesDownloaded)
	}
	if progress.CurrentUpdateBytesToDownload != 1024 {
		t.Errorf("CurrentUpdateBytesToDownload = %d, want 1024", progress.CurrentUpdateBytesToDownload)
	}
	if progress.CurrentUpdatePercentComplete != 50 {
		t.Errorf("CurrentUpdatePercentComplete = %d, want 50", progress.CurrentUpdatePercentComplete)
	}
}

func TestToIDownloadProgress_NilDispatch(t *testing.T) {
	result, err := toIDownloadProgress(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}
