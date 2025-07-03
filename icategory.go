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
	"fmt"

	"github.com/go-ole/go-ole"
)

type _ = ole.IDispatch // Prevent goimports from removing the import

// ICategory represents the category to which an update belongs.
// https://docs.microsoft.com/en-us/windows/win32/api/wuapi/nn-wuapi-icategory
type ICategory struct {
	disp        *ole.IDispatch
	CategoryID  string
	Children    []*ICategory
	Description string
	Image       *IImageInformation
	Name        string
	Order       int32
	Parent      *ICategory
	Type        string
	Updates     []*IUpdate
}

// Package-level injectable function variables for UT mock
var (
	toIDispatchErrFunc      = toIDispatchErr
	toStringErrFunc         = toStringErr
	toInt32ErrFunc          = toInt32Err
	toICategoriesFunc       func(*ole.IDispatch) ([]*ICategory, error)
	toIImageInformationFunc = toIImageInformation
	toICategoryFunc         = toICategory
)

func init() {
	toICategoriesFunc = toICategories
}

func toICategories(categoriesDisp *ole.IDispatch) ([]*ICategory, error) {
	count, err := toInt32ErrFunc(getProperty(categoriesDisp, "Count"))
	if err != nil {
		return nil, err
	}

	categories := make([]*ICategory, 0, count)
	for i := 0; i < int(count); i++ {
		categoryDisp, err := toIDispatchErrFunc(getProperty(categoriesDisp, "Item", i))
		if err != nil {
			return nil, err
		}
		if categoryDisp == nil {
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("categoryDisp is nil at index %d", i)
		}
		category, err := toICategoryFunc(categoryDisp)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func toICategory(categoryDisp *ole.IDispatch) (*ICategory, error) {
	var err error
	iCategory := &ICategory{
		disp: categoryDisp,
	}

	if iCategory.CategoryID, err = toStringErrFunc(getProperty(categoryDisp, "CategoryID")); err != nil {
		return nil, err
	}

	childrenDisp, err := toIDispatchErrFunc(getProperty(categoryDisp, "Children"))
	if err != nil {
		return nil, err
	}
	if childrenDisp != nil {
		if iCategory.Children, err = toICategoriesFunc(childrenDisp); err != nil {
			return nil, err
		}
	}

	if iCategory.Description, err = toStringErrFunc(getProperty(categoryDisp, "Description")); err != nil {
		return nil, err
	}

	imageDisp, err := toIDispatchErrFunc(getProperty(categoryDisp, "Image"))
	if err != nil {
		return nil, err
	}
	if imageDisp != nil {
		if iCategory.Image, err = toIImageInformationFunc(imageDisp); err != nil {
			return nil, err
		}
	}

	if iCategory.Name, err = toStringErrFunc(getProperty(categoryDisp, "Name")); err != nil {
		return nil, err
	}

	if iCategory.Order, err = toInt32ErrFunc(getProperty(categoryDisp, "Order")); err != nil {
		return nil, err
	}

	// parentDisp, err := toIDispatchErrFunc(getProperty(categoryDisp, "Parent"))
	// if err != nil {
	// 	return nil, err
	// }
	// if parentDisp != nil {
	// 	if iCategory.Parent, err = toICategory(parentDisp); err != nil {
	// 		return nil, err
	// 	}
	// }

	if iCategory.Type, err = toStringErrFunc(getProperty(categoryDisp, "Type")); err != nil {
		return nil, err
	}

	// updatesDisp, err := toIDispatchErrFunc(getProperty(categoryDisp, "Updates"))
	// if err != nil {
	// 	return nil, err
	// }
	// if updatesDisp != nil {
	// 	if iCategory.Updates, err = toIUpdates(updatesDisp); err != nil {
	// 		return nil, err
	// 	}
	// }

	return iCategory, nil
}
