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

// IUpdateSession2 represents an enhanced update session interface.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iupdatesession2
type IUpdateSession2 struct {
	disp *ole.IDispatch
}

func toIUpdateSession2(updateSession2Disp *ole.IDispatch) (*IUpdateSession2, error) {
	if updateSession2Disp == nil {
		return nil, nil
	}
	return &IUpdateSession2{
		disp: updateSession2Disp,
	}, nil
}

// NewUpdateSession2 creates a new IUpdateSession2 interface.
func NewUpdateSession2() (*IUpdateSession2, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.Session")
	if err != nil {
		return nil, err
	}
	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return toIUpdateSession2(disp)
}

// CreateUpdateDownloader returns an IUpdateDownloader interface for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-createupdatedownloader
func (iUpdateSession2 *IUpdateSession2) CreateUpdateDownloader() (*IUpdateDownloader, error) {
	updateDownloaderDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSession2.disp, "CreateUpdateDownloader"))
	if err != nil {
		return nil, err
	}
	return toIUpdateDownloader(updateDownloaderDisp)
}

// CreateUpdateInstaller returns an IUpdateInstaller2 interface for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-createupdateinstaller
func (iUpdateSession2 *IUpdateSession2) CreateUpdateInstaller() (*IUpdateInstaller2, error) {
	updateInstallerDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSession2.disp, "CreateUpdateInstaller"))
	if err != nil {
		return nil, err
	}
	return toIUpdateInstaller2(updateInstallerDisp)
}

// CreateUpdateSearcher returns an IUpdateSearcher2 interface for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-createupdatesearcher
func (iUpdateSession2 *IUpdateSession2) CreateUpdateSearcher() (*IUpdateSearcher2, error) {
	updateSearcherDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSession2.disp, "CreateUpdateSearcher"))
	if err != nil {
		return nil, err
	}
	return toIUpdateSearcher2(updateSearcherDisp)
}

// CreateUpdateServiceManager returns an IUpdateServiceManager interface for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-createupdateservicemanager
func (iUpdateSession2 *IUpdateSession2) CreateUpdateServiceManager() (*IUpdateServiceManager, error) {
	updateServiceManagerDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSession2.disp, "CreateUpdateServiceManager"))
	if err != nil {
		return nil, err
	}
	return toIUpdateServiceManager(updateServiceManagerDisp)
}

// GetClientApplicationID gets the client application ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-get_clientapplicationid
func (iUpdateSession2 *IUpdateSession2) GetClientApplicationID() (string, error) {
	return toStringErr(oleutil.GetProperty(iUpdateSession2.disp, "ClientApplicationID"))
}

// SetClientApplicationID sets the client application ID.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-put_clientapplicationid
func (iUpdateSession2 *IUpdateSession2) SetClientApplicationID(id string) error {
	_, err := oleutil.PutProperty(iUpdateSession2.disp, "ClientApplicationID", id)
	return err
}

// GetReadOnly gets whether the session is read-only.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-get_readonly
func (iUpdateSession2 *IUpdateSession2) GetReadOnly() (bool, error) {
	return toBoolErr(oleutil.GetProperty(iUpdateSession2.disp, "ReadOnly"))
}

// GetWebProxy gets the web proxy for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-get_webproxy
func (iUpdateSession2 *IUpdateSession2) GetWebProxy() (*IWebProxy, error) {
	webProxyDisp, err := toIDispatchErr(oleutil.GetProperty(iUpdateSession2.disp, "WebProxy"))
	if err != nil {
		return nil, err
	}
	return toIWebProxy(webProxyDisp)
}

// SetWebProxy sets the web proxy for this session.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nf-wuapi-iupdatesession2-put_webproxy
func (iUpdateSession2 *IUpdateSession2) SetWebProxy(webProxy *IWebProxy) error {
	_, err := oleutil.PutProperty(iUpdateSession2.disp, "WebProxy", webProxy.disp)
	return err
}
