package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IDownloadResult represents the result of a download operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-idownloadresult
type IDownloadResult struct {
	disp       *ole.IDispatch
	HResult    int32
	ResultCode int32 // enum https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-operationresultcode
}

func toIDownloadResult(downloadResultDisp *ole.IDispatch) (*IDownloadResult, error) {
	var err error
	iDownloadResult := &IDownloadResult{
		disp: downloadResultDisp,
	}

	if iDownloadResult.HResult, err = toInt32Err(oleutil.GetProperty(downloadResultDisp, "HResult")); err != nil {
		return nil, err
	}

	if iDownloadResult.ResultCode, err = toInt32Err(oleutil.GetProperty(downloadResultDisp, "ResultCode")); err != nil {
		return nil, err
	}

	return iDownloadResult, nil
}

// GetUpdateResult returns an IUpdateDownloadResult interface that contains the download information for a specified update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-idownloadresult-getupdateresult
func (iDownloadResult *IDownloadResult) GetUpdateResult(updateIndex int32) (*IUpdateDownloadResult, error) {
	// TODO
	return nil, nil
}
