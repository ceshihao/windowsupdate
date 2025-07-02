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

// IUpdateInstaller2 represents an enhanced update installer interface.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateinstaller2
type IUpdateInstaller2 struct {
	disp *ole.IDispatch
}

func toIUpdateInstaller2(updateInstaller2Disp *ole.IDispatch) (*IUpdateInstaller2, error) {
	if updateInstaller2Disp == nil {
		return nil, nil
	}
	return &IUpdateInstaller2{
		disp: updateInstaller2Disp,
	}, nil
}

// BeginInstall begins an asynchronous installation operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-begininstall
func (iUpdateInstaller2 *IUpdateInstaller2) BeginInstall(onProgressChanged interface{}, onCompleted interface{}, state interface{}) (*IUpdateInstallResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "BeginInstall", onProgressChanged, onCompleted, state))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// BeginUninstall begins an asynchronous uninstallation operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-beginuninstall
func (iUpdateInstaller2 *IUpdateInstaller2) BeginUninstall(onProgressChanged interface{}, onCompleted interface{}, state interface{}) (*IUpdateInstallResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "BeginUninstall", onProgressChanged, onCompleted, state))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// EndInstall ends an asynchronous installation operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-endinstall
func (iUpdateInstaller2 *IUpdateInstaller2) EndInstall(asyncResult *IUpdateInstallResult) (*IUpdateInstallResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "EndInstall", asyncResult))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// EndUninstall ends an asynchronous uninstallation operation.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-enduninstall
func (iUpdateInstaller2 *IUpdateInstaller2) EndUninstall(asyncResult *IUpdateInstallResult) (*IUpdateInstallResult, error) {
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "EndUninstall", asyncResult))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// Install installs updates synchronously.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-install
func (iUpdateInstaller2 *IUpdateInstaller2) Install(updates []*IUpdate) (*IUpdateInstallResult, error) {
	updatesDisp, err := toIUpdateCollection(updates)
	if err != nil {
		return nil, err
	}
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "Install", updatesDisp))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// Uninstall uninstalls updates synchronously.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-uninstall
func (iUpdateInstaller2 *IUpdateInstaller2) Uninstall(updates []*IUpdate) (*IUpdateInstallResult, error) {
	updatesDisp, err := toIUpdateCollection(updates)
	if err != nil {
		return nil, err
	}
	resultDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateInstaller2.disp, "Uninstall", updatesDisp))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstallResult(resultDisp)
}

// GetAllowSourcePrompts gets whether source prompts are allowed.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-get_allowsourceprompts
func (iUpdateInstaller2 *IUpdateInstaller2) GetAllowSourcePrompts() (bool, error) {
	return toBoolErr(oleutil.GetProperty(iUpdateInstaller2.disp, "AllowSourcePrompts"))
}

// SetAllowSourcePrompts sets whether source prompts are allowed.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-put_allowsourceprompts
func (iUpdateInstaller2 *IUpdateInstaller2) SetAllowSourcePrompts(allow bool) error {
	_, err := oleutil.PutProperty(iUpdateInstaller2.disp, "AllowSourcePrompts", allow)
	return err
}

// GetForceQuiet gets whether force quiet mode is enabled.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-get_forcequiet
func (iUpdateInstaller2 *IUpdateInstaller2) GetForceQuiet() (bool, error) {
	return toBoolErr(oleutil.GetProperty(iUpdateInstaller2.disp, "ForceQuiet"))
}

// SetForceQuiet sets whether force quiet mode is enabled.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-put_forcequiet
func (iUpdateInstaller2 *IUpdateInstaller2) SetForceQuiet(force bool) error {
	_, err := oleutil.PutProperty(iUpdateInstaller2.disp, "ForceQuiet", force)
	return err
}

// GetUpdates gets the updates to install.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-get_updates
func (iUpdateInstaller2 *IUpdateInstaller2) GetUpdates() ([]*IUpdate, error) {
	updatesDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateInstaller2.disp, "Updates"))
	if err != nil {
		return nil, err
	}
	return toIUpdates(updatesDisp)
}

// SetUpdates sets the updates to install.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateinstaller2-put_updates
func (iUpdateInstaller2 *IUpdateInstaller2) SetUpdates(updates []*IUpdate) error {
	updatesDisp, err := toIUpdateCollection(updates)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(iUpdateInstaller2.disp, "Updates", updatesDisp)
	return err
} 