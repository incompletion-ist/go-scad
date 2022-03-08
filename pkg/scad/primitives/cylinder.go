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

package primitives

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/values"
)

// Cylinder is a cylinder.
type Cylinder struct {
	cylinder scad.AutoFunctionName //nolint:golint,structcheck,unused

	H      values.Float `scad:"h"`
	R      values.Float `scad:"r"`
	R1     values.Float `scad:"r1"`
	R2     values.Float `scad:"r2"`
	D      values.Float `scad:"d"`
	D1     values.Float `scad:"d1"`
	D2     values.Float `scad:"d2"`
	Center values.Bool  `scad:"center"`
}
