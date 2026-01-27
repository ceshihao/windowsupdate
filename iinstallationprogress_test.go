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

func TestToIInstallationProgress_NilDispatch(t *testing.T) {
	result, err := toIInstallationProgress(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIInstallationProgress_StructureFields(t *testing.T) {
	progress := &IInstallationProgress{
		CurrentUpdatePercentComplete: 60,
		PercentComplete:              80,
		CurrentUpdateIndex:           1,
	}
	if progress.CurrentUpdatePercentComplete != 60 {
		t.Errorf("CurrentUpdatePercentComplete not set correctly")
	}
	if progress.PercentComplete != 80 {
		t.Errorf("PercentComplete not set correctly")
	}
	if progress.CurrentUpdateIndex != 1 {
		t.Errorf("CurrentUpdateIndex not set correctly")
	}
}

func TestIInstallationProgress_PercentComplete(t *testing.T) {
	progress := &IInstallationProgress{
		PercentComplete: 45,
	}
	if progress.PercentComplete != 45 {
		t.Errorf("PercentComplete = %d, want 45", progress.PercentComplete)
	}
}

func TestIInstallationProgress_GetUpdateResult_NilDispatch(t *testing.T) {
	defer func() {
		_ = recover()
	}()

	progress := &IInstallationProgress{
		disp: nil,
	}
	result, err := progress.GetUpdateResult(0)
	if err == nil && result != nil {
		t.Errorf("expected error or panic for nil dispatch, got result=%v, err=%v", result, err)
	}
}

