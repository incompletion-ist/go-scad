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

package transformation_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/primitive3d"
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/transformation"
	"go.incompletion.ist/go-scad/value"
)

func Example() {
	thing := scad.Apply(
		primitive3d.Cube{
			Size: value.NewFloat(5),
		},
		transformation.Color{
			C: value.NewFloatXYZ(0.1, 0.2, 0.3),
		},
		transformation.Translate{
			V: value.NewFloatXYZ(10, 15, 20),
		},
	)

	content, _ := scad.FunctionContent(thing)
	fmt.Println(content)
	// Output: translate(v=[10, 15, 20]) {
	//   color(c=[0.1, 0.2, 0.3]) {
	//     cube(size=5);
	//   }
	// }
}
