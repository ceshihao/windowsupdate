package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestIUpdateCollection_Add_Clear_Count_Item_Remove_GetUpdates(t *testing.T) {
	coll := &IUpdateCollection{disp: &ole.IDispatch{}}
	update := &IUpdate{disp: &ole.IDispatch{}}
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) {
			return fakeVariant(0), nil
		}, nil,
		func() {
			_ = coll.Add(update)
			_ = coll.Clear()
			_, _ = coll.Count()
			_, _ = coll.Item(0)
			_ = coll.Remove(0)
			_, _ = coll.GetUpdates()
		},
	)
}
