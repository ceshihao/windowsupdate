/*
Copyright 2022 Zheng Dayu
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
