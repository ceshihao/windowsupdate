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
func TestNewSystemInformation(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	sysInfo, err := NewSystemInformation()
	if err != nil {
		t.Fatalf("NewSystemInformation failed: %v", err)
	}
	if sysInfo == nil {
		t.Fatal("NewSystemInformation returned nil")
	}

	// RebootRequired should be a valid boolean (true or false)
	// This is a read-only property, just verify we can read it
	_ = sysInfo.RebootRequired
}
