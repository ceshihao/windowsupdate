package windowsupdate

import (
	"github.com/go-ole/go-ole"
)

// IWebProxy contains the HTTP proxy settings.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iwebproxy
type IWebProxy struct {
	disp               *ole.Dispatch
	Address            string
	AutoDetect         bool
	BypassList         []string
	BypassProxyOnLocal bool
	ReadOnly           bool
	UserName           string
}

func toIWebProxy(webProxyDisp *ole.IDispatch) (*IWebProxy, error) {
	// TODO
	return nil, nil
}
