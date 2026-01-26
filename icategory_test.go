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

func TestICategory_StructureFields(t *testing.T) {
	category := &ICategory{
		CategoryID:  "cat-001",
		Description: "Test category description",
		Name:        "Test Category",
		Order:       10,
		Type:        "Software",
	}
	if category.CategoryID != "cat-001" {
		t.Errorf("CategoryID not set correctly, got %s", category.CategoryID)
	}
	if category.Name != "Test Category" {
		t.Errorf("Name not set correctly, got %s", category.Name)
	}
	if category.Order != 10 {
		t.Errorf("Order not set correctly, got %d", category.Order)
	}
	if category.Description != "Test category description" {
		t.Errorf("Description not set correctly, got %s", category.Description)
	}
	if category.Type != "Software" {
		t.Errorf("Type not set correctly, got %s", category.Type)
	}
}
