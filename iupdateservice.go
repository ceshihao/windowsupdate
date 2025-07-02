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
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateService represents an update service.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateservice
type IUpdateService struct {
	disp                 *ole.IDispatch
	AuthorizationCabPath string
	CanRegisterWithAU    bool
	IsDefaultAUService   bool
	IsManaged            bool
	IsRegisteredWithAU   bool
	IssueDate            *time.Time
	OfflineScanCabPath   string
	RedirectUrls         []string
	ServiceID            string
	ServiceUrl           string
	SetupPrefix          string
}

func toIUpdateService(updateServiceDisp *ole.IDispatch) (*IUpdateService, error) {
	if updateServiceDisp == nil {
		return nil, nil
	}

	var err error
	iUpdateService := &IUpdateService{
		disp: updateServiceDisp,
	}

	if iUpdateService.AuthorizationCabPath, err = toStringErr(oleutil.GetProperty(updateServiceDisp, "AuthorizationCabPath")); err != nil {
		return nil, err
	}

	if iUpdateService.CanRegisterWithAU, err = toBoolErr(oleutil.GetProperty(updateServiceDisp, "CanRegisterWithAU")); err != nil {
		return nil, err
	}

	if iUpdateService.IsDefaultAUService, err = toBoolErr(oleutil.GetProperty(updateServiceDisp, "IsDefaultAUService")); err != nil {
		return nil, err
	}

	if iUpdateService.IsManaged, err = toBoolErr(oleutil.GetProperty(updateServiceDisp, "IsManaged")); err != nil {
		return nil, err
	}

	if iUpdateService.IsRegisteredWithAU, err = toBoolErr(oleutil.GetProperty(updateServiceDisp, "IsRegisteredWithAU")); err != nil {
		return nil, err
	}

	if iUpdateService.IssueDate, err = toTimeErr(oleutil.GetProperty(updateServiceDisp, "IssueDate")); err != nil {
		return nil, err
	}

	if iUpdateService.OfflineScanCabPath, err = toStringErr(oleutil.GetProperty(updateServiceDisp, "OfflineScanCabPath")); err != nil {
		return nil, err
	}

	redirectUrlsDisp, err := toIDispatchErr(oleutil.GetProperty(updateServiceDisp, "RedirectUrls"))
	if err != nil {
		return nil, err
	}
	if redirectUrlsDisp != nil {
		if iUpdateService.RedirectUrls, err = iStringCollectionToStringArrayErr(redirectUrlsDisp, nil); err != nil {
			return nil, err
		}
	}

	if iUpdateService.ServiceID, err = toStringErr(oleutil.GetProperty(updateServiceDisp, "ServiceID")); err != nil {
		return nil, err
	}

	if iUpdateService.ServiceUrl, err = toStringErr(oleutil.GetProperty(updateServiceDisp, "ServiceUrl")); err != nil {
		return nil, err
	}

	if iUpdateService.SetupPrefix, err = toStringErr(oleutil.GetProperty(updateServiceDisp, "SetupPrefix")); err != nil {
		return nil, err
	}

	return iUpdateService, nil
}

// GetPropertyValue gets a property value from the service.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservice-getpropertyvalue
func (iUpdateService *IUpdateService) GetPropertyValue(propertyName string) (interface{}, error) {
	return oleutil.CallMethod(iUpdateService.disp, "GetPropertyValue", propertyName)
}

// SetPropertyValue sets a property value for the service.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservice-setpropertyvalue
func (iUpdateService *IUpdateService) SetPropertyValue(propertyName string, propertyValue interface{}) error {
	_, err := oleutil.CallMethod(iUpdateService.disp, "SetPropertyValue", propertyName, propertyValue)
	return err
}
