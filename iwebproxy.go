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

// IWebProxy contains the HTTP proxy settings.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-iwebproxy
type IWebProxy struct {
	disp               *ole.IDispatch
	Address            string
	AutoDetect         bool
	BypassList         []string
	BypassProxyOnLocal bool
	ReadOnly           bool
	UserName           string
}

func toIWebProxy(webProxyDisp *ole.IDispatch) (*IWebProxy, error) {
	if webProxyDisp == nil {
		return nil, nil
	}
	var err error
	proxy := &IWebProxy{
		disp: webProxyDisp,
	}
	if proxy.Address, err = toStringErr(oleutil.GetProperty(webProxyDisp, "Address")); err != nil {
		return nil, err
	}
	if proxy.AutoDetect, err = toBoolErr(oleutil.GetProperty(webProxyDisp, "AutoDetect")); err != nil {
		return nil, err
	}
	bypassListDisp, err := toIDispatchErr(oleutil.GetProperty(webProxyDisp, "BypassList"))
	if err != nil {
		return nil, err
	}
	if bypassListDisp != nil {
		if proxy.BypassList, err = iStringCollectionToStringArrayErr(bypassListDisp, nil); err != nil {
			return nil, err
		}
	}
	if proxy.BypassProxyOnLocal, err = toBoolErr(oleutil.GetProperty(webProxyDisp, "BypassProxyOnLocal")); err != nil {
		return nil, err
	}
	if proxy.ReadOnly, err = toBoolErr(oleutil.GetProperty(webProxyDisp, "ReadOnly")); err != nil {
		return nil, err
	}
	if proxy.UserName, err = toStringErr(oleutil.GetProperty(webProxyDisp, "UserName")); err != nil {
		return nil, err
	}
	return proxy, nil
}
