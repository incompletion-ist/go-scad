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

package primitive3d_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/primitive3d"
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/value"
)

func ExampleCylinder() {
	cylinder := primitive3d.Cylinder{
		H:  value.NewFloat(50),
		D1: value.NewFloat(20),
		D2: value.NewFloat(5),
	}

	content, _ := scad.FunctionContent(cylinder)
	fmt.Println(content)
	// Output: cylinder(d1=20, d2=5, h=50);
}
