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

import (
	"fmt"
	"strconv"
	"strings"
)

// IntSets represents an explicitly settable set of integer sets.
type IntSets struct {
	value [][]int
	set   bool
}

// Set explicitly sets the given value.
func (i *IntSets) Set(value [][]int) {
	i.value = value
	i.set = true
}

// IsSet returns a boolean indicating if IntSets has been explicitly set.
func (i IntSets) IsSet() bool {
	return i.set
}

// Value returns the IntSet's stored value.
func (i IntSets) Value() [][]int {
	return i.value
}

// GetParameterValue returns the string representation of IntSets, and a boolean
// indicating if it was explicitly set.
func (i IntSets) GetParameterValue() (string, bool) {
	intsStrings := make([]string, len(i.value))

	for j, intValues := range i.value {
		intStrings := make([]string, len(intValues))
		for k, intValue := range intValues {
			intStrings[k] = strconv.FormatInt(int64(intValue), 10)
		}

		intsStrings[j] = fmt.Sprintf("[%s]", strings.Join(intStrings, ", "))
	}

	return fmt.Sprintf("[ %s ]", strings.Join(intsStrings, ", ")), i.set
}

// NewIntSets returns a new IntSets with the given value explicitly set.
func NewIntSets(value ...[]int) IntSets {
	var i IntSets
	i.Set(value)

	return i
}
