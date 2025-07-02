package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestIUpdateCollection_Add_Clear_Count_Item_Remove_GetUpdates(t *testing.T) {
	coll := &IUpdateCollection{disp: &ole.IDispatch{}}
	update := &IUpdate{disp: &ole.IDispatch{}}
	withGetProperty(func(_ *ole.IDispatch, _ string) (*mockVariant, error) {
		return &mockVariant{v: 0}, nil
	}, func() {
		_ = coll.Add(update)
		_ = coll.Clear()
		_, _ = coll.Count()
		_, _ = coll.Item(0)
		_ = coll.Remove(0)
		_, _ = coll.GetUpdates()
	})
}
