package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

// func fakeVariant(val interface{}) *ole.VARIANT {
// 	v := &ole.VARIANT{}
// 	// 这里只做空壳，实际用不到Value
// 	return v
// }

func TestToIUpdates(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(0), nil
		}, nil,
		func() {
			_, _ = toIUpdates(&ole.IDispatch{})
		},
	)
}

func TestToIUpdatesIdentities(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(0), nil
		}, nil,
		func() {
			_, _ = toIUpdatesIdentities(&ole.IDispatch{})
		},
	)
}

func TestToIUpdate(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(0), nil
		}, nil,
		func() {
			_, _ = toIUpdate(&ole.IDispatch{})
		},
	)
}

func TestToIUpdateCollection(t *testing.T) {
	_, _ = toIUpdateCollection([]*IUpdate{})
}

func TestAcceptEula(t *testing.T) {
	u := &IUpdate{disp: &ole.IDispatch{}}
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
		return fakeVariant(nil), nil
	}, func() {
		_ = u.AcceptEula()
	})
}
