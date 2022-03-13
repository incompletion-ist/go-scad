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
	"strings"
)

// floatTuplesString returns a string for tuples of floats.
func floatTuplesString(floatsTuples ...[]float64) string {
	floatTupleStrings := make([]string, len(floatsTuples))

	for i, floatsTuple := range floatsTuples {
		floatTupleStrings[i] = floatTupleString(floatsTuple...)
	}

	return fmt.Sprintf("[ %s ]", strings.Join(floatTupleStrings, ", "))
}

// FloatsXY represents a list of XY floats that can be explicitly set.
type FloatsXY struct {
	value [][2]float64
	set   bool
}

// Set sets the given values.
func (xy *FloatsXY) Set(value ...[2]float64) {
	xy.value = value
	xy.set = true
}

// IsSet returns a boolean indicating if FloatsXY has been explicitly set.
func (xy FloatsXY) IsSet() bool {
	return xy.set
}

// Value returns the value stored in the FloatsXY.
func (xy FloatsXY) Value() [][2]float64 {
	return xy.value
}

// GetParameterValue returns a string value for the FloatsXY, and a boolean
// indicating if its value was explicity set.
func (xy FloatsXY) GetParameterValue() (string, bool) {
	valuesSlice := make([][]float64, len(xy.value))
	for i, values := range xy.value {
		valuesSlice[i] = make([]float64, len(values))
		for j, value := range values {
			valuesSlice[i][j] = value
		}
	}

	valueString := floatTuplesString(valuesSlice...)

	return valueString, xy.set
}

// NewFloatsXY creates a new FloatsXY with its value explicitly set.
func NewFloatsXY(value ...[2]float64) FloatsXY {
	var xy FloatsXY
	xy.Set(value...)

	return xy
}
