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

package primitive3d

import (
	"go.incompletion.ist/go-scad/value"
)

// Cylinder is a cylinder.
type Cylinder struct {
	H      value.Float `scad:"h"`
	R      value.Float `scad:"r"`
	R1     value.Float `scad:"r1"`
	R2     value.Float `scad:"r2"`
	D      value.Float `scad:"d"`
	D1     value.Float `scad:"d1"`
	D2     value.Float `scad:"d2"`
	Center value.Bool  `scad:"center"`

	FA value.Float `scad:"$fa"`
	FS value.Float `scad:"$fs"`
	FN value.Int   `scad:"$fn"`
}
