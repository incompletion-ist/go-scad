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

	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/extrusion"
	"github.com/micahkemp/scad/pkg/scad/flat"
	"github.com/micahkemp/scad/pkg/scad/transforms"
	"github.com/micahkemp/scad/pkg/scad/values"
)

func ExampleRotateExtrude() {
	halfDonut := scad.Apply(
		flat.Circle{
			D: values.NewFloat(5),
		},
		transforms.Translate{
			V: values.NewFloatXYZ(5, 0, 0),
		},
		extrusion.RotateExtrude{
			Angle: values.NewFloat(180),
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