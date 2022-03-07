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

package booleans

import "github.com/micahkemp/scad/pkg/scad"

// Union is a union boolean operation.
type Union struct {
	union scad.AutoFunctionName //nolint:golint,structcheck,unused

	Children []interface{}
}

// UnionedWith applies a Union operation.
func UnionedWith(additions ...interface{}) scad.ChildWrapper {
	return func(i interface{}) interface{} {
		children := make([]interface{}, len(additions)+1)
		children = append(children, i)
		children = append(children, additions)

		return Union{
			Children: children,
		}
	}
}
