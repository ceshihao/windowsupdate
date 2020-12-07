package windowsupdate

import (
	"github.com/go-ole/go-ole"
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
	Order       int64
	Parent      *ICategory
	Type        string
	Updates     []*IUpdate
}

func toICategories(categoriesDisp *ole.IDispatch) ([]*ICategory, error) {
	// TODO
	return nil, nil
}
