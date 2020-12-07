package windowsupdate

import (
	"github.com/go-ole/go-ole"
)

// IUpdateIdentity represents the unique identifier of an update.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iupdateidentity
type IUpdateIdentity struct {
	disp           *ole.IDispatch
	RevisionNumber int64
	UpdateID       string
}

func toIUpdateIdentity(updateIdentityDisp *ole.IDispatch) (*IUpdateIdentity, error) {
	// TODO
	return nil, nil
}
