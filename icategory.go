package windowsupdate

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// ICategory represents the category to which an update belongs.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-icategory
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

func toICategories(categoriesDisp *ole.IDispatch) ([]*ICategory, error) {
	count, err := toInt32Err(oleutil.GetProperty(categoriesDisp, "Count"))
	if err != nil {
		return nil, err
	}

	categories := make([]*ICategory, 0, count)
	for i := 0; i < int(count); i++ {
		categoryDisp, err := toIDispatchErr(oleutil.GetProperty(categoriesDisp, "Item", i))
		if err != nil {
			return nil, err
		}

		category, err := toICategory(categoryDisp)
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

	if iCategory.CategoryID, err = toStringErr(oleutil.GetProperty(categoryDisp, "CategoryID")); err != nil {
		return nil, err
	}

	childrenDisp, err := toIDispatchErr(oleutil.GetProperty(categoryDisp, "Children"))
	if err != nil {
		return nil, err
	}
	if childrenDisp != nil {
		if iCategory.Children, err = toICategories(childrenDisp); err != nil {
			return nil, err
		}
	}

	if iCategory.Description, err = toStringErr(oleutil.GetProperty(categoryDisp, "Description")); err != nil {
		return nil, err
	}

	imageDisp, err := toIDispatchErr(oleutil.GetProperty(categoryDisp, "Image"))
	if err != nil {
		return nil, err
	}
	if imageDisp != nil {
		if iCategory.Image, err = toIImageInformation(imageDisp); err != nil {
			return nil, err
		}
	}

	if iCategory.Name, err = toStringErr(oleutil.GetProperty(categoryDisp, "Name")); err != nil {
		return nil, err
	}

	if iCategory.Order, err = toInt32Err(oleutil.GetProperty(categoryDisp, "Order")); err != nil {
		return nil, err
	}

	// parentDisp, err := toIDispatchErr(oleutil.GetProperty(categoryDisp, "Parent"))
	// if err != nil {
	// 	return nil, err
	// }
	// if parentDisp != nil {
	// 	if iCategory.Parent, err = toICategory(parentDisp); err != nil {
	// 		return nil, err
	// 	}
	// }

	if iCategory.Type, err = toStringErr(oleutil.GetProperty(categoryDisp, "Type")); err != nil {
		return nil, err
	}

	// updatesDisp, err := toIDispatchErr(oleutil.GetProperty(categoryDisp, "Updates"))
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
