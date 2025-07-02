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

// IUpdateServiceManager represents a manager for update services.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdateservicemanager
type IUpdateServiceManager struct {
	disp *ole.IDispatch
}

func toIUpdateServiceManager(updateServiceManagerDisp *ole.IDispatch) (*IUpdateServiceManager, error) {
	if updateServiceManagerDisp == nil {
		return nil, nil
	}
	return &IUpdateServiceManager{
		disp: updateServiceManagerDisp,
	}, nil
}

// NewUpdateServiceManager creates a new IUpdateServiceManager interface.
func NewUpdateServiceManager() (*IUpdateServiceManager, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.ServiceManager")
	if err != nil {
		return nil, err
	}
	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return toIUpdateServiceManager(disp)
}

// AddScanPackageService adds a scan package service.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-addscanpackageservice
func (iUpdateServiceManager *IUpdateServiceManager) AddScanPackageService(serviceName, scanFileLocation string) (*IUpdateService, error) {
	serviceDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateServiceManager.disp, "AddScanPackageService", serviceName, scanFileLocation))
	if err != nil {
		return nil, err
	}
	return toIUpdateService(serviceDisp)
}

// AddService adds a service to the service manager.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-addservice
func (iUpdateServiceManager *IUpdateServiceManager) AddService(serviceID, authorizationCabPath string) (*IUpdateService, error) {
	serviceDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateServiceManager.disp, "AddService", serviceID, authorizationCabPath))
	if err != nil {
		return nil, err
	}
	return toIUpdateService(serviceDisp)
}

// GetDefaultAUNotificationLevel gets the default automatic update notification level.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-getdefaultaunotificationlevel
func (iUpdateServiceManager *IUpdateServiceManager) GetDefaultAUNotificationLevel() (int32, error) {
	return toInt32Err(oleutil.CallMethod(iUpdateServiceManager.disp, "GetDefaultAUNotificationLevel"))
}

// GetDefaultAUScheduledInstallationDay gets the default automatic update scheduled installation day.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-getdefaultauscheduledinstallationday
func (iUpdateServiceManager *IUpdateServiceManager) GetDefaultAUScheduledInstallationDay() (int32, error) {
	return toInt32Err(oleutil.CallMethod(iUpdateServiceManager.disp, "GetDefaultAUScheduledInstallationDay"))
}

// GetDefaultAUScheduledInstallationTime gets the default automatic update scheduled installation time.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-getdefaultauscheduledinstallationtime
func (iUpdateServiceManager *IUpdateServiceManager) GetDefaultAUScheduledInstallationTime() (int32, error) {
	return toInt32Err(oleutil.CallMethod(iUpdateServiceManager.disp, "GetDefaultAUScheduledInstallationTime"))
}

// QueryServiceRegistration queries for service registration.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-queryserviceregistration
func (iUpdateServiceManager *IUpdateServiceManager) QueryServiceRegistration(serviceID string) (*IUpdateServiceRegistration, error) {
	registrationDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateServiceManager.disp, "QueryServiceRegistration", serviceID))
	if err != nil {
		return nil, err
	}
	return toIUpdateServiceRegistration(registrationDisp)
}

// RegisterServiceWithAU registers a service with automatic updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-registerservicewithau
func (iUpdateServiceManager *IUpdateServiceManager) RegisterServiceWithAU(serviceID string) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "RegisterServiceWithAU", serviceID)
	return err
}

// RemoveService removes a service from the service manager.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-removeservice
func (iUpdateServiceManager *IUpdateServiceManager) RemoveService(serviceID string) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "RemoveService", serviceID)
	return err
}

// SetDefaultAUNotificationLevel sets the default automatic update notification level.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-setdefaultaunotificationlevel
func (iUpdateServiceManager *IUpdateServiceManager) SetDefaultAUNotificationLevel(level int32) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "SetDefaultAUNotificationLevel", level)
	return err
}

// SetDefaultAUScheduledInstallationDay sets the default automatic update scheduled installation day.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-setdefaultauscheduledinstallationday
func (iUpdateServiceManager *IUpdateServiceManager) SetDefaultAUScheduledInstallationDay(day int32) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "SetDefaultAUScheduledInstallationDay", day)
	return err
}

// SetDefaultAUScheduledInstallationTime sets the default automatic update scheduled installation time.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-setdefaultauscheduledinstallationtime
func (iUpdateServiceManager *IUpdateServiceManager) SetDefaultAUScheduledInstallationTime(time int32) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "SetDefaultAUScheduledInstallationTime", time)
	return err
}

// UnregisterServiceWithAU unregisters a service from automatic updates.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdateservicemanager-unregisterservicewithau
func (iUpdateServiceManager *IUpdateServiceManager) UnregisterServiceWithAU(serviceID string) error {
	_, err := oleutil.CallMethod(iUpdateServiceManager.disp, "UnregisterServiceWithAU", serviceID)
	return err
}
