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

// Rotate is a rotate transform.
type Rotate struct {
	// Only one of A, Axyz may be set.
	A    value.Float    `scad:"a"`
	Axyz value.FloatXYZ `scad:"a"`

	V value.FloatXYZ `scad:"v"`

	Children []interface{}
}

// Wrap wraps a child with this Rotate.
func (rotate Rotate) Wrap(child interface{}) scad.Wrapper {
	rotate.Children = append([]interface{}{child}, rotate.Children...)

	return rotate
}
