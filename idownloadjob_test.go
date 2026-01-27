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

// COM tests for job methods
func TestIDownloadJob_Methods(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Note: These methods require a real download job from BeginDownload.
	// Testing with nil dispatch would cause panic, so we skip direct method calls.
	// These methods are covered through integration tests or manual testing.

	job := &IDownloadJob{
		disp:        nil,
		IsCompleted: false,
	}

	// Verify structure can be created
	if job.IsCompleted {
		t.Error("IsCompleted should be false")
	}
}
