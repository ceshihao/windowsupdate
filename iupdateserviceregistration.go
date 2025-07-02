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

// IUpdateServiceRegistration represents an update service registration.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateserviceregistration
type IUpdateServiceRegistration struct {
	disp                        *ole.IDispatch
	AuthorizationCabPath        string
	IsPendingRegistrationWithAU bool
	IsRegisteredWithAU          bool
	ServiceID                   string
	State                       int32
}

func toIUpdateServiceRegistration(updateServiceRegistrationDisp *ole.IDispatch) (*IUpdateServiceRegistration, error) {
	if updateServiceRegistrationDisp == nil {
		return nil, nil
	}

	var err error
	iUpdateServiceRegistration := &IUpdateServiceRegistration{
		disp: updateServiceRegistrationDisp,
	}

	if iUpdateServiceRegistration.AuthorizationCabPath, err = toStringErr(oleutil.GetProperty(updateServiceRegistrationDisp, "AuthorizationCabPath")); err != nil {
		return nil, err
	}

	if iUpdateServiceRegistration.IsPendingRegistrationWithAU, err = toBoolErr(oleutil.GetProperty(updateServiceRegistrationDisp, "IsPendingRegistrationWithAU")); err != nil {
		return nil, err
	}

	if iUpdateServiceRegistration.IsRegisteredWithAU, err = toBoolErr(oleutil.GetProperty(updateServiceRegistrationDisp, "IsRegisteredWithAU")); err != nil {
		return nil, err
	}

	if iUpdateServiceRegistration.ServiceID, err = toStringErr(oleutil.GetProperty(updateServiceRegistrationDisp, "ServiceID")); err != nil {
		return nil, err
	}

	if iUpdateServiceRegistration.State, err = toInt32Err(oleutil.GetProperty(updateServiceRegistrationDisp, "State")); err != nil {
		return nil, err
	}

	return iUpdateServiceRegistration, nil
}

// GetPropertyValue gets a property value from the service registration.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateserviceregistration-getpropertyvalue
func (iUpdateServiceRegistration *IUpdateServiceRegistration) GetPropertyValue(propertyName string) (interface{}, error) {
	return oleutil.CallMethod(iUpdateServiceRegistration.disp, "GetPropertyValue", propertyName)
}

// SetPropertyValue sets a property value for the service registration.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateserviceregistration-setpropertyvalue
func (iUpdateServiceRegistration *IUpdateServiceRegistration) SetPropertyValue(propertyName string, propertyValue interface{}) error {
	_, err := oleutil.CallMethod(iUpdateServiceRegistration.disp, "SetPropertyValue", propertyName, propertyValue)
	return err
}
