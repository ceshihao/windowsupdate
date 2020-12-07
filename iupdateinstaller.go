package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateInstaller installs or uninstalls updates from or onto a computer.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateinstaller
type IUpdateInstaller struct {
	disp                *ole.IDispatch
	AllowSourcePrompts  bool
	ClientApplicationID string
	IsBusy              bool
	IsForced            bool
	// ParentHwnd                       HWND
	// ParentWindow                     IUnknown
	RebootRequiredBeforeInstallation bool
	Updates                          []*IUpdate
}

func toIUpdateInstaller(updateInstallerDisp *ole.IDispatch) (*IUpdateInstaller, error) {
	var err error
	iUpdateInstaller := &IUpdateInstaller{
		disp: updateInstallerDisp,
	}

	if iUpdateInstaller.AllowSourcePrompts, err = toBoolErr(oleutil.GetProperty(updateInstallerDisp, "AllowSourcePrompts")); err != nil {
		return nil, err
	}

	if iUpdateInstaller.ClientApplicationID, err = toStringErr(oleutil.GetProperty(updateInstallerDisp, "ClientApplicationID")); err != nil {
		return nil, err
	}

	if iUpdateInstaller.IsBusy, err = toBoolErr(oleutil.GetProperty(updateInstallerDisp, "IsBusy")); err != nil {
		return nil, err
	}

	if iUpdateInstaller.IsForced, err = toBoolErr(oleutil.GetProperty(updateInstallerDisp, "IsForced")); err != nil {
		return nil, err
	}

	if iUpdateInstaller.RebootRequiredBeforeInstallation, err = toBoolErr(oleutil.GetProperty(updateInstallerDisp, "RebootRequiredBeforeInstallation")); err != nil {
		return nil, err
	}

	updatesDisp, err := toIDispatchErr(oleutil.GetProperty(updateInstallerDisp, "Updates"))
	if err != nil {
		return nil, err
	}
	if updatesDisp != nil {
		if iUpdateInstaller.Updates, err = toIUpdates(updatesDisp); err != nil {
			return nil, err
		}
	}

	return iUpdateInstaller, nil
}

// Install starts a synchronous installation of the updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller-install
func (iUpdateInstaller *IUpdateInstaller) Install(updates []*IUpdate) (*IInstallationResult, error) {
	updatesDisp, err := toIUpdateCollection(updates)
	if err != nil {
		return nil, err
	}
	if _, err = oleutil.PutProperty(iUpdateInstaller.disp, "Updates", updatesDisp); err != nil {
		return nil, err
	}

	installationResultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller.disp, "Install"))
	if err != nil {
		return nil, err
	}
	return toIInstallationResult(installationResultDisp)
}
