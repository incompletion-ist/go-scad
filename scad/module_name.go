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

// ModuleName represents an OpenSCAD module's name. Its presence in a
// struct causes that type to be split into its own module.
type ModuleName string

// GetModuleName returns the value of the ModuleName. It satisfies the
// ModuleNameGetter interface.
func (name ModuleName) GetModuleName() string {
	return string(name)
}
