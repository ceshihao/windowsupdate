package windowsupdate

import (
	"errors"
	"testing"

	"github.com/go-ole/go-ole"
)

func toIUpdateExceptionTest(disp *ole.IDispatch) (*IUpdateException, error) {
	if disp == nil {
		return nil, nil
	}
	e := &IUpdateException{disp: disp}
	if v, err := getProperty(disp, "Context"); err != nil {
		return nil, err
	} else {
		e.Context = v.Value().(int32)
	}
	if v, err := getProperty(disp, "HResult"); err != nil {
		return nil, err
	} else {
		e.HResult = v.Value().(int64)
	}
	if v, err := getProperty(disp, "Message"); err != nil {
		return nil, err
	} else {
		e.Message = v.Value().(string)
	}
	return e, nil
}

func TestToIUpdateException_AllSuccess(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		m := map[string]interface{}{"Context": int32(2), "HResult": int64(123), "Message": "msg"}
		return &mockVariant{v: m[prop]}, nil
	}, func() {
		obj, err := toIUpdateExceptionTest(&ole.IDispatch{})
		if err != nil || obj.Context != int32(2) || obj.HResult != int64(123) || obj.Message != "msg" {
			t.Errorf("unexpected: %+v, err=%v", obj, err)
		}
	})
}

func TestToIUpdateException_ErrorCases(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) { return nil, errors.New("err") }, func() {
		_, err := toIUpdateExceptionTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Context" {
			return &mockVariant{v: int32(2)}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIUpdateExceptionTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Context" {
			return &mockVariant{v: int32(2)}, nil
		}
		if prop == "HResult" {
			return &mockVariant{v: int64(123)}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIUpdateExceptionTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestToIUpdateException_NilInput(t *testing.T) {
	obj, err := toIUpdateExceptionTest(nil)
	if err != nil || obj != nil {
		t.Errorf("unexpected: %+v, err=%v", obj, err)
	}
}
