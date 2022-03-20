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

package extrusion

import (
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/value"
)

// LinearExtrude is a linear extrude.
type LinearExtrude struct {
	linearExtrude scad.AutoFunctionName `scad:"linear_extrude"` //nolint:golint,structcheck,unused

	Height value.Float `scad:"height"`
	Twist  value.Float `scad:"twist"`
	Center value.Bool  `scad:"center"`
	Slices value.Int   `scad:"slices"`

	// Only one of Scale or ScaleXY should be set.
	Scale   value.Float   `scad:"scale"`
	ScaleXY value.FloatXY `scad:"scale"`

	FN value.Int `scad:"$fn"`

	Children []interface{}
}

// Wrap wraps a child with this LinearExtrude.
func (extrude LinearExtrude) Wrap(child interface{}) scad.Wrapper {
	extrude.Children = append([]interface{}{child}, extrude.Children...)

	return extrude
}
