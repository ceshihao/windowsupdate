package windowsupdate

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToIUpdateServiceManager(t *testing.T) {
	_, _ = toIUpdateServiceManager(&ole.IDispatch{})
}

func TestNewUpdateServiceManager(t *testing.T) {
	_, _ = NewUpdateServiceManager()
}

func TestIUpdateServiceManager_Methods(t *testing.T) {
	m := &IUpdateServiceManager{disp: &ole.IDispatch{}}
	_, _ = m.AddScanPackageService("", "")
	_, _ = m.AddService("", "")
	_, _ = m.GetDefaultAUNotificationLevel()
	_, _ = m.GetDefaultAUScheduledInstallationDay()
	_, _ = m.GetDefaultAUScheduledInstallationTime()
	_, _ = m.QueryServiceRegistration("")
	_ = m.RegisterServiceWithAU("")
	_ = m.RemoveService("")
	_ = m.SetDefaultAUNotificationLevel(0)
	_ = m.SetDefaultAUScheduledInstallationDay(0)
	_ = m.SetDefaultAUScheduledInstallationTime(0)
	_ = m.UnregisterServiceWithAU("")
}
