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

package transforms

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/values"
)

// Resize is a resize transform.
type Resize struct {
	resize scad.AutoFunctionName //nolint:golint,structcheck,unused

	NewSize values.FloatXYZ `scad:"newsize"`
	Auto    values.Bool     `scad:"auto"`

	Children []interface{}
}
