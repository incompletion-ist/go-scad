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

// floatTupleStrings returns a string for a tuple of floats.
func floatTupleString(floats ...float64) string {
	floatStrings := make([]string, len(floats))

	for i, float := range floats {
		floatStrings[i] = strconv.FormatFloat(float, 'f', -1, 64)
	}

	return fmt.Sprintf("[%s]", strings.Join(floatStrings, ", "))
}

// FloatXY represents a tuple of XY float64 values that can be explicitly
// set.
type FloatXY struct {
	valueX float64
	valueY float64
	set    bool
}

// Set explicitly sets the given values.
func (xy *FloatXY) Set(x, y float64) {
	xy.valueX = x
	xy.valueY = y
	xy.set = true
}

// IsSet returns a boolean indicating if FloatXY has been explicitly set.
func (xy FloatXY) IsSet() bool {
	return xy.set
}

// ValueX returns the stored value for X.
func (xy FloatXY) ValueX() float64 {
	return xy.valueX
}

// ValueY returns the stored value for Y.
func (xy FloatXY) ValueY() float64 {
	return xy.valueY
}

// GetParameterValue returns a string value for the FloatXY, and a boolean
// indicating if its value was explicity set.
func (xy FloatXY) GetParameterValue() (string, bool) {
	value := floatTupleString(xy.valueX, xy.valueY)

	return value, xy.set
}

// NewFloatXY creates a new FloatXY with its value explicitly set.
func NewFloatXY(x, y float64) FloatXY {
	var xy FloatXY
	xy.Set(x, y)

	return xy
}
