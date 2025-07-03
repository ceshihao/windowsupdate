package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIImageInformation(t *testing.T) {
	m := map[string]interface{}{
		"AltText": "alt",
		"Height":  int64(1),
		"Source":  "src",
		"Width":   int64(1),
	}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(m[prop]), nil
		}, nil,
		func() {
			_, _ = toIImageInformation(&ole.IDispatch{})
		},
	)
}

func TestToIImageInformation_Err(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return nil, nil // 模拟属性获取失败
		}, nil,
		func() {
			_, _ = toIImageInformation(&ole.IDispatch{})
		},
	)
}

func TestToIImageInformation_AltOnly(t *testing.T) {
	m := map[string]interface{}{
		"AltText": "alt",
	}
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(m[prop]), nil
		}, nil,
		func() {
			_, _ = toIImageInformation(&ole.IDispatch{})
		},
	)
}
