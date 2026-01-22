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
	"errors"
	"testing"
	"time"

	"github.com/go-ole/go-ole"
)

func TestToIDispatchErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toIDispatchErr(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toIDispatchErr(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != nil {
			t.Errorf("expected nil result for nil value variant, got %v", result)
		}
	})
}

func TestToInt64Err(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toInt64Err(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != 0 {
			t.Errorf("expected 0 result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toInt64Err(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestToInt32Err(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toInt32Err(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != 0 {
			t.Errorf("expected 0 result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toInt32Err(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestToFloat64Err(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toFloat64Err(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != 0 {
			t.Errorf("expected 0 result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toFloat64Err(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestToFloat32Err(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toFloat32Err(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != 0 {
			t.Errorf("expected 0 result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toFloat32Err(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestToStringErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toStringErr(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != "" {
			t.Errorf("expected empty string result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toStringErr(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "" {
			t.Errorf("expected empty string for nil value variant, got %v", result)
		}
	})
}

func TestToBoolErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toBoolErr(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != false {
			t.Errorf("expected false result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toBoolErr(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != false {
			t.Errorf("expected false for nil value variant, got %v", result)
		}
	})
}

func TestToTimeErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := toTimeErr(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
	})

	t.Run("WithNilVariant", func(t *testing.T) {
		v := &ole.VARIANT{}
		result, err := toTimeErr(v, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != nil {
			t.Errorf("expected nil for nil value variant, got %v", result)
		}
	})
}

func TestVariantToIDispatch(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToIDispatch(v)
		if result != nil {
			t.Errorf("expected nil for nil value variant, got %v", result)
		}
	})
}

func TestVariantToInt64(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToInt64(v)
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestVariantToInt32(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToInt32(v)
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestVariantToFloat64(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToFloat64(v)
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestVariantToFloat32(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToFloat32(v)
		if result != 0 {
			t.Errorf("expected 0 for nil value variant, got %v", result)
		}
	})
}

func TestVariantToString(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToString(v)
		if result != "" {
			t.Errorf("expected empty string for nil value variant, got %v", result)
		}
	})
}

func TestVariantToBool(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToBool(v)
		if result != false {
			t.Errorf("expected false for nil value variant, got %v", result)
		}
	})
}

func TestVariantToTime(t *testing.T) {
	t.Run("WithNilValue", func(t *testing.T) {
		v := &ole.VARIANT{}
		result := variantToTime(v)
		if result != nil {
			t.Errorf("expected nil for nil value variant, got %v", result)
		}
	})
}

func TestIStringCollectionToStringArrayErr(t *testing.T) {
	t.Run("WithError", func(t *testing.T) {
		expectedErr := errors.New("test error")
		result, err := iStringCollectionToStringArrayErr(nil, expectedErr)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
	})

	t.Run("WithNilDispatch", func(t *testing.T) {
		result, err := iStringCollectionToStringArrayErr(nil, nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != nil {
			t.Errorf("expected nil for nil dispatch, got %v", result)
		}
	})
}

func TestToIUpdateInstallationResult_NilDispatch(t *testing.T) {
	result, err := toIUpdateInstallationResult(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestToIStringCollection_NilDispatch(t *testing.T) {
	result, err := toIStringCollection(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

// TestIStringCollection_ToSlice_EmptyCollection tests ToSlice with zero count
func TestIStringCollection_ToSlice_EmptyCollection(t *testing.T) {
	sc := &IStringCollection{
		Count: 0,
	}
	result, err := sc.ToSlice()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

// TestStructureFields verifies that struct fields are correctly defined
func TestStructureFields(t *testing.T) {
	t.Run("IUpdateInstallationResult", func(t *testing.T) {
		r := &IUpdateInstallationResult{
			HResult:        -1,
			RebootRequired: true,
			ResultCode:     OperationResultCodeOrcSucceeded,
		}
		if r.HResult != -1 {
			t.Errorf("HResult not set correctly")
		}
		if !r.RebootRequired {
			t.Errorf("RebootRequired not set correctly")
		}
		if r.ResultCode != OperationResultCodeOrcSucceeded {
			t.Errorf("ResultCode not set correctly")
		}
	})

	t.Run("IUpdateIdentity", func(t *testing.T) {
		r := &IUpdateIdentity{
			RevisionNumber: 123,
			UpdateID:       "test-update-id",
		}
		if r.RevisionNumber != 123 {
			t.Errorf("RevisionNumber not set correctly")
		}
		if r.UpdateID != "test-update-id" {
			t.Errorf("UpdateID not set correctly")
		}
	})

	t.Run("IDownloadResult", func(t *testing.T) {
		r := &IDownloadResult{
			HResult:    0,
			ResultCode: OperationResultCodeOrcSucceeded,
		}
		if r.HResult != 0 {
			t.Errorf("HResult not set correctly")
		}
		if r.ResultCode != OperationResultCodeOrcSucceeded {
			t.Errorf("ResultCode not set correctly")
		}
	})

	t.Run("IInstallationResult", func(t *testing.T) {
		r := &IInstallationResult{
			HResult:        0,
			RebootRequired: false,
			ResultCode:     OperationResultCodeOrcSucceeded,
		}
		if r.HResult != 0 {
			t.Errorf("HResult not set correctly")
		}
		if r.RebootRequired {
			t.Errorf("RebootRequired not set correctly")
		}
		if r.ResultCode != OperationResultCodeOrcSucceeded {
			t.Errorf("ResultCode not set correctly")
		}
	})

	t.Run("IUpdateException", func(t *testing.T) {
		r := &IUpdateException{
			Context: UpdateExceptionContextUecGeneral,
			HResult: 0x80070005,
			Message: "Access denied",
		}
		if r.Context != UpdateExceptionContextUecGeneral {
			t.Errorf("Context not set correctly")
		}
		if r.HResult != 0x80070005 {
			t.Errorf("HResult not set correctly")
		}
		if r.Message != "Access denied" {
			t.Errorf("Message not set correctly")
		}
	})

	t.Run("IUpdateHistoryEntry", func(t *testing.T) {
		now := time.Now()
		r := &IUpdateHistoryEntry{
			ClientApplicationID: "test-app",
			Date:                &now,
			Description:         "Test update",
			HResult:             0,
			Operation:           UpdateOperationUoInstallation,
			ResultCode:          OperationResultCodeOrcSucceeded,
			ServerSelection:     ServerSelectionSsWindowsUpdate,
			ServiceID:           "service-id",
			SupportUrl:          "https://example.com",
			Title:               "Test Update Title",
			UninstallationNotes: "Notes",
			UninstallationSteps: []string{"step1", "step2"},
			UnmappedResultCode:  0,
		}
		if r.ClientApplicationID != "test-app" {
			t.Errorf("ClientApplicationID not set correctly")
		}
		if r.Date == nil {
			t.Errorf("Date not set correctly")
		}
		if r.Title != "Test Update Title" {
			t.Errorf("Title not set correctly")
		}
	})

	t.Run("ISearchResult", func(t *testing.T) {
		r := &ISearchResult{
			ResultCode:     OperationResultCodeOrcSucceeded,
			RootCategories: nil,
			Updates:        nil,
			Warnings:       nil,
		}
		if r.ResultCode != OperationResultCodeOrcSucceeded {
			t.Errorf("ResultCode not set correctly")
		}
	})

	t.Run("IWebProxy", func(t *testing.T) {
		r := &IWebProxy{
			Address:            "http://proxy:8080",
			AutoDetect:         true,
			BypassList:         []string{"localhost", "127.0.0.1"},
			BypassProxyOnLocal: true,
			ReadOnly:           false,
			UserName:           "user",
		}
		if r.Address != "http://proxy:8080" {
			t.Errorf("Address not set correctly")
		}
		if !r.AutoDetect {
			t.Errorf("AutoDetect not set correctly")
		}
		if len(r.BypassList) != 2 {
			t.Errorf("BypassList not set correctly")
		}
	})

	t.Run("IUpdateSession", func(t *testing.T) {
		r := &IUpdateSession{
			ClientApplicationID: "test-client",
			ReadOnly:            true,
			WebProxy:            nil,
		}
		if r.ClientApplicationID != "test-client" {
			t.Errorf("ClientApplicationID not set correctly")
		}
		if !r.ReadOnly {
			t.Errorf("ReadOnly not set correctly")
		}
	})

	t.Run("IStringCollection", func(t *testing.T) {
		r := &IStringCollection{
			Count:    5,
			ReadOnly: true,
		}
		if r.Count != 5 {
			t.Errorf("Count not set correctly")
		}
		if !r.ReadOnly {
			t.Errorf("ReadOnly not set correctly")
		}
	})

	t.Run("IUpdateDownloadResult", func(t *testing.T) {
		r := &IUpdateDownloadResult{
			HResult:    0,
			ResultCode: OperationResultCodeOrcSucceeded,
		}
		if r.HResult != 0 {
			t.Errorf("HResult not set correctly")
		}
		if r.ResultCode != OperationResultCodeOrcSucceeded {
			t.Errorf("ResultCode not set correctly")
		}
	})
}
