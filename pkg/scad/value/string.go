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

// String represents an explicitly settable string.
type String struct {
	value string
	set   bool
}

// Set explicitly sets the given value.
func (s *String) Set(value string) {
	s.value = value
	s.set = true
}

// IsSet returns a boolean indicating if String has been explicitly set.
func (s String) IsSet() bool {
	return s.set
}

// Value returns a boolean value representing the stored value, implicit
// or explicit.
func (s String) Value() string {
	return s.value
}

// GetParameterValue returns a string value representing String, and a boolean
// indicating if it was explicitly set.
func (s String) GetParameterValue() (string, bool) {
	return s.value, s.set
}

// NewString returns a new String with the given value explicitly set.
func NewString(value string) String {
	var s String
	s.Set(value)

	return s
}
