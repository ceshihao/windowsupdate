package windowsupdate

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// mockVariant 用于 UT，模拟 *ole.VARIANT 并自定义 Value()
type mockVariant struct {
	ole.VARIANT
	val interface{}
}

func (v *mockVariant) Value() interface{} {
	return v.val
}

// fakeVariant 用于 UT 中模拟 *ole.VARIANT 返回值
func fakeVariant(val interface{}) *ole.VARIANT {
	return (*ole.VARIANT)(unsafe.Pointer(&mockVariant{val: val}))
}

// getMockValue 用于 UT，从 *ole.VARIANT 获取 mockVariant 的 val 字段
func getMockValue(v *ole.VARIANT) interface{} {
	return (*mockVariant)(unsafe.Pointer(v)).val
}

// WithOleutilMock 用于 UT 注入 oleutil.GetProperty/CallMethod mock
// getMock/callMock 允许为 nil，test 逻辑在 mock 环境下运行
func WithOleutilMock(getMock func(*ole.IDispatch, string, ...interface{}) (*ole.VARIANT, error), callMock func(*ole.IDispatch, string, ...interface{}) (*ole.VARIANT, error), test func()) {
	oldGet := getProperty
	oldCall := callMethod
	if getMock != nil {
		getProperty = getMock
	}
	if callMock != nil {
		callMethod = callMock
	}
	defer func() {
		getProperty = oldGet
		callMethod = oldCall
	}()
	test()
}
