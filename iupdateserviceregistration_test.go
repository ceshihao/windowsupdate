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

func TestToIUpdateServiceRegistration_NilDispatch(t *testing.T) {
	result, err := toIUpdateServiceRegistration(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIUpdateServiceRegistration_StructureFields(t *testing.T) {
	reg := &IUpdateServiceRegistration{
		IsPendingRegistrationWithAU: true,
		RegistrationState:           1,
	}
	if !reg.IsPendingRegistrationWithAU {
		t.Errorf("IsPendingRegistrationWithAU not set correctly")
	}
	if reg.RegistrationState != 1 {
		t.Errorf("RegistrationState = %d, want 1", reg.RegistrationState)
	}
}
