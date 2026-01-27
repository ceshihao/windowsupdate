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

import (
	"testing"

	"github.com/go-ole/go-ole"
)

// COM tests
func TestNewWindowsUpdateAgentInfo(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}
	if agentInfo == nil {
		t.Fatal("NewWindowsUpdateAgentInfo returned nil")
	}
}

func TestIWindowsUpdateAgentInfo_GetApiMajorVersion(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	apiMajor, err := agentInfo.GetApiMajorVersion()
	if err != nil {
		t.Fatalf("GetApiMajorVersion failed: %v", err)
	}
	// API major version should be >= 1
	if apiMajor < 1 {
		t.Errorf("ApiMajorVersion seems invalid: %d", apiMajor)
	}
}

func TestIWindowsUpdateAgentInfo_GetApiMinorVersion(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	apiMinor, err := agentInfo.GetApiMinorVersion()
	if err != nil {
		t.Fatalf("GetApiMinorVersion failed: %v", err)
	}
	// API minor version should be >= 0
	if apiMinor < 0 {
		t.Errorf("ApiMinorVersion seems invalid: %d", apiMinor)
	}
}

func TestIWindowsUpdateAgentInfo_GetProductVersionString(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	agentInfo, err := NewWindowsUpdateAgentInfo()
	if err != nil {
		t.Fatalf("NewWindowsUpdateAgentInfo failed: %v", err)
	}

	productVersion, err := agentInfo.GetProductVersionString()
	if err != nil {
		t.Fatalf("GetProductVersionString failed: %v", err)
	}
	if productVersion == "" {
		t.Error("ProductVersionString is empty")
	}
}
