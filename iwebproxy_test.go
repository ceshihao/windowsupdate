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
		proxy.Address = v.Value().(string)
	}
	if v, err := getProperty(disp, "AutoDetect"); err != nil {
		return nil, err
	} else {
		proxy.AutoDetect = v.Value().(bool)
	}
	if v, err := getProperty(disp, "BypassList"); err != nil {
		return nil, err
	} else {
		proxy.BypassList = v.Value().([]string)
	}
	if v, err := getProperty(disp, "BypassProxyOnLocal"); err != nil {
		return nil, err
	} else {
		proxy.BypassProxyOnLocal = v.Value().(bool)
	}
	if v, err := getProperty(disp, "ReadOnly"); err != nil {
		return nil, err
	} else {
		proxy.ReadOnly = v.Value().(bool)
	}
	if v, err := getProperty(disp, "UserName"); err != nil {
		return nil, err
	} else {
		proxy.UserName = v.Value().(string)
	}
	return proxy, nil
}

func TestToIWebProxy_AllSuccess(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		m := map[string]interface{}{
			"Address":            "1.2.3.4:8080",
			"AutoDetect":         false,
			"BypassList":         []string{"localhost", "127.0.0.1"},
			"BypassProxyOnLocal": true,
			"ReadOnly":           true,
			"UserName":           "user1",
		}
		return &mockVariant{v: m[prop]}, nil
	}, func() {
		obj, err := toIWebProxyTest(&ole.IDispatch{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if obj.Address != "1.2.3.4:8080" || obj.AutoDetect != false || len(obj.BypassList) != 2 || obj.BypassList[0] != "localhost" || !obj.BypassProxyOnLocal || !obj.ReadOnly || obj.UserName != "user1" {
			t.Errorf("unexpected struct values: %+v", obj)
		}
	})
}

func TestToIWebProxy_ErrorCases(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		return nil, errors.New("mock error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for Address")
		}
	})

	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Address" {
			return &mockVariant{v: "1.2.3.4:8080"}, nil
		}
		return nil, errors.New("auto error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for AutoDetect")
		}
	})

	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Address" {
			return &mockVariant{v: "1.2.3.4:8080"}, nil
		}
		if prop == "AutoDetect" {
			return &mockVariant{v: false}, nil
		}
		return nil, errors.New("bypasslist error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for BypassList")
		}
	})

	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Address" {
			return &mockVariant{v: "1.2.3.4:8080"}, nil
		}
		if prop == "AutoDetect" {
			return &mockVariant{v: false}, nil
		}
		if prop == "BypassList" {
			return &mockVariant{v: []string{"localhost", "127.0.0.1"}}, nil
		}
		return nil, errors.New("bypasslocal error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for BypassProxyOnLocal")
		}
	})

	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Address" {
			return &mockVariant{v: "1.2.3.4:8080"}, nil
		}
		if prop == "AutoDetect" {
			return &mockVariant{v: false}, nil
		}
		if prop == "BypassList" {
			return &mockVariant{v: []string{"localhost", "127.0.0.1"}}, nil
		}
		if prop == "BypassProxyOnLocal" {
			return &mockVariant{v: true}, nil
		}
		return nil, errors.New("readonly error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for ReadOnly")
		}
	})

	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "Address" {
			return &mockVariant{v: "1.2.3.4:8080"}, nil
		}
		if prop == "AutoDetect" {
			return &mockVariant{v: false}, nil
		}
		if prop == "BypassList" {
			return &mockVariant{v: []string{"localhost", "127.0.0.1"}}, nil
		}
		if prop == "BypassProxyOnLocal" {
			return &mockVariant{v: true}, nil
		}
		if prop == "ReadOnly" {
			return &mockVariant{v: true}, nil
		}
		return nil, errors.New("username error")
	}, func() {
		_, err := toIWebProxyTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for UserName")
		}
	})
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
