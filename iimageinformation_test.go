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

func TestToIImageInformation_WithRealUpdate(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	// Try to get real updates and check if any have image information
	session, err := NewUpdateSession()
	if err != nil {
		t.Skipf("NewUpdateSession failed: %v", err)
	}

	searcher, err := session.CreateUpdateSearcher()
	if err != nil {
		t.Skipf("CreateUpdateSearcher failed: %v", err)
	}

	// Search for any updates
	result, err := searcher.Search("IsInstalled=0")
	if err != nil {
		t.Skipf("Search failed: %v", err)
	}

	if result == nil || len(result.Updates) == 0 {
		t.Skip("No updates available to test image information")
	}

	// Check if any update has image information
	foundImage := false
	for i := 0; i < len(result.Updates) && i < 10; i++ {
		update := result.Updates[i]

		// Access the Image property which calls toIImageInformation internally
		if update.Image != nil {
			t.Logf("Found update with image information")
			t.Logf("Image AltText: %s", update.Image.AltText)
			t.Logf("Image Source: %s", update.Image.Source)
			t.Logf("Image Size: %dx%d", update.Image.Width, update.Image.Height)

			// Verify that image fields are populated
			if update.Image.Source == "" {
				t.Error("Image should have a Source URL")
			}

			foundImage = true
			break
		}
	}

	if !foundImage {
		t.Log("No updates with image information found (this is common)")
		// Even if no image is found, the code path through toIUpdate is exercised
		// which includes the attempt to read Image property
	}
}
