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

package die

import (
	"fmt"

	"github.com/micahkemp/scad/examples/dice/dimples"
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/booleans"
	"github.com/micahkemp/scad/pkg/scad/transformation"
)

// dimplesPlacement defines the rotation and translation for each Dimples arrangement.
var dimplesPlacement = []struct {
	count     int
	rotateA   int
	rotateV   [3]int
	translate [3]int
}{
	{
		count:     1,
		translate: [3]int{0, 0, -1},
	},
	{
		count:     2,
		rotateA:   90,
		rotateV:   [3]int{0, 1, 0},
		translate: [3]int{-1, 0, 0},
	},
	{
		count:     3,
		rotateA:   90,
		rotateV:   [3]int{1, 0, 0},
		translate: [3]int{0, 1, 0},
	},
	{
		count:     4,
		rotateA:   90,
		rotateV:   [3]int{-1, 0, 0},
		translate: [3]int{0, -1, 0},
	},
	{
		count:     5,
		rotateA:   90,
		rotateV:   [3]int{0, -1, 0},
		translate: [3]int{1, 0, 0},
	},
	{
		count:     6,
		rotateA:   180,
		rotateV:   [3]int{0, 1, 0},
		translate: [3]int{0, 0, 1},
	},
}

// Dimples is all of the Dimples for a Die.
type Dimples struct {
	Name scad.ModuleName `scad:"die_dimples"`

	// Dimple defines the Dimple characteristics.
	Dimple dimples.Dimple

	// Width is the width of the face the Dimples will be arranged for.
	Width float64
}

// EncodeFunction implements custom encoding for scad.EncodeFunction.
func (d Dimples) EncodeFunction() (interface{}, error) {
	children := make([]interface{}, len(dimplesPlacement))

	for i, rotation := range dimplesPlacement {
		dimples := scad.Wrap(
			dimples.Dimples{
				Count:  rotation.count,
				Dimple: d.Dimple,
				Width:  d.Width,
				Name:   scad.ModuleName(fmt.Sprintf("dimples_%d", rotation.count)),
			},
			transformation.RotateAround(
				float64(rotation.rotateA),
				float64(rotation.rotateV[0]),
				float64(rotation.rotateV[1]),
				float64(rotation.rotateV[2]),
			),
			transformation.TranslateTo(
				float64(rotation.translate[0])*d.Width/2,
				float64(rotation.translate[1])*d.Width/2,
				float64(rotation.translate[2])*d.Width/2,
			),
		)

		children[i] = dimples
	}

	return booleans.Union{
		Children: children,
	}, nil
}
