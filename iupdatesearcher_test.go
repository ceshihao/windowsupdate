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

func TestToIUpdateSearcher(t *testing.T) {
	WithOleutilMock(
		func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(0), nil }, nil,
		func() {
			_, _ = toIUpdateSearcher(&ole.IDispatch{})
		},
	)
}

func TestIUpdateSearcher_Search(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.Search("")
		},
	)
}

func TestIUpdateSearcher_QueryHistory(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.QueryHistory(0, 1)
		},
	)
}

func TestIUpdateSearcher_GetTotalHistoryCount(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.GetTotalHistoryCount()
		},
	)
}

func TestIUpdateSearcher_QueryHistoryAll(t *testing.T) {
	WithOleutilMock(nil, func(_ *ole.IDispatch, _ string, _ ...interface{}) (*ole.VARIANT, error) { return fakeVariant(nil), nil },
		func() {
			s := &IUpdateSearcher{disp: &ole.IDispatch{}}
			_, _ = s.QueryHistoryAll()
		},
	)
}
