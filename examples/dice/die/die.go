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
	"github.com/micahkemp/scad/examples/dice/dimples"
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/booleans"
	"github.com/micahkemp/scad/pkg/scad/primitives"
	"github.com/micahkemp/scad/pkg/scad/transformation"
	"github.com/micahkemp/scad/pkg/scad/values"
)

// Die is a die.
type Die struct {
	Name scad.ModuleName `scad:"die"`

	// Dimple defines the Die's Dimple characteristics.
	Dimple dimples.Dimple

	// Width is the width of the Die's face.
	Width float64
}

// EncodeFunction implements custom encoding for scad.EncodeFunction.
func (d Die) EncodeFunction() (interface{}, error) {
	return scad.Apply(
		primitives.Cube{Size: values.NewFloat(d.Width)},
		transformation.Translate{
			V: values.NewFloatXYZ(-d.Width/2, -d.Width/2, -d.Width/2),
		},
		booleans.Difference{
			Children: []interface{}{
				Dimples{
					Dimple: d.Dimple,
					Width:  d.Width,
				},
			},
		},
	), nil
}

// DieSamples is a map of sample names to their sample Die.
var DieSamples = map[string]interface{}{
	"die_basic": Die{
		Dimple: dimples.Dimple{
			Depth:    2,
			Diameter: 10,
		},
		Width: 60,
	},
	"die_shallow_dimple": Die{
		Dimple: dimples.Dimple{
			Depth:    0.5,
			Diameter: 10,
		},
		Width: 60,
	},
}
