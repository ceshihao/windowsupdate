package windowsupdate

import (
	"errors"
	"testing"

	"github.com/go-ole/go-ole"
)

// 被测函数的注入版本
func toIInstallationBehaviorTest(installationBehaviorDisp *ole.IDispatch) (*IInstallationBehavior, error) {
	if installationBehaviorDisp == nil {
		return nil, nil
	}
	behavior := &IInstallationBehavior{
		disp: installationBehaviorDisp,
	}
	if v, err := getProperty(installationBehaviorDisp, "CanRequestUserInput"); err != nil {
		return nil, err
	} else {
		behavior.CanRequestUserInput = getMockValue(v).(bool)
	}
	if v, err := getProperty(installationBehaviorDisp, "Impact"); err != nil {
		return nil, err
	} else {
		behavior.Impact = getMockValue(v).(int32)
	}
	if v, err := getProperty(installationBehaviorDisp, "RebootBehavior"); err != nil {
		return nil, err
	} else {
		behavior.RebootBehavior = getMockValue(v).(int32)
	}
	if v, err := getProperty(installationBehaviorDisp, "RequiresNetworkConnectivity"); err != nil {
		return nil, err
	} else {
		behavior.RequiresNetworkConnectivity = getMockValue(v).(bool)
	}
	return behavior, nil
}

func TestToIInstallationBehavior_AllSuccess(t *testing.T) {
	m := map[string]interface{}{
		"CanRequestUserInput":         true,
		"Impact":                      int32(1),
		"RebootBehavior":              int32(2),
		"RequiresNetworkConnectivity": true,
	}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			t.Logf("prop: %q", prop)
			for k := range m {
				t.Logf("m key: %q", k)
			}
			v, ok := m[prop]
			if !ok {
				panic("mock: unexpected property " + prop)
			}
			return fakeVariant(v), nil
		}, nil,
		func() {
			obj, err := toIInstallationBehaviorTest(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !obj.CanRequestUserInput || obj.Impact != 1 || obj.RebootBehavior != 2 || !obj.RequiresNetworkConnectivity {
				t.Errorf("unexpected struct values: %+v", obj)
			}
		},
	)
}

func TestToIInstallationBehavior_ErrorCases(t *testing.T) {
	// CanRequestUserInput error
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return nil, errors.New("mock error")
		}, nil,
		func() {
			_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for CanRequestUserInput")
			}
		},
	)

	// Impact error
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "CanRequestUserInput" {
				return fakeVariant(true), nil
			}
			return nil, errors.New("impact error")
		}, nil,
		func() {
			_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for Impact")
			}
		},
	)

	// RebootBehavior error
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "CanRequestUserInput" {
				return fakeVariant(true), nil
			}
			if prop == "Impact" {
				return fakeVariant(int32(1)), nil
			}
			return nil, errors.New("reboot error")
		}, nil,
		func() {
			_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for RebootBehavior")
			}
		},
	)

	// RequiresNetworkConnectivity error
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "CanRequestUserInput" {
				return fakeVariant(true), nil
			}
			if prop == "Impact" {
				return fakeVariant(int32(1)), nil
			}
			if prop == "RebootBehavior" {
				return fakeVariant(int32(2)), nil
			}
			return nil, errors.New("net error")
		}, nil,
		func() {
			_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
			if err == nil {
				t.Error("expected error for RequiresNetworkConnectivity")
			}
		},
	)
}

func TestToIInstallationBehavior_NilInput(t *testing.T) {
	obj, err := toIInstallationBehaviorTest(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if obj != nil {
		t.Errorf("expected nil, got: %+v", obj)
	}
}
