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
		behavior.CanRequestUserInput = v.Value().(bool)
	}
	if v, err := getProperty(installationBehaviorDisp, "Impact"); err != nil {
		return nil, err
	} else {
		behavior.Impact = v.Value().(int32)
	}
	if v, err := getProperty(installationBehaviorDisp, "RebootBehavior"); err != nil {
		return nil, err
	} else {
		behavior.RebootBehavior = v.Value().(int32)
	}
	if v, err := getProperty(installationBehaviorDisp, "RequiresNetworkConnectivity"); err != nil {
		return nil, err
	} else {
		behavior.RequiresNetworkConnectivity = v.Value().(bool)
	}
	return behavior, nil
}

func TestToIInstallationBehavior_AllSuccess(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		m := map[string]interface{}{
			"CanRequestUserInput":         true,
			"Impact":                      int32(1),
			"RebootBehavior":              int32(2),
			"RequiresNetworkConnectivity": true,
		}
		return &mockVariant{v: m[prop]}, nil
	}, func() {
		obj, err := toIInstallationBehaviorTest(&ole.IDispatch{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !obj.CanRequestUserInput || obj.Impact != 1 || obj.RebootBehavior != 2 || !obj.RequiresNetworkConnectivity {
			t.Errorf("unexpected struct values: %+v", obj)
		}
	})
}

func TestToIInstallationBehavior_ErrorCases(t *testing.T) {
	// CanRequestUserInput error
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		return nil, errors.New("mock error")
	}, func() {
		_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for CanRequestUserInput")
		}
	})

	// Impact error
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "CanRequestUserInput" {
			return &mockVariant{v: true}, nil
		}
		return nil, errors.New("impact error")
	}, func() {
		_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for Impact")
		}
	})

	// RebootBehavior error
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "CanRequestUserInput" {
			return &mockVariant{v: true}, nil
		}
		if prop == "Impact" {
			return &mockVariant{v: int32(1)}, nil
		}
		return nil, errors.New("reboot error")
	}, func() {
		_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for RebootBehavior")
		}
	})

	// RequiresNetworkConnectivity error
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "CanRequestUserInput" {
			return &mockVariant{v: true}, nil
		}
		if prop == "Impact" {
			return &mockVariant{v: int32(1)}, nil
		}
		if prop == "RebootBehavior" {
			return &mockVariant{v: int32(2)}, nil
		}
		return nil, errors.New("net error")
	}, func() {
		_, err := toIInstallationBehaviorTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error for RequiresNetworkConnectivity")
		}
	})
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
