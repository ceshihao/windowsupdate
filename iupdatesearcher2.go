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

// IUpdateSearcher2 represents an enhanced update searcher interface.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatesearcher2
type IUpdateSearcher2 struct {
	disp *ole.IDispatch
}

func toIUpdateSearcher2(updateSearcher2Disp *ole.IDispatch) (*IUpdateSearcher2, error) {
	if updateSearcher2Disp == nil {
		return nil, nil
	}
	return &IUpdateSearcher2{
		disp: updateSearcher2Disp,
	}, nil
}

// BeginSearch begins an asynchronous search operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-beginsearch
func (iUpdateSearcher2 *IUpdateSearcher2) BeginSearch(criteria string, onProgressChanged interface{}, onCompleted interface{}, state interface{}) (*ISearchResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSearcher2.disp, "BeginSearch", criteria, onProgressChanged, onCompleted, state))
	if err != nil {
		return nil, err
	}
	return toISearchResult(resultDisp)
}

// EndSearch ends an asynchronous search operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-endsearch
func (iUpdateSearcher2 *IUpdateSearcher2) EndSearch(asyncResult *ISearchResult) (*ISearchResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSearcher2.disp, "EndSearch", asyncResult))
	if err != nil {
		return nil, err
	}
	return toISearchResult(resultDisp)
}

// Search searches for updates synchronously.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-search
func (iUpdateSearcher2 *IUpdateSearcher2) Search(criteria string) (*ISearchResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSearcher2.disp, "Search", criteria))
	if err != nil {
		return nil, err
	}
	return toISearchResult(resultDisp)
}

// GetCanAutomaticallyUpgradeService gets whether the service can be automatically upgraded.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-get_canautomaticallyupgradeservice
func (iUpdateSearcher2 *IUpdateSearcher2) GetCanAutomaticallyUpgradeService() (bool, error) {
	return toBoolErr(oleutil.GetProperty(iUpdateSearcher2.disp, "CanAutomaticallyUpgradeService"))
}

// SetCanAutomaticallyUpgradeService sets whether the service can be automatically upgraded.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-put_canautomaticallyupgradeservice
func (iUpdateSearcher2 *IUpdateSearcher2) SetCanAutomaticallyUpgradeService(can bool) error {
	_, err := oleutil.PutProperty(iUpdateSearcher2.disp, "CanAutomaticallyUpgradeService", can)
	return err
}

// GetClientApplicationID gets the client application ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-get_clientapplicationid
func (iUpdateSearcher2 *IUpdateSearcher2) GetClientApplicationID() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateSearcher2.disp, "ClientApplicationID"))
}

// SetClientApplicationID sets the client application ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-put_clientapplicationid
func (iUpdateSearcher2 *IUpdateSearcher2) SetClientApplicationID(id string) error {
	_, err := oleutil.PutProperty(iUpdateSearcher2.disp, "ClientApplicationID", id)
	return err
}

// GetIncludePotentiallySupersededUpdates gets whether potentially superseded updates are included.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-get_includepotentiallysupersededupdates
func (iUpdateSearcher2 *IUpdateSearcher2) GetIncludePotentiallySupersededUpdates() (bool, error) {
	return toBoolErr(oleutil.GetProperty(iUpdateSearcher2.disp, "IncludePotentiallySupersededUpdates"))
}

// SetIncludePotentiallySupersededUpdates sets whether potentially superseded updates are included.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-put_includepotentiallysupersededupdates
func (iUpdateSearcher2 *IUpdateSearcher2) SetIncludePotentiallySupersededUpdates(include bool) error {
	_, err := oleutil.PutProperty(iUpdateSearcher2.disp, "IncludePotentiallySupersededUpdates", include)
	return err
}

// GetServerSelection gets the server selection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-get_serverselection
func (iUpdateSearcher2 *IUpdateSearcher2) GetServerSelection() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateSearcher2.disp, "ServerSelection"))
}

// SetServerSelection sets the server selection.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-put_serverselection
func (iUpdateSearcher2 *IUpdateSearcher2) SetServerSelection(selection int32) error {
	_, err := oleutil.PutProperty(iUpdateSearcher2.disp, "ServerSelection", selection)
	return err
}

// GetTotalHistoryCount gets the total history count.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-get_totalhistorycount
func (iUpdateSearcher2 *IUpdateSearcher2) GetTotalHistoryCount() (int32, error) {
	return toInt32Err(oleutil.GetProperty(iUpdateSearcher2.disp, "TotalHistoryCount"))
}

// QueryHistory queries the update history.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher2-queryhistory
func (iUpdateSearcher2 *IUpdateSearcher2) QueryHistory(startIndex, count int32, criteria string) ([]*IUpdateHistoryEntry, error) {
	historyDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSearcher2.disp, "QueryHistory", startIndex, count, criteria))
	if err != nil {
		return nil, err
	}
	return toIUpdateHistoryEntries(historyDisp)
}
