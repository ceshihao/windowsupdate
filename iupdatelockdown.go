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

// IUpdateLockdown represents an update lockdown interface.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatelockdown
type IUpdateLockdown struct {
	disp *ole.IDispatch
}

func toIUpdateLockdown(updateLockdownDisp *ole.IDispatch) (*IUpdateLockdown, error) {
	if updateLockdownDisp == nil {
		return nil, nil
	}
	return &IUpdateLockdown{
		disp: updateLockdownDisp,
	}, nil
}

// NewUpdateLockdown creates a new IUpdateLockdown interface.
func NewUpdateLockdown() (*IUpdateLockdown, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.Lockdown")
	if err != nil {
		return nil, err
	}
	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return toIUpdateLockdown(disp)
}

// GetLockdownPolicy gets the lockdown policy.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatelockdown-getlockdownpolicy
func (iUpdateLockdown *IUpdateLockdown) GetLockdownPolicy() (int32, error) {
	return toInt32Err(oleutil.CallMethod(iUpdateLockdown.disp, "GetLockdownPolicy"))
}

// SetLockdownPolicy sets the lockdown policy.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatelockdown-setlockdownpolicy
func (iUpdateLockdown *IUpdateLockdown) SetLockdownPolicy(policy int32) error {
	_, err := oleutil.CallMethod(iUpdateLockdown.disp, "SetLockdownPolicy", policy)
	return err
} 