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

package primitive2d_test

import (
	"fmt"

	"go.incompletion.ist/go-scad/primitive2d"
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/value"
)

func ExamplePolygon() {
	polygon := primitive2d.Polygon{
		Points: value.NewFloatsXY(
			// outer points
			[2]float64{10, 0},
			[2]float64{0, 10},
			[2]float64{-10, 0},
			[2]float64{0, -10},
			// inner points
			[2]float64{5, 0},
			[2]float64{0, 5},
			[2]float64{-5, 0},
			[2]float64{0, -5},
		),
		Paths: value.NewIntSets(
			// outer points
			[]int{0, 1, 2, 3},
			// inner points
			[]int{4, 5, 6, 7},
		),
	}

	content, _ := scad.FunctionContent(polygon)
	fmt.Println(content)
	// Output: polygon(paths=[ [0, 1, 2, 3], [4, 5, 6, 7] ], points=[ [10, 0], [0, 10], [-10, 0], [0, -10], [5, 0], [0, 5], [-5, 0], [0, -5] ]);
}
