//go:build windows
// +build windows

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

import "testing"

func TestToIWebProxy_NilDispatch(t *testing.T) {
	// Test with nil dispatch - should handle gracefully
	proxy := &IWebProxy{
		disp: nil,
	}
	if proxy.disp != nil {
		t.Errorf("expected nil dispatch")
	}
}

func TestIWebProxy_StructureFields(t *testing.T) {
	proxy := &IWebProxy{
		Address:            "http://proxy:8080",
		AutoDetect:         true,
		BypassList:         []string{"localhost", "127.0.0.1"},
		BypassProxyOnLocal: true,
		ReadOnly:           false,
		UserName:           "user",
	}
	if proxy.Address != "http://proxy:8080" {
		t.Errorf("Address not set correctly, got %s", proxy.Address)
	}
	if !proxy.AutoDetect {
		t.Errorf("AutoDetect not set correctly")
	}
	if !proxy.BypassProxyOnLocal {
		t.Errorf("BypassProxyOnLocal not set correctly")
	}
	if len(proxy.BypassList) != 2 {
		t.Errorf("BypassList length incorrect, got %d", len(proxy.BypassList))
	}
}
