// Copyright 2022 Micah Kemp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package value

import "strconv"

// Int represents a float64 value that can be explicitly set.
type Int struct {
	value int
	set   bool
}

// Set explicitly sets the given value.
func (f *Int) Set(value int) {
	f.value = value
	f.set = true
}

// IsSet returns a boolean indicating if Int has been explicitly set.
func (f Int) IsSet() bool {
	return f.set
}

// Value returns the Int's stored value.
func (f Int) Value() int {
	return f.value
}

// GetParameterValue returns the string representation of Int, and a boolean
// indicating if it was explicitly set.
func (f Int) GetParameterValue() (string, bool) {
	return strconv.FormatInt(int64(f.value), 10), f.set
}

// NewInt returns a new Int with the given value explicitly set.
func NewInt(value int) Int {
	var f Int
	f.Set(value)

	return f
}
