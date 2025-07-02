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
		r.ResultCode = v.Value().(int32)
	}
	if v, err := getProperty(disp, "HResult"); err != nil {
		return nil, err
	} else {
		r.HResult = v.Value().(int32)
	}
	return r, nil
}

func TestToIUpdateDownloadResult_AllSuccess(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		m := map[string]interface{}{"ResultCode": int32(1), "HResult": int32(2)}
		return &mockVariant{v: m[prop]}, nil
	}, func() {
		obj, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
		if err != nil || obj.ResultCode != 1 || obj.HResult != 2 {
			t.Errorf("unexpected: %+v, err=%v", obj, err)
		}
	})
}

func TestToIUpdateDownloadResult_ErrorCases(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) { return nil, errors.New("err") }, func() {
		_, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "ResultCode" {
			return &mockVariant{v: int32(1)}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIUpdateDownloadResultTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestToIUpdateDownloadResult_NilInput(t *testing.T) {
	obj, err := toIUpdateDownloadResultTest(nil)
	if err != nil || obj != nil {
		t.Errorf("unexpected: %+v, err=%v", obj, err)
	}
}
