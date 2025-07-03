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

	"github.com/go-ole/go-ole"
)

func toIUpdateDownloadResultTest(disp *ole.IDispatch) (*IUpdateDownloadResult, error) {
	if disp == nil {
		return nil, nil
	}
	r := &IUpdateDownloadResult{disp: disp}
	if v, err := getProperty(disp, "ResultCode"); err != nil {
		return nil, err
	} else {
		r.ResultCode = getMockValue(v).(int32)
	}
	if v, err := getProperty(disp, "HResult"); err != nil {
		return nil, err
	} else {
		r.HResult = getMockValue(v).(int32)
	}
	return r, nil
}

func TestToIUpdateDownloadResult_AllSuccess(t *testing.T) {
	m := map[string]interface{}{"ResultCode": int32(1), "HResult": int32(2)}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(m[prop]), nil
		}, nil,
		func() {
			obj, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
			if err != nil || obj.ResultCode != 1 || obj.HResult != 2 {
				t.Errorf("unexpected: %+v, err=%v", obj, err)
			}
		},
	)
}

func TestToIUpdateDownloadResult_ErrorCases(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return nil, errors.New("err")
		}, nil,
		func() {
			_, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error")
			}
		},
	)
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "ResultCode" {
				return fakeVariant(int32(1)), nil
			}
			return nil, errors.New("err")
		}, nil,
		func() {
			_, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error")
			}
		},
	)
}

func TestToIUpdateDownloadResult_NilInput(t *testing.T) {
	obj, err := toIUpdateDownloadResultTest(nil)
	if err != nil || obj != nil {
		t.Errorf("unexpected: %+v, err=%v", obj, err)
	}
}
