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

package scad

// AutoFunctionName is an empty type that is used to use a struct field's name
// or "scad" tag value as the Name value when encoding to a Function. It enables
// this by always returning an empty string for GetFunctionName.
type AutoFunctionName struct{}

// GetFunctionName returns an empty string.
func (fnName AutoFunctionName) GetFunctionName() string {
	return ""
}
