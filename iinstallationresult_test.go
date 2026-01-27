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

func TestIInstallationResult_StructureFields(t *testing.T) {
	result := &IInstallationResult{
		HResult:        0,
		RebootRequired: true,
		ResultCode:     OperationResultCodeOrcSucceeded,
	}
	if result.HResult != 0 {
		t.Errorf("HResult not set correctly, got %d", result.HResult)
	}
	if !result.RebootRequired {
		t.Errorf("RebootRequired not set correctly, got %v", result.RebootRequired)
	}
	if result.ResultCode != OperationResultCodeOrcSucceeded {
		t.Errorf("ResultCode not set correctly")
	}
}

func TestIInstallationResult_RebootRequired(t *testing.T) {
	result := &IInstallationResult{
		RebootRequired: false,
	}
	if result.RebootRequired {
		t.Errorf("RebootRequired should be false")
	}
}

func TestIInstallationResult_ErrorScenarios(t *testing.T) {
	testCases := []struct {
		name           string
		hresult        int32
		resultCode     int32
		rebootRequired bool
	}{
		{"Success_NoReboot", 0, OperationResultCodeOrcSucceeded, false},
		{"Success_WithReboot", 0, OperationResultCodeOrcSucceeded, true},
		{"Failed", -2147024891, OperationResultCodeOrcFailed, false},
		{"Aborted", 0, OperationResultCodeOrcAborted, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := &IInstallationResult{
				HResult:        tc.hresult,
				ResultCode:     tc.resultCode,
				RebootRequired: tc.rebootRequired,
			}
			if result.HResult != tc.hresult {
				t.Errorf("HResult = %d, want %d", result.HResult, tc.hresult)
			}
			if result.ResultCode != tc.resultCode {
				t.Errorf("ResultCode = %d, want %d", result.ResultCode, tc.resultCode)
			}
			if result.RebootRequired != tc.rebootRequired {
				t.Errorf("RebootRequired = %v, want %v", result.RebootRequired, tc.rebootRequired)
			}
		})
	}
}

// Note: toIInstallationResult and GetUpdateResult require actual installation results
// from COM objects. These are tested indirectly when installations are performed
// in integration tests. Unit testing these requires real Windows Update operations
// which is not suitable for automated unit tests.
