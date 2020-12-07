package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IInstallationResult represents the result of an installation or uninstallation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iinstallationresult
type IInstallationResult struct {
	disp           *ole.IDispatch
	HResult        int32
	RebootRequired bool
	ResultCode     int32 // enum https://docs.microsoft.com/en-us/windows/win32/api/wuapi/ne-wuapi-operationresultcode
}

func toIInstallationResult(installationResultDisp *ole.IDispatch) (*IInstallationResult, error) {
	var err error
	iInstallationResult := &IInstallationResult{
		disp: installationResultDisp,
	}

	if iInstallationResult.HResult, err = toInt32Err(oleutil.GetProperty(installationResultDisp, "HResult")); err != nil {
		return nil, err
	}

	if iInstallationResult.RebootRequired, err = toBoolErr(oleutil.GetProperty(installationResultDisp, "RebootRequired")); err != nil {
		return nil, err
	}

	if iInstallationResult.ResultCode, err = toInt32Err(oleutil.GetProperty(installationResultDisp, "ResultCode")); err != nil {
		return nil, err
	}

	return iInstallationResult, nil
}
