package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateSearcher searches for updates on a server.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iupdatesearcher
type IUpdateSearcher struct {
	disp                                *ole.IDispatch
	CanAutomaticallyUpgradeService      bool
	ClientApplicationID                 string
	IncludePotentiallySupersededUpdates bool
	Online                              bool
	ServerSelection                     int32
	ServiceID                           string
}

func toIUpdateSearcher(updateSearcherDisp *ole.IDispatch) (*IUpdateSearcher, error) {
	var err error
	iUpdateSearcher := &IUpdateSearcher{
		disp: updateSearcherDisp,
	}

	if iUpdateSearcher.CanAutomaticallyUpgradeService, err = toBoolErr(oleutil.GetProperty(updateSearcherDisp, "CanAutomaticallyUpgradeService")); err != nil {
		return nil, err
	}

	if iUpdateSearcher.ClientApplicationID, err = toStringErr(oleutil.GetProperty(updateSearcherDisp, "ClientApplicationID")); err != nil {
		return nil, err
	}

	if iUpdateSearcher.IncludePotentiallySupersededUpdates, err = toBoolErr(oleutil.GetProperty(updateSearcherDisp, "IncludePotentiallySupersededUpdates")); err != nil {
		return nil, err
	}

	if iUpdateSearcher.Online, err = toBoolErr(oleutil.GetProperty(updateSearcherDisp, "Online")); err != nil {
		return nil, err
	}

	if iUpdateSearcher.ServerSelection, err = toInt32Err(oleutil.GetProperty(updateSearcherDisp, "ServerSelection")); err != nil {
		return nil, err
	}

	if iUpdateSearcher.ServiceID, err = toStringErr(oleutil.GetProperty(updateSearcherDisp, "ServiceID")); err != nil {
		return nil, err
	}

	return iUpdateSearcher, nil
}

// Search performs a synchronous search for updates. The search uses the search options that are currently configured.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nf-wuapi-iupdatesearcher-search
func (iUpdateSearcher *IUpdateSearcher) Search(criteria string) (*ISearchResult, error) {
	searchResultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSearcher.disp, "Search", criteria))
	if err != nil {
		return nil, err
	}
	return toISearchResult(searchResultDisp)
}
