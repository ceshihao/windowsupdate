package windowsupdate

import (
	"errors"
	"testing"

	"github.com/go-ole/go-ole"
)

func toIWebProxyTest(disp *ole.IDispatch) (*IWebProxy, error) {
	if disp == nil {
		return nil, nil
	}
	proxy := &IWebProxy{disp: disp}
	if v, err := getProperty(disp, "Address"); err != nil {
		return nil, err
	} else {
		proxy.Address = getMockValue(v).(string)
	}
	if v, err := getProperty(disp, "AutoDetect"); err != nil {
		return nil, err
	} else {
		proxy.AutoDetect = getMockValue(v).(bool)
	}
	if v, err := getProperty(disp, "BypassList"); err != nil {
		return nil, err
	} else {
		proxy.BypassList = getMockValue(v).([]string)
	}
	if v, err := getProperty(disp, "BypassProxyOnLocal"); err != nil {
		return nil, err
	} else {
		proxy.BypassProxyOnLocal = getMockValue(v).(bool)
	}
	if v, err := getProperty(disp, "ReadOnly"); err != nil {
		return nil, err
	} else {
		proxy.ReadOnly = getMockValue(v).(bool)
	}
	if v, err := getProperty(disp, "UserName"); err != nil {
		return nil, err
	} else {
		proxy.UserName = getMockValue(v).(string)
	}
	return proxy, nil
}

func TestToIWebProxy_AllSuccess(t *testing.T) {
	m := map[string]interface{}{
		"Address":            "1.2.3.4:8080",
		"AutoDetect":         false,
		"BypassList":         []string{"localhost", "127.0.0.1"},
		"BypassProxyOnLocal": true,
		"ReadOnly":           true,
		"UserName":           "user1",
	}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(m[prop]), nil
		}, nil,
		func() {
			obj, err := toIWebProxyTest(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if obj.Address != "1.2.3.4:8080" || obj.AutoDetect != false || len(obj.BypassList) != 2 || obj.BypassList[0] != "localhost" || !obj.BypassProxyOnLocal || !obj.ReadOnly || obj.UserName != "user1" {
				t.Errorf("unexpected struct values: %+v", obj)
			}
		},
	)
}

func TestToIWebProxy_ErrorCases(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return nil, errors.New("mock error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for Address")
			}
		},
	)

	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Address" {
				return fakeVariant("1.2.3.4:8080"), nil
			}
			return nil, errors.New("auto error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for AutoDetect")
			}
		},
	)

	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Address" {
				return fakeVariant("1.2.3.4:8080"), nil
			}
			if prop == "AutoDetect" {
				return fakeVariant(false), nil
			}
			return nil, errors.New("bypasslist error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for BypassList")
			}
		},
	)

	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Address" {
				return fakeVariant("1.2.3.4:8080"), nil
			}
			if prop == "AutoDetect" {
				return fakeVariant(false), nil
			}
			if prop == "BypassList" {
				return fakeVariant([]string{"localhost", "127.0.0.1"}), nil
			}
			return nil, errors.New("bypasslocal error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for BypassProxyOnLocal")
			}
		},
	)

	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Address" {
				return fakeVariant("1.2.3.4:8080"), nil
			}
			if prop == "AutoDetect" {
				return fakeVariant(false), nil
			}
			if prop == "BypassList" {
				return fakeVariant([]string{"localhost", "127.0.0.1"}), nil
			}
			if prop == "BypassProxyOnLocal" {
				return fakeVariant(true), nil
			}
			return nil, errors.New("readonly error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for ReadOnly")
			}
		},
	)

	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Address" {
				return fakeVariant("1.2.3.4:8080"), nil
			}
			if prop == "AutoDetect" {
				return fakeVariant(false), nil
			}
			if prop == "BypassList" {
				return fakeVariant([]string{"localhost", "127.0.0.1"}), nil
			}
			if prop == "BypassProxyOnLocal" {
				return fakeVariant(true), nil
			}
			if prop == "ReadOnly" {
				return fakeVariant(true), nil
			}
			return nil, errors.New("username error")
		}, nil,
		func() {
			_, err := toIWebProxyTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for UserName")
			}
		},
	)
}

func TestToIWebProxy_NilInput(t *testing.T) {
	obj, err := toIWebProxyTest(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if obj != nil {
		t.Errorf("expected nil, got: %+v", obj)
	}
}
