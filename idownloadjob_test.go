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

func TestToIDownloadJob_NilDispatch(t *testing.T) {
	result, err := toIDownloadJob(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIDownloadJob_StructureFields(t *testing.T) {
	job := &IDownloadJob{
		IsCompleted: false,
	}
	if job.IsCompleted {
		t.Errorf("IsCompleted should be false")
	}
}

func TestIDownloadJob_Methods_NilDispatch(t *testing.T) {
	job := &IDownloadJob{
		disp:        nil,
		IsCompleted: false,
	}

	// CleanUp
	func() {
		defer func() { _ = recover() }()
		_ = job.CleanUp()
	}()

	// RequestAbort
	func() {
		defer func() { _ = recover() }()
		_ = job.RequestAbort()
	}()

	// GetProgress
	func() {
		defer func() { _ = recover() }()
		progress, err := job.GetProgress()
		if err == nil && progress != nil {
			t.Errorf("expected error or panic for nil dispatch, got progress=%v, err=%v", progress, err)
		}
	}()
}

// TestIDownloadJob_BeginDownloadAndMethods exercises toIDownloadJob, CleanUp, RequestAbort, GetProgress via real COM.
func TestIDownloadJob_BeginDownloadAndMethods(t *testing.T) {
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

	// BeginDownload with empty updates yields a job that completes immediately or very quickly
	job, err := downloader.BeginDownload([]*IUpdate{})
	if err != nil {
		t.Skipf("BeginDownload failed (may need WU service): %v", err)
		return
	}
	if job == nil {
		t.Fatal("BeginDownload returned nil job")
	}

	// GetProgress returns current progress (covers toIDownloadProgress path when used from job)
	progress, err := job.GetProgress()
	if err != nil {
		t.Logf("GetProgress returned error (non-fatal): %v", err)
	}
	if progress != nil {
		// If we got progress, calling GetUpdateResult(0) covers IDownloadProgress.GetUpdateResult
		_, _ = progress.GetUpdateResult(0)
	}

	// RequestAbort before CleanUp so disp is still valid
	_ = job.RequestAbort()

	// CleanUp releases resources; safe to call on completed job
	err = job.CleanUp()
	if err != nil {
		t.Logf("CleanUp returned error (non-fatal): %v", err)
	}
}
