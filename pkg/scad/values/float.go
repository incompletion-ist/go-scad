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

package values

import "strconv"

// Float represents a float64 value that can be explicitly set.
type Float struct {
	value float64
	set   bool
}

// Set explicitly sets the given value.
func (f *Float) Set(value float64) {
	f.value = value
	f.set = true
}

// IsSet returns a boolean indicating if Float has been explicitly set.
func (f Float) IsSet() bool {
	return f.set
}

// Value returns the Float's stored value.
func (f Float) Value() float64 {
	return f.value
}

// GetParameterValue returns the string representation of Float, and a boolean
// indicating if it was explicitly set.
func (f Float) GetParameterValue() (string, bool) {
	return strconv.FormatFloat(f.value, 'f', -1, 64), f.set
}

// NewFloat returns a new Float with the given value explicitly set.
func NewFloat(value float64) Float {
	var f Float
	f.Set(value)

	return f
}
