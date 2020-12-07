package windowsupdate

import (
	"github.com/go-ole/go-ole"
)

// IUpdateDownloadResult contains the properties that indicate the status of a download operation for an update.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatedownloadresult
type IUpdateDownloadResult struct {
	disp       *ole.IDispatch
	HResult    int32
	ResultCode int32
}

func toIUpdateDownloadResult(iUpdateDownloadResultDisp *ole.IDispatch) (*IUpdateDownloadResult, error) {
	// TODO
	return nil, nil
}
