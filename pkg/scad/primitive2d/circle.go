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

package primitive2d

import "github.com/micahkemp/scad/pkg/scad/value"

// Circle is a circle.
type Circle struct {
	// Only one of R or D should be set.
	R value.Float `scad:"r"`
	D value.Float `scad:"d"`

	FA value.Float `scad:"$fa"`
	FS value.Float `scad:"$fs"`
	FN value.Int   `scad:"$fn"`
}
