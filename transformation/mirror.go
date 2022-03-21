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

// Mirror is a mirror transformation.
type Mirror struct {
	V value.FloatXYZ

	Children []interface{}
}

// Wrap wraps a child with this Mirror.
func (mirror Mirror) Wrap(child interface{}) scad.Wrapper {
	mirror.Children = append([]interface{}{child}, mirror.Children...)

	return mirror
}
