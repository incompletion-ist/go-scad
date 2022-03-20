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

package boolean_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/boolean"
	"go.incompletion.ist/go-scad/primitive3d"
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/transformation"
	"go.incompletion.ist/go-scad/value"
)

func ExampleDifference() {
	var wallThickness float64 = 1
	var outerSize float64 = 20

	outerCube := primitive3d.Cube{Size: value.NewFloat(outerSize)}
	innerCube := scad.Apply(
		primitive3d.Cube{SizeXYZ: value.NewFloatXYZ(
			outerSize-2*wallThickness,
			outerSize-2*wallThickness,
			outerSize,
		)},
		transformation.Translate{
			V: value.NewFloatXYZ(wallThickness, wallThickness, wallThickness),
		},
	)

	shell := scad.Apply(
		outerCube,
		boolean.Difference{
			Children: []interface{}{innerCube},
		},
	)

	content, _ := scad.FunctionContent(shell)
	fmt.Println(content)
	// Output: difference() {
	//   cube(size=20);
	//   translate(v=[1, 1, 1]) {
	//     cube(size=[18, 18, 20]);
	//   }
	// }
}
