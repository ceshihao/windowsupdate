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
