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

func TestIUpdateDownloadResult_StructureFields(t *testing.T) {
	result := &IUpdateDownloadResult{
		HResult:    0,
		ResultCode: OperationResultCodeOrcSucceeded,
	}
	if result.HResult != 0 {
		t.Errorf("HResult not set correctly, got %d", result.HResult)
	}
	if result.ResultCode != OperationResultCodeOrcSucceeded {
		t.Errorf("ResultCode not set correctly")
	}
}

func TestIUpdateDownloadResult_ErrorCodes(t *testing.T) {
	// Test with various error codes
	testCases := []struct {
		name       string
		hresult    int32
		resultCode int32
	}{
		{"Success", 0, OperationResultCodeOrcSucceeded},
		{"Failed", -2147024891, OperationResultCodeOrcFailed},
		{"InProgress", 0, OperationResultCodeOrcInProgress},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := &IUpdateDownloadResult{
				HResult:    tc.hresult,
				ResultCode: tc.resultCode,
			}
			if result.HResult != tc.hresult {
				t.Errorf("HResult = %d, want %d", result.HResult, tc.hresult)
			}
			if result.ResultCode != tc.resultCode {
				t.Errorf("ResultCode = %d, want %d", result.ResultCode, tc.resultCode)
			}
		})
	}
}
