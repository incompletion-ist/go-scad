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

// Bool represents an explicitly settable bool.
type Bool struct {
	value bool
	set   bool
}

// Set explicitly sets the given value.
func (b *Bool) Set(value bool) {
	b.value = value
	b.set = true
}

// IsSet returns a boolean indicating if Bool has been explicitly set.
func (b Bool) IsSet() bool {
	return b.set
}

// Value returns a boolean value representing the stored value, implicit
// or explicit.
func (b Bool) Value() bool {
	return b.value
}

// GetParameterValue returns a string value representing Bool, and a boolean
// indicating if it was explicitly set.
func (b Bool) GetParameterValue() (string, bool) {
	return strconv.FormatBool(b.value), b.set
}

// NewBool returns a new Bool with the given value explicitly set.
func NewBool(value bool) Bool {
	var b Bool
	b.Set(value)

	return b
}
