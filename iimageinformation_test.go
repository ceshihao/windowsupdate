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

func TestIImageInformation_StructureFields(t *testing.T) {
	image := &IImageInformation{
		AltText: "Image alt text",
		Height:  100,
		Width:   200,
		Source:  "https://example.com/image.png",
	}
	if image.AltText != "Image alt text" {
		t.Errorf("AltText not set correctly, got %s", image.AltText)
	}
	if image.Height != 100 {
		t.Errorf("Height not set correctly, got %d", image.Height)
	}
	if image.Width != 200 {
		t.Errorf("Width not set correctly, got %d", image.Width)
	}
	if image.Source != "https://example.com/image.png" {
		t.Errorf("Source not set correctly, got %s", image.Source)
	}
}

func TestToIImageInformation_NilDispatch(t *testing.T) {
	defer func() {
		_ = recover()
	}()

	result, err := toIImageInformation(nil)
	if err == nil && result != nil {
		t.Errorf("expected error or panic for nil dispatch, got result=%v, err=%v", result, err)
	}
}
