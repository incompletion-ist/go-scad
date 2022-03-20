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

// Package value provides explicitly settable values to be used for OpenSCAD
// function calls. Only explicitly set values are passed as function arguments.
// The types in this package adhere to the scad.ParameterValueGetter interface.
package value

import (
	"fmt"
	"reflect"
	"strconv"
)

// explicitlyValuable is the union of types available for explicitValue.
type explicitlyValuable interface {
	string | bool | int | float64
}

// explicitValue stores a value for a explicitlyValuable type.
type explicitValue[T explicitlyValuable] struct {
	// value isn't actually unused, but the available golangci-lint version thinks it is
	value T //nolint:golint,structcheck
}

// newExplicitValue returns a pointer to a new explicitValue for the given value.
func newExplicitValue[T explicitlyValuable](value T) *explicitValue[T] {
	var e explicitValue[T]

	e.value = value

	return &e
}

// ValueOk returns the stored value and a boolean indicating if the value was set
// (not a nil pointer). If the pointer is nil, it returns the zero value for
// its stored type.
func (e *explicitValue[T]) ValueOk() (T, bool) {
	var value T

	if e != nil {
		value = e.value
	}

	return value, e != nil
}

// Value returns the stored value (or the zero value, if a nil pointer).
func (e *explicitValue[T]) Value() T {
	value, _ := e.ValueOk()

	return value
}

// GetParameterValue returns the string representation for the stored value. It
// always returns true, as this method is not on a pointer.
func (e explicitValue[T]) GetParameterValue() (string, bool) {
	storedValueV := reflect.ValueOf(e.value)
	storedValueK := storedValueV.Kind()

	if storedValueK == reflect.String {
		return fmt.Sprintf("%q", storedValueV.String()), true
	}

	if storedValueK == reflect.Int {
		return strconv.FormatInt(storedValueV.Int(), 10), true
	}

	if storedValueK == reflect.Float64 {
		return strconv.FormatFloat(storedValueV.Float(), 'f', -1, 64), true
	}

	if storedValueK == reflect.Bool {
		return strconv.FormatBool(storedValueV.Bool()), true
	}

	// this should be unreachable
	return "", false
}
