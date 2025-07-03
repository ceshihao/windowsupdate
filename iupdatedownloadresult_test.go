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
