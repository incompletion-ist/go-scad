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
	"go.incompletion.ist/scad/scad"
	"go.incompletion.ist/scad/value"
)

// Resize is a resize transform.
type Resize struct {
	NewSize value.FloatXYZ `scad:"newsize"`
	Auto    value.Bool     `scad:"auto"`

	Children []interface{}
}

// Wrap wraps a child with this Resize.
func (resize Resize) Wrap(child interface{}) scad.Wrapper {
	resize.Children = append([]interface{}{child}, resize.Children...)

	return resize
}
