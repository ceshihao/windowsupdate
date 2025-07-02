//go:build test || !production
// +build test !production

package windowsupdate

import (
	"errors"

	"github.com/go-ole/go-ole"
)

type mockVariant struct{ v interface{} }

func (m *mockVariant) Value() interface{} { return m.v }

// getProperty 是全局mock函数，测试时可覆盖
var getProperty = func(disp *ole.IDispatch, prop string) (*mockVariant, error) {
	return nil, errors.New("not implemented")
}

// withGetProperty 用于临时替换getProperty，便于测试
func withGetProperty(f func(*ole.IDispatch, string) (*mockVariant, error), test func()) {
	old := getProperty
	getProperty = f
	defer func() { getProperty = old }()
	test()
}
