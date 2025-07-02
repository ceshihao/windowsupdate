package windowsupdate

import (
	"errors"
	"testing"

	"github.com/go-ole/go-ole"
)

func toIImageInformationTest(disp *ole.IDispatch) (*IImageInformation, error) {
	if disp == nil {
		return nil, nil
	}
	img := &IImageInformation{disp: disp}
	if v, err := getProperty(disp, "AltText"); err != nil {
		return nil, err
	} else {
		img.AltText = v.Value().(string)
	}
	if v, err := getProperty(disp, "Height"); err != nil {
		return nil, err
	} else {
		img.Height = v.Value().(int64)
	}
	if v, err := getProperty(disp, "Source"); err != nil {
		return nil, err
	} else {
		img.Source = v.Value().(string)
	}
	if v, err := getProperty(disp, "Width"); err != nil {
		return nil, err
	} else {
		img.Width = v.Value().(int64)
	}
	return img, nil
}

func TestToIImageInformation_AllSuccess(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		m := map[string]interface{}{"AltText": "alt", "Height": int64(1), "Source": "src", "Width": int64(2)}
		return &mockVariant{v: m[prop]}, nil
	}, func() {
		obj, err := toIImageInformationTest(&ole.IDispatch{})
		if err != nil || obj.AltText != "alt" || obj.Height != 1 || obj.Source != "src" || obj.Width != 2 {
			t.Errorf("unexpected: %+v, err=%v", obj, err)
		}
	})
}

func TestToIImageInformation_ErrorCases(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) { return nil, errors.New("err") }, func() {
		_, err := toIImageInformationTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "AltText" {
			return &mockVariant{v: "alt"}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIImageInformationTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "AltText" {
			return &mockVariant{v: "alt"}, nil
		}
		if prop == "Height" {
			return &mockVariant{v: int64(1)}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIImageInformationTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
	withGetProperty(func(_ *ole.IDispatch, prop string) (*mockVariant, error) {
		if prop == "AltText" {
			return &mockVariant{v: "alt"}, nil
		}
		if prop == "Height" {
			return &mockVariant{v: int64(1)}, nil
		}
		if prop == "Source" {
			return &mockVariant{v: "src"}, nil
		}
		return nil, errors.New("err")
	}, func() {
		_, err := toIImageInformationTest(&ole.IDispatch{})
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestToIImageInformation_NilInput(t *testing.T) {
	obj, err := toIImageInformationTest(nil)
	if err != nil || obj != nil {
		t.Errorf("unexpected: %+v, err=%v", obj, err)
	}
}
