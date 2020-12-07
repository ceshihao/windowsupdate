package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// ISearchResult represents the result of a search.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-isearchresult
type ISearchResult struct {
	disp           *ole.IDispatch
	ResultCode     int32 // enum https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/ne-wuapi-operationresultcode
	RootCategories []*ICategory
	Updates        []*IUpdate
	Warnings       []*IUpdateException
}

func toISearchResult(searchResultDisp *ole.IDispatch) (*ISearchResult, error) {
	var err error
	iSearchResult := &ISearchResult{
		disp: searchResultDisp,
	}

	if iSearchResult.ResultCode, err = toInt32Err(oleutil.GetProperty(searchResultDisp, "ResultCode")); err != nil {
		return nil, err
	}

	rootCategoriesDisp, err := toIDispatchErr(oleutil.GetProperty(searchResultDisp, "RootCategories"))
	if err != nil {
		return nil, err
	}
	if iSearchResult.RootCategories, err = toICategories(rootCategoriesDisp); err != nil {
		return nil, err
	}

	updatesDisp, err := toIDispatchErr(oleutil.GetProperty(searchResultDisp, "Updates"))
	if err != nil {
		return nil, err
	}
	if iSearchResult.Updates, err = toIUpdates(updatesDisp); err != nil {
		return nil, err
	}

	warningsDisp, err := toIDispatchErr(oleutil.GetProperty(searchResultDisp, "Warnings"))
	if err != nil {
		return nil, err
	}
	if iSearchResult.Warnings, err = toIUpdateExceptions(warningsDisp); err != nil {
		return nil, err
	}

	return iSearchResult, nil
}
