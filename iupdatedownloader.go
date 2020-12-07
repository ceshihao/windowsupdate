package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateDownloader downloads updates from the server.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatedownloaders
type IUpdateDownloader struct {
	disp                *ole.IDispatch
	ClientApplicationID string
	IsForced            bool
	Priority            int32 // enum https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-downloadpriority
	Updates             []*IUpdate
}

func toIUpdateDownloader(updateDownloaderDisp *ole.IDispatch) (*IUpdateDownloader, error) {
	var err error
	iUpdateDownloader := &IUpdateDownloader{
		disp: updateDownloaderDisp,
	}

	if iUpdateDownloader.ClientApplicationID, err = toStringErr(oleutil.GetProperty(updateDownloaderDisp, "ClientApplicationID")); err != nil {
		return nil, err
	}

	if iUpdateDownloader.IsForced, err = toBoolErr(oleutil.GetProperty(updateDownloaderDisp, "IsForced")); err != nil {
		return nil, err
	}

	if iUpdateDownloader.Priority, err = toInt32Err(oleutil.GetProperty(updateDownloaderDisp, "Priority")); err != nil {
		return nil, err
	}

	updatesDisp, err := toIDispatchErr(oleutil.GetProperty(updateDownloaderDisp, "Updates"))
	if err != nil {
		return nil, err
	}
	if updatesDisp != nil {
		if iUpdateDownloader.Updates, err = toIUpdates(updatesDisp); err != nil {
			return nil, err
		}
	}

	return iUpdateDownloader, nil
}

// Download starts a synchronous download of the content files that are associated with the updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatedownloader-download
func (iUpdateDownloader *IUpdateDownloader) Download(updates []*IUpdate) (*IDownloadResult, error) {
	updatesDisp, err := toIUpdateCollection(updates)
	if err != nil {
		return nil, err
	}
	if _, err = oleutil.PutProperty(iUpdateDownloader.disp, "Updates", updatesDisp); err != nil {
		return nil, err
	}

	downloadResultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateDownloader.disp, "Download"))
	if err != nil {
		return nil, err
	}
	return toIDownloadResult(downloadResultDisp)
}
