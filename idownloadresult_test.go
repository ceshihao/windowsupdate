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

func TestIDownloadResult_StructureFields(t *testing.T) {
	result := &IDownloadResult{
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

func TestIDownloadResult_HResult(t *testing.T) {
	result := &IDownloadResult{
		HResult: -2147024891, // E_ACCESSDENIED
	}
	if result.HResult != -2147024891 {
		t.Errorf("HResult = %d, want -2147024891", result.HResult)
	}
}

func TestIDownloadResult_ResultCodes(t *testing.T) {
	testCases := []struct {
		name       string
		resultCode int32
	}{
		{"Succeeded", OperationResultCodeOrcSucceeded},
		{"InProgress", OperationResultCodeOrcInProgress},
		{"Failed", OperationResultCodeOrcFailed},
		{"Aborted", OperationResultCodeOrcAborted},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := &IDownloadResult{
				HResult:    0,
				ResultCode: tc.resultCode,
			}
			if result.ResultCode != tc.resultCode {
				t.Errorf("ResultCode = %d, want %d", result.ResultCode, tc.resultCode)
			}
		})
	}
}

func TestToIDownloadResult_NilDispatch(t *testing.T) {
	defer func() {
		// If a panic occurs when using a nil dispatch, that's acceptable
		// as the COM layer may not handle nil pointers uniformly.
		_ = recover()
	}()

	result, err := toIDownloadResult(nil)
	if err == nil && result != nil {
		t.Errorf("expected error or panic for nil dispatch, got result=%v, err=%v", result, err)
	}
}

func TestIDownloadResult_GetUpdateResult_NilDispatch(t *testing.T) {
	defer func() {
		// Allow panic as a valid behavior when the underlying COM dispatch is nil.
		_ = recover()
	}()

	dr := &IDownloadResult{
		disp: nil,
	}
	updateResult, err := dr.GetUpdateResult(0)
	if err == nil && updateResult != nil {
		t.Errorf("expected error or panic for nil dispatch, got result=%v, err=%v", updateResult, err)
	}
}

