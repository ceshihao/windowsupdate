package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIUpdates(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) {
		return &mockVariant{v: 0}, nil
	}, func() {
		_, _ = toIUpdates(&ole.IDispatch{})
	})
}

func TestToIUpdatesIdentities(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) {
		return &mockVariant{v: 0}, nil
	}, func() {
		_, _ = toIUpdatesIdentities(&ole.IDispatch{})
	})
}

func TestToIUpdate(t *testing.T) {
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) {
		return &mockVariant{v: 0}, nil
	}, func() {
		_, _ = toIUpdate(&ole.IDispatch{})
	})
}

func TestToIUpdateCollection(t *testing.T) {
	_, _ = toIUpdateCollection([]*IUpdate{})
}

func TestAcceptEula(t *testing.T) {
	u := &IUpdate{disp: &ole.IDispatch{}}
	_ = u.AcceptEula()
}
