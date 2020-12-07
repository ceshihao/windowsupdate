package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateIdentity represents the unique identifier of an update.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iupdateidentity
type IUpdateIdentity struct {
	disp           *ole.IDispatch
	RevisionNumber int32
	UpdateID       string
}

func toIUpdateIdentity(updateIdentityDisp *ole.IDispatch) (*IUpdateIdentity, error) {
	var err error
	iUpdateIdentity := &IUpdateIdentity{
		disp: updateIdentityDisp,
	}

	if iUpdateIdentity.RevisionNumber, err = toInt32Err(oleutil.GetProperty(updateIdentityDisp, "RevisionNumber")); err != nil {
		return nil, err
	}

	if iUpdateIdentity.UpdateID, err = toStringErr(oleutil.GetProperty(updateIdentityDisp, "UpdateID")); err != nil {
		return nil, err
	}

	return iUpdateIdentity, nil
}
