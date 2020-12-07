package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdateSession represents a session in which the caller can perform operations that involve updates.
// For example, this interface represents sessions in which the caller performs a search, download, installation, or uninstallation operation.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iupdatesession
type IUpdateSession struct {
	disp                *ole.IDispatch
	ClientApplicationID string
	ReadOnly            bool
	WebProxy            *IWebProxy
}

func toIUpdateSession(updateSessionDisp *ole.IDispatch) (*IUpdateSession, error) {
	var err error
	iUpdateSession := &IUpdateSession{
		disp: updateSessionDisp,
	}

	if iUpdateSession.ClientApplicationID, err = toStringErr(oleutil.GetProperty(updateSessionDisp, "ClientApplicationID")); err != nil {
		return nil, err
	}

	if iUpdateSession.ReadOnly, err = toBoolErr(oleutil.GetProperty(updateSessionDisp, "ReadOnly")); err != nil {
		return nil, err
	}

	webProxyDisp, err := toIDispatchErr(oleutil.GetProperty(updateSessionDisp, "WebProxy"))
	if err != nil {
		return nil, err
	}

	if iUpdateSession.WebProxy, err = toIWebProxy(webProxyDisp); err != nil {
		return nil, err
	}

	return iUpdateSession, nil
}

// NewUpdateSession creates a new IUpdateSession interface.
func NewUpdateSession() (*IUpdateSession, error) {
	unknown, err := oleutil.CreateObject("Microsoft.Update.Session")
	if err != nil {
		return nil, err
	}
	disp, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return toIUpdateSession(disp)
}

// CreateUpdateSearcher returns an IUpdateSearcher interface for this session.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nf-wuapi-iupdatesession-createupdatesearcher
func (iUpdateSession *IUpdateSession) CreateUpdateSearcher() (*IUpdateSearcher, error) {
	updateSearcherDisp, err := toIDispatchErr(oleutil.CallMethod(iUpdateSession.disp, "CreateUpdateSearcher"))
	if err != nil {
		return nil, err
	}

	return toIUpdateSearcher(updateSearcherDisp)
}
