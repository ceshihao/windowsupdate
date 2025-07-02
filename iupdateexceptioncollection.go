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

// IUpdateExceptionCollection represents a collection of update exceptions.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateexceptioncollection
type IUpdateExceptionCollection struct {
	disp *ole.IDispatch
}

func toIUpdateExceptionCollection(updateExceptionCollectionDisp *ole.IDispatch) (*IUpdateExceptionCollection, error) {
	if updateExceptionCollectionDisp == nil {
		return nil, nil
	}
	return &IUpdateExceptionCollection{
		disp: updateExceptionCollectionDisp,
	}, nil
}

// Count returns the number of exceptions in the collection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateexceptioncollection-get_count
func (iUpdateExceptionCollection *IUpdateExceptionCollection) Count() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateExceptionCollection.disp, "Count"))
}

// Item returns the exception at the specified index.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateexceptioncollection-item
func (iUpdateExceptionCollection *IUpdateExceptionCollection) Item(index int32) (*IUpdateException, error) {
	exceptionDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateExceptionCollection.disp, "Item", index))
	if err != nil {
		return nil, err
	}
	return toIUpdateException(exceptionDisp)
}

// GetExceptions returns all exceptions in the collection as a slice.
func (iUpdateExceptionCollection *IUpdateExceptionCollection) GetExceptions() ([]*IUpdateException, error) {
	count, err := iUpdateExceptionCollection.Count()
	if err != nil {
		return nil, err
	}

	exceptions := make([]*IUpdateException, 0, count)
	for i := 0; i < int(count); i++ {
		exception, err := iUpdateExceptionCollection.Item(int32(i))
		if err != nil {
			return nil, err
		}
		exceptions = append(exceptions, exception)
	}
	return exceptions, nil
} 