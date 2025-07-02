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
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateCollection represents a collection of updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatecollection
type IUpdateCollection struct {
	disp *ole.IDispatch
}

// NewUpdateCollection creates a new IUpdateCollection interface.
func NewUpdateCollection() (*IUpdateCollection, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.UpdateCollection")
	if err != nil {
		return nil, err
	}
	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return &IUpdateCollection{disp: disp}, nil
}

// Add adds an update to the collection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatecollection-add
func (iUpdateCollection *IUpdateCollection) Add(update *IUpdate) error {
	_, err := oleutil.CallMethod(iUpdateCollection.disp, "Add", update.disp)
	return err
}

// Clear removes all updates from the collection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatecollection-clear
func (iUpdateCollection *IUpdateCollection) Clear() error {
	_, err := oleutil.CallMethod(iUpdateCollection.disp, "Clear")
	return err
}

// Count returns the number of updates in the collection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatecollection-get_count
func (iUpdateCollection *IUpdateCollection) Count() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateCollection.disp, "Count"))
}

// Item returns the update at the specified index.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatecollection-item
func (iUpdateCollection *IUpdateCollection) Item(index int32) (*IUpdate, error) {
	updateDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateCollection.disp, "Item", index))
	if err != nil {
		return nil, err
	}
	return toIUpdate(updateDisp)
}

// Remove removes an update from the collection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatecollection-remove
func (iUpdateCollection *IUpdateCollection) Remove(index int32) error {
	_, err := oleutil.CallMethod(iUpdateCollection.disp, "Remove", index)
	return err
}

// GetUpdates returns all updates in the collection as a slice.
func (iUpdateCollection *IUpdateCollection) GetUpdates() ([]*IUpdate, error) {
	return toIUpdates(iUpdateCollection.disp)
}
