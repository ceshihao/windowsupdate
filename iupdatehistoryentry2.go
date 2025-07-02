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
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateHistoryEntry2 represents an enhanced update history entry interface.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatehistoryentry2
type IUpdateHistoryEntry2 struct {
	disp *ole.IDispatch
}

func toIUpdateHistoryEntry2(updateHistoryEntry2Disp *ole.IDispatch) (*IUpdateHistoryEntry2, error) {
	if updateHistoryEntry2Disp == nil {
		return nil, nil
	}
	return &IUpdateHistoryEntry2{
		disp: updateHistoryEntry2Disp,
	}, nil
}

// GetCategories gets the categories of the update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_categories
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetCategories() ([]*ICategory, error) {
	categoriesDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Categories"))
	if err != nil {
		return nil, err
	}
	return toICategories(categoriesDisp)
}

// GetClientApplicationID gets the client application ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_clientapplicationid
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetClientApplicationID() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "ClientApplicationID"))
}

// GetDate gets the date of the history entry.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_date
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetDate() (*time.Time, error) {
	return toTimeErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Date"))
}

// GetDescription gets the description of the update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_description
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetDescription() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Description"))
}

// GetHResult gets the HRESULT of the operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_hresult
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetHResult() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "HResult"))
}

// GetIdentity gets the identity of the update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_identity
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetIdentity() (*IUpdateIdentity, error) {
	identityDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Identity"))
	if err != nil {
		return nil, err
	}
	return toIUpdateIdentity(identityDisp)
}

// GetOperation gets the operation performed.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_operation
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetOperation() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Operation"))
}

// GetResultCode gets the result code of the operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_resultcode
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetResultCode() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "ResultCode"))
}

// GetServiceID gets the service ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_serviceid
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetServiceID() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "ServiceID"))
}

// GetSupportUrl gets the support URL.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_supporturl
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetSupportUrl() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "SupportUrl"))
}

// GetTitle gets the title of the update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_title
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetTitle() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "Title"))
}

// GetUnmappedResultCode gets the unmapped result code.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatehistoryentry2-get_unmappedresultcode
func (iUpdateHistoryEntry2 *IUpdateHistoryEntry2) GetUnmappedResultCode() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateHistoryEntry2.disp, "UnmappedResultCode"))
}
