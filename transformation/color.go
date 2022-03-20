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

package transformation

import (
	"go.incompletion.ist/go-scad/scad"
	"go.incompletion.ist/go-scad/value"
)

// Color is a color.
type Color struct {
	C     value.FloatXYZ `scad:"c"`
	Alpha value.Float    `scad:"alpha"`

	Children []interface{}
}

// Wrap wraps a child with this Color.
func (color Color) Wrap(child interface{}) scad.Wrapper {
	color.Children = append([]interface{}{child}, color.Children...)

	return color
}
