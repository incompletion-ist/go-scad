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

package value_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/value"
)

func ExampleNewFloatsXYRelative() {
	// starting at [1, 1], move in one direction by 2 units to form a square
	floatsXY := value.NewFloatsXYRelative(
		[2]float64{1, 1},
		[2]float64{-2, 0},
		[2]float64{0, -2},
		[2]float64{2, 0},
	)

	paramValue, _ := floatsXY.GetParameterValue()
	fmt.Printf("paramValue: %s", paramValue)
	// Output: paramValue: [ [1, 1], [-1, 1], [-1, -1], [1, -1] ]
}
