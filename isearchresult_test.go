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

func TestISearchResult_StructureFields(t *testing.T) {
	result := &ISearchResult{
		ResultCode:     OperationResultCodeOrcSucceeded,
		RootCategories: nil,
		Updates:        nil,
		Warnings:       nil,
	}
	if result.ResultCode != OperationResultCodeOrcSucceeded {
		t.Errorf("ResultCode not set correctly")
	}
	if result.RootCategories != nil {
		t.Errorf("RootCategories should be nil")
	}
	if result.Updates != nil {
		t.Errorf("Updates should be nil")
	}
	if result.Warnings != nil {
		t.Errorf("Warnings should be nil")
	}
}
