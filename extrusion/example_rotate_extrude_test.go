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

package extrusion_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/extrusion"
	"go.incompletion.ist/go-scad/primitive2d"
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/transformation"
	"go.incompletion.ist/go-scad/value"
)

func ExampleRotateExtrude() {
	halfDonut := scad.Apply(
		primitive2d.Circle{
			D: value.NewFloat(5),
		},
		transformation.Translate{
			V: value.NewFloatXYZ(5, 0, 0),
		},
		extrusion.RotateExtrude{
			Angle: value.NewFloat(180),
		},
	)

	content, _ := scad.FunctionContent(halfDonut)
	fmt.Println(content)
	// Output: rotate_extrude(angle=180) {
	//   translate(v=[5, 0, 0]) {
	//     circle(d=5);
	//   }
	// }
}
