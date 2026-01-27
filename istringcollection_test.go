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

func TestToIStringCollection_NilDispatch(t *testing.T) {
	result, err := toIStringCollection(nil)
	if err != nil {
		t.Errorf("expected no error for nil dispatch, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for nil dispatch, got %v", result)
	}
}

func TestIStringCollection_ToSlice_EmptyCollection(t *testing.T) {
	sc := &IStringCollection{
		Count: 0,
	}
	result, err := sc.ToSlice()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestIStringCollectionToStringArrayErr_WithError(t *testing.T) {
	testErr := ole.NewError(0x80070005)
	result, err := iStringCollectionToStringArrayErr(nil, testErr)
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestIStringCollectionToStringArrayErr_WithNilDispatch(t *testing.T) {
	result, err := iStringCollectionToStringArrayErr(nil, nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("expected nil for nil dispatch, got %v", result)
	}
}

func TestIStringCollection_StructureFields(t *testing.T) {
	sc := &IStringCollection{
		Count:    10,
		ReadOnly: true,
	}
	if sc.Count != 10 {
		t.Errorf("Count not set correctly, got %d, want 10", sc.Count)
	}
	if !sc.ReadOnly {
		t.Errorf("ReadOnly not set correctly, got %v, want true", sc.ReadOnly)
	}
}

// COM tests
func TestNewStringCollection(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}
	if collection == nil {
		t.Fatal("NewStringCollection returned nil")
	}
	if collection.disp == nil {
		t.Fatal("collection.disp is nil")
	}
	if collection.Count != 0 {
		t.Errorf("Count = %d, want 0", collection.Count)
	}
}

func TestIStringCollection_AddAndItem(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	// Add items
	testStrings := []string{"item1", "item2", "item3"}
	for _, str := range testStrings {
		_, err := collection.Add(str)
		if err != nil {
			t.Errorf("Add(%q) failed: %v", str, err)
		}
	}

	if collection.Count != int32(len(testStrings)) {
		t.Errorf("Count = %d, want %d", collection.Count, len(testStrings))
	}

	// Get items
	for i, expected := range testStrings {
		item, err := collection.Item(int32(i))
		if err != nil {
			t.Errorf("Item(%d) failed: %v", i, err)
		}
		if item != expected {
			t.Errorf("Item(%d) = %q, want %q", i, item, expected)
		}
	}
}

func TestIStringCollection_ToSlice(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	testStrings := []string{"apple", "banana", "cherry"}
	for _, str := range testStrings {
		_, err := collection.Add(str)
		if err != nil {
			t.Fatalf("Add failed: %v", err)
		}
	}

	slice, err := collection.ToSlice()
	if err != nil {
		t.Fatalf("ToSlice failed: %v", err)
	}
	if len(slice) != len(testStrings) {
		t.Errorf("ToSlice returned %d items, want %d", len(slice), len(testStrings))
	}
	for i, expected := range testStrings {
		if slice[i] != expected {
			t.Errorf("slice[%d] = %q, want %q", i, slice[i], expected)
		}
	}
}

func TestIStringCollection_Clear(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	collection.Add("item1")
	collection.Add("item2")

	err = collection.Clear()
	if err != nil {
		t.Errorf("Clear failed: %v", err)
	}
	if collection.Count != 0 {
		t.Errorf("Count after Clear = %d, want 0", collection.Count)
	}
}

func TestIStringCollection_Insert(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	collection.Add("first")
	collection.Add("third")

	err = collection.Insert(1, "second")
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	if collection.Count != 3 {
		t.Errorf("Count = %d, want 3", collection.Count)
	}

	item, err := collection.Item(1)
	if err != nil {
		t.Errorf("Item(1) failed: %v", err)
	}
	if item != "second" {
		t.Errorf("Item(1) = %q, want \"second\"", item)
	}
}

func TestIStringCollection_RemoveAt(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	collection.Add("item1")
	collection.Add("item2")
	collection.Add("item3")

	err = collection.RemoveAt(1)
	if err != nil {
		t.Errorf("RemoveAt failed: %v", err)
	}
	if collection.Count != 2 {
		t.Errorf("Count = %d, want 2", collection.Count)
	}

	item, err := collection.Item(1)
	if err != nil {
		t.Errorf("Item(1) failed: %v", err)
	}
	if item != "item3" {
		t.Errorf("Item(1) = %q, want \"item3\"", item)
	}
}

func TestIStringCollection_Operations(t *testing.T) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	collection, err := NewStringCollection()
	if err != nil {
		t.Fatalf("NewStringCollection failed: %v", err)
	}

	// Add items
	idx, err := collection.Add("test1")
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if idx != 0 {
		t.Errorf("First Add should return index 0, got %d", idx)
	}

	idx, err = collection.Add("test2")
	if err != nil {
		t.Fatalf("Add second item failed: %v", err)
	}
	if idx != 1 {
		t.Errorf("Second Add should return index 1, got %d", idx)
	}

	// Verify count
	if collection.Count != 2 {
		t.Errorf("Count should be 2, got %d", collection.Count)
	}

	// Get item
	item, err := collection.Item(0)
	if err != nil {
		t.Fatalf("Item(0) failed: %v", err)
	}
	if item != "test1" {
		t.Errorf("Item(0) should be 'test1', got '%s'", item)
	}

	// ToSlice
	slice, err := collection.ToSlice()
	if err != nil {
		t.Fatalf("ToSlice failed: %v", err)
	}
	if len(slice) != 2 {
		t.Errorf("ToSlice should return 2 items, got %d", len(slice))
	}
	if slice[0] != "test1" || slice[1] != "test2" {
		t.Errorf("ToSlice returned wrong values: %v", slice)
	}

	// Insert
	err = collection.Insert(1, "inserted")
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
	if collection.Count != 3 {
		t.Errorf("Count after insert should be 3, got %d", collection.Count)
	}

	// RemoveAt
	err = collection.RemoveAt(1)
	if err != nil {
		t.Fatalf("RemoveAt failed: %v", err)
	}
	if collection.Count != 2 {
		t.Errorf("Count after remove should be 2, got %d", collection.Count)
	}

	// Clear
	err = collection.Clear()
	if err != nil {
		t.Fatalf("Clear failed: %v", err)
	}
	if collection.Count != 0 {
		t.Errorf("Count after clear should be 0, got %d", collection.Count)
	}
}
