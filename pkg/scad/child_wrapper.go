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

// ChildWrapper is a function that wraps an interface within another
// interface.
type ChildWrapper func(interface{}) interface{}

// Wrap applies a set of ChildWrapper functions create a new interface
// from the input interface.
func Wrap(i interface{}, wrappers ...ChildWrapper) interface{} {
	for _, wrapper := range wrappers {
		i = wrapper(i)
	}

	return i
}

// Wrapper is the interface for types that implement Wrap.
type Wrapper interface {
	// Wrap returns a new Wrapper by wrapping the given interface.
	Wrap(interface{}) Wrapper
}

// Apply returns a new interface by wrapping the child with each given Wrapper in order.
//
// The returned value is an interface{}, not Wrapper, because is is possible for wrappers
// to be empty, causing child to be returned directly.
func Apply(child interface{}, wrappers ...Wrapper) interface{} {
	wrapped := child

	for _, wrapper := range wrappers {
		wrapped = wrapper.Wrap(wrapped)
	}

	return wrapped
}
