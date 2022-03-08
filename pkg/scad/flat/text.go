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

package flat

import (
	"github.com/micahkemp/scad/pkg/scad/values"
)

// Text is text.
type Text struct {
	Text      values.String `scad:"text"`
	Size      values.Float  `scad:"size"`
	Font      values.String `scad:"font"`
	Halign    values.String `scad:"halign"`
	Valign    values.String `scad:"valign"`
	Spacing   values.Float  `scad:"spacing"`
	Direction values.String `scad:"direction"`
	Language  values.String `scad:"language"`
	Script    values.String `scad:"script"`
}
