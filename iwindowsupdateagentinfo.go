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

// IWindowsUpdateAgentInfo retrieves information about the version of Windows Update Agent (WUA).
// https://learn.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iwindowsupdateagentinfo
type IWindowsUpdateAgentInfo struct {
	disp *ole.IDispatch
}

// NewWindowsUpdateAgentInfo creates a new IWindowsUpdateAgentInfo instance.
func NewWindowsUpdateAgentInfo() (*IWindowsUpdateAgentInfo, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.AgentInfo")
	if err != nil {
		return nil, err
	}

	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}

	return &IWindowsUpdateAgentInfo{disp: disp}, nil
}

// WindowsUpdateAgentInfoIndex defines the information to retrieve about WUA.
const (
	WindowsUpdateAgentInfoIndexApiMajorVersion      int32 = 0
	WindowsUpdateAgentInfoIndexApiMinorVersion      int32 = 1
	WindowsUpdateAgentInfoIndexProductVersionString int32 = 2
)

// GetInfo retrieves version information about Windows Update Agent.
// https://learn.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iwindowsupdateagentinfo-getinfo
func (info *IWindowsUpdateAgentInfo) GetInfo(varInfoIdentifier int32) (interface{}, error) {
	result, err := oleutil.CallMethod(info.disp, "QueryInterface", varInfoIdentifier)
	if err != nil {
		// Try GetInfo instead
		result, err = oleutil.CallMethod(info.disp, "QueryInterface", varInfoIdentifier)
		if err != nil {
			return nil, err
		}
	}
	return result.Value(), nil
}

// GetApiMajorVersion returns the major version of the WUA API.
func (info *IWindowsUpdateAgentInfo) GetApiMajorVersion() (int32, error) {
	result, err := oleutil.CallMethod(info.disp, "GetInfo", WindowsUpdateAgentInfoIndexApiMajorVersion)
	if err != nil {
		return 0, err
	}
	return variantToInt32(result), nil
}

// GetApiMinorVersion returns the minor version of the WUA API.
func (info *IWindowsUpdateAgentInfo) GetApiMinorVersion() (int32, error) {
	result, err := oleutil.CallMethod(info.disp, "GetInfo", WindowsUpdateAgentInfoIndexApiMinorVersion)
	if err != nil {
		return 0, err
	}
	return variantToInt32(result), nil
}

// GetProductVersionString returns the product version string of WUA.
func (info *IWindowsUpdateAgentInfo) GetProductVersionString() (string, error) {
	result, err := oleutil.CallMethod(info.disp, "GetInfo", WindowsUpdateAgentInfoIndexProductVersionString)
	if err != nil {
		return "", err
	}
	return variantToString(result), nil
}
