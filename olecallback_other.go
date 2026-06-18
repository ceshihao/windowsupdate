//go:build !windows

/*
Copyright 2026 Zheng Dayu
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

import "github.com/go-ole/go-ole"

// newNoopDispatch is a no-op stub on non-Windows platforms.
// The COM async methods are only functional on Windows.
func newNoopDispatch() *ole.IDispatch {
	return nil
}
