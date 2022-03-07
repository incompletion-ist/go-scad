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

package dimples

import (
	"math"

	"github.com/micahkemp/scad/pkg/primitives"
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/transforms"
	"github.com/micahkemp/scad/pkg/values"
)

// Dimple is a sphere that will be removed from a Die.
type Dimple struct {
	Name scad.ModuleName `scad:"dimple"`

	// Depth is how deep the dimple should be.
	Depth float64

	// Diameter is how wide the dimple should be where it intersects the Die.
	Diameter float64
}

// EncodeFunction implements custom encoding for scad.EncodeFunction.
func (d Dimple) EncodeFunction() (interface{}, error) {
	sphereRadius := (math.Pow(d.Depth, 2) + math.Pow(d.Diameter/2, 2)) / (2 * d.Depth)

	return scad.Wrap(
		primitives.Sphere{R: values.NewFloat(sphereRadius)},
		transforms.TranslateTo(0, 0, sphereRadius-d.Depth),
		transforms.RotateAround(180, 1, 0, 0),
	), nil
}
