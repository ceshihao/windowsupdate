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
	"errors"
	"testing"

	"github.com/go-ole/go-ole"
)

func TestToICategories_HappyPath(t *testing.T) {
	// 构造 categoriesDisp，模拟 Count=2，Item(0/1) 返回不同 categoryDisp
	calls := 0
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, args ...interface{}) (*ole.VARIANT, error) {
			switch prop {
			case "Count":
				return fakeVariant(int32(2)), nil
			case "Item":
				calls++
				return fakeVariant(&ole.IDispatch{}), nil
			case "CategoryID", "Description", "Name", "Type":
				return fakeVariant("mock"), nil
			case "Order":
				return fakeVariant(int32(1)), nil
			case "Image":
				return fakeVariant(&ole.IDispatch{}), nil
			case "Children":
				return fakeVariant(&ole.IDispatch{}), nil
			}
			return nil, errors.New("unexpected prop")
		}, nil,
		func() {
			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return &ole.IDispatch{}, nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			oldToICategories := toICategoriesFunc
			toICategoriesFunc = func(_ *ole.IDispatch) ([]*ICategory, error) {
				return []*ICategory{{Name: "mock"}}, nil
			}
			defer func() { toICategoriesFunc = oldToICategories }()

			cats, err := toICategories(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(cats) != 2 || cats[0].Name != "mock" {
				t.Errorf("unexpected cats: %+v", cats)
			}
		},
	)
}

func TestToICategories_CountErr(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "Count" {
				return nil, errors.New("count error")
			}
			return nil, nil
		}, nil,
		func() {
			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return &ole.IDispatch{}, nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			oldToICategories := toICategoriesFunc
			toICategoriesFunc = func(_ *ole.IDispatch) ([]*ICategory, error) {
				return []*ICategory{{Name: "mock"}}, nil
			}
			defer func() { toICategoriesFunc = oldToICategories }()

			_, err := toICategories(&ole.IDispatch{})
			if err == nil || err.Error() != "count error" {
				t.Errorf("expected count error, got %v", err)
			}
		},
	)
}

func TestToICategory_HappyPath(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			switch prop {
			case "CategoryID":
				return fakeVariant("catid"), nil
			case "Children":
				return fakeVariant(&ole.IDispatch{}), nil
			case "Description":
				return fakeVariant("desc"), nil
			case "Image":
				return fakeVariant(&ole.IDispatch{}), nil
			case "Name":
				return fakeVariant("name"), nil
			case "Order":
				return fakeVariant(int32(1)), nil
			case "Type":
				return fakeVariant("type"), nil
			}
			return nil, errors.New("unexpected prop")
		}, nil,
		func() {
			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return &ole.IDispatch{}, nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToICategories := toICategoriesFunc
			toICategoriesFunc = func(_ *ole.IDispatch) ([]*ICategory, error) {
				return []*ICategory{{Name: "child"}}, nil
			}
			defer func() { toICategoriesFunc = oldToICategories }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			cat, err := toICategory(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cat.Name != "name" || cat.Image.AltText != "img" || len(cat.Children) != 1 || cat.Children[0].Name != "child" {
				t.Errorf("unexpected cat: %+v", cat)
			}
		},
	)
}

func TestToICategory_DescriptionErr(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			if prop == "CategoryID" {
				return fakeVariant("catid"), nil
			}
			if prop == "Description" {
				return nil, errors.New("desc error")
			}
			return fakeVariant(nil), nil
		}, nil,
		func() {
			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return &ole.IDispatch{}, nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			oldToICategories := toICategoriesFunc
			toICategoriesFunc = func(_ *ole.IDispatch) ([]*ICategory, error) {
				return []*ICategory{{Name: "mock"}}, nil
			}
			defer func() { toICategoriesFunc = oldToICategories }()

			_, err := toICategory(&ole.IDispatch{})
			if err == nil || err.Error() != "desc error" {
				t.Errorf("expected desc error, got %v", err)
			}
		},
	)
}

func TestToICategory_ChildrenNil(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			switch prop {
			case "CategoryID", "Description", "Name", "Type":
				return fakeVariant("mock"), nil
			case "Order":
				return fakeVariant(int32(1)), nil
			case "Image":
				return fakeVariant(&ole.IDispatch{}), nil
			case "Children":
				return fakeVariant(nil), nil
			}
			return nil, errors.New("unexpected prop")
		}, nil,
		func() {
			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return nil, nil // childrenDisp 为 nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			cat, err := toICategory(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cat.Children != nil && len(cat.Children) != 0 {
				t.Errorf("expected nil or empty children, got %+v", cat.Children)
			}
		},
	)
}

func TestToICategory_ImageNil(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, prop string, _ ...interface{}) (*ole.VARIANT, error) {
			switch prop {
			case "Count":
				return fakeVariant(int32(0)), nil
			case "CategoryID", "Description", "Name", "Type":
				return fakeVariant("mock"), nil
			case "Order":
				return fakeVariant(int32(1)), nil
			case "Image":
				return fakeVariant(nil), nil
			case "Children":
				return fakeVariant(&ole.IDispatch{}), nil
			}
			panic("unexpected prop: " + prop)
		}, nil,
		func() {
			oldToStringErr := toStringErrFunc
			toStringErrFunc = func(v *ole.VARIANT, err error) (string, error) {
				if err != nil {
					return "", err
				}
				return getMockValue(v).(string), nil
			}
			defer func() { toStringErrFunc = oldToStringErr }()

			oldToInt32Err := toInt32ErrFunc
			toInt32ErrFunc = func(v *ole.VARIANT, err error) (int32, error) {
				if err != nil {
					return 0, err
				}
				return getMockValue(v).(int32), nil
			}
			defer func() { toInt32ErrFunc = oldToInt32Err }()

			oldToIDispatchErr := toIDispatchErrFunc
			toIDispatchErrFunc = func(v *ole.VARIANT, err error) (*ole.IDispatch, error) {
				if err != nil {
					return nil, err
				}
				return &ole.IDispatch{}, nil
			}
			defer func() { toIDispatchErrFunc = oldToIDispatchErr }()

			oldToIImageInformation := toIImageInformationFunc
			toIImageInformationFunc = func(_ *ole.IDispatch) (*IImageInformation, error) {
				return &IImageInformation{AltText: "img"}, nil
			}
			defer func() { toIImageInformationFunc = oldToIImageInformation }()

			cat, err := toICategory(&ole.IDispatch{})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cat.Image != nil && cat.Image.AltText != "img" {
				t.Errorf("unexpected image: %+v", cat.Image)
			}
		},
	)
}
