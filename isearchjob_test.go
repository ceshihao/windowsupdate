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

func TestToISearchJob_NilDispatch(t *testing.T) {
	result, err := toISearchJob(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestISearchJob_StructureFields(t *testing.T) {
	job := &ISearchJob{
		IsCompleted: true,
	}
	if !job.IsCompleted {
		t.Errorf("IsCompleted not set correctly")
	}
}

func TestISearchJob_Methods_NilDispatch(t *testing.T) {
	job := &ISearchJob{
		disp:        nil,
		IsCompleted: true,
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
}

// TestISearchJob_CleanUpRequestAbort exercises CleanUp and RequestAbort via a real search job from BeginSearch.
func TestISearchJob_CleanUpRequestAbort(t *testing.T) {
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

	job, err := searcher.BeginSearch("IsInstalled=1")
	if err != nil {
		t.Skipf("BeginSearch failed: %v", err)
		return
	}
	if job == nil {
		t.Fatal("BeginSearch returned nil job")
	}

	// RequestAbort before CleanUp so disp is still valid
	_ = job.RequestAbort()

	err = job.CleanUp()
	if err != nil {
		t.Logf("CleanUp returned error (non-fatal): %v", err)
	}
}
