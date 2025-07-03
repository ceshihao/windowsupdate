package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIUpdateSearcher(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(0), nil }, nil,
		func() {
			_, _ = toIUpdateSearcher(&ole.IDispatch{})
		},
	)
}

func TestIUpdateSearcher_Search(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.Search("")
		},
	)
}

func TestIUpdateSearcher_QueryHistory(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.QueryHistory(0, 1)
		},
	)
}

func TestIUpdateSearcher_GetTotalHistoryCount(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.GetTotalHistoryCount()
		},
	)
}

func TestIUpdateSearcher_QueryHistoryAll(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.QueryHistoryAll()
		},
	)
}
