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
		e.Context = getMockValue(v).(int32)
	}
	if v, err := getProperty(disp, "HResult"); err != nil {
		return nil, err
	} else {
		e.HResult = getMockValue(v).(int64)
	}
	if v, err := getProperty(disp, "Message"); err != nil {
		return nil, err
	} else {
		e.Message = getMockValue(v).(string)
	}
	return e, nil
}

func TestToIUpdateException_AllSuccess(t *testing.T) {
	m := map[string]interface{}{"Context": int32(2), "HResult": int64(123), "Message": "msg"}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(m[prop]), nil
		}, nil,
		func() {
			obj, err := toIUpdateExceptionTest(&ole.IDispatch{})
			if err != nil || obj.Context != int32(2) || obj.HResult != int64(123) || obj.Message != "msg" {
				t.Errorf("unexpected: %+v, err=%v", obj, err)
			}
		},
	)
}

func TestToIUpdateException_ErrorCases(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return nil, errors.New("err")
		}, nil,
		func() {
			_, err := toIUpdateExceptionTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error")
			}
		},
	)
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Context" {
				return fakeVariant(int32(2)), nil
			}
			return nil, errors.New("err")
		}, nil,
		func() {
			_, err := toIUpdateExceptionTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error")
			}
		},
	)
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Context" {
				return fakeVariant(int32(2)), nil
			}
			if prop == "HResult" {
				return fakeVariant(int64(123)), nil
			}
			return nil, errors.New("err")
		}, nil,
		func() {
			_, err := toIUpdateExceptionTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error")
			}
		},
	)
}

func TestToIUpdateException_NilInput(t *testing.T) {
	obj, err := toIUpdateExceptionTest(nil)
	if err != nil || obj != nil {
		t.Errorf("unexpected: %+v, err=%v", obj, err)
	}
}
