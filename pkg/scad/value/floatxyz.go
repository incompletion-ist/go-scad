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

// floatXY enables embedding a FloatXY in FloatXYZ without exporting it.
type floatXY struct {
	FloatXY
}

// FloatXYZ represents a tuple of XYZ float64 values that can be explicitly
// set.
type FloatXYZ struct {
	floatXY
	valueZ float64
	set    bool
}

// Set explicitly sets the given values.
func (xyz *FloatXYZ) Set(x, y, z float64) {
	xyz.valueX = x
	xyz.valueY = y
	xyz.valueZ = z
	xyz.set = true
}

// IsSet returns a boolean indicating if FloatXYZ has been explicitly set.
func (xyz FloatXYZ) IsSet() bool {
	return xyz.set
}

// ValueZ returns the stored value for Z.
func (xyz FloatXYZ) ValueZ() float64 {
	return xyz.valueZ
}

// GetParameterValue returns a string value for the FloatXYZ, and a boolean
// indicating if its value was explicity set.
func (xyz FloatXYZ) GetParameterValue() (string, bool) {
	values := strings.Join([]string{
		strconv.FormatFloat(xyz.valueX, 'f', -1, 64),
		strconv.FormatFloat(xyz.valueY, 'f', -1, 64),
		strconv.FormatFloat(xyz.valueZ, 'f', -1, 64),
	}, ", ")

	value := fmt.Sprintf("[%s]", values)

	return value, xyz.set
}

// NewFloatXYZ creates a new FloatXYZ with its value explicitly set.
func NewFloatXYZ(x, y, z float64) FloatXYZ {
	var xyz FloatXYZ
	xyz.Set(x, y, z)

	return xyz
}
