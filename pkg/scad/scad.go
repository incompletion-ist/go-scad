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

// Package scad enables abstraction of OpenSCAD code in Go.
package scad

import (
	"fmt"
	"reflect"
	"strings"
)

// FunctionContent returns the OpenSCAD content for an input interface.
func FunctionContent(i interface{}) (string, error) {
	fn, err := Encode(i)
	if err != nil {
		return "", err
	}

	return fn.content()
}

// Write writes a given interface as a Function to the given location.
func Write(p string, i interface{}) error {
	fn, err := Encode(i)
	if err != nil {
		return err
	}

	return fn.Write(p)
}

// WriteMap writes each interface as a Function to a path of the key.
func WriteMap(samples map[string]interface{}) error {
	for name, sample := range samples {
		if err := Write(name, sample); err != nil {
			return err
		}
	}

	return nil
}

// FunctionNameGetter is the interface for types that implement GetFunctionName.
type FunctionNameGetter interface {
	// GetFunctionName returns a string representing the function name. The returned
	// string may be empty.
	GetFunctionName() string
}

// ParameterValueGetter is the interface for types that implement GetParameterValue.
type ParameterValueGetter interface {
	// GetParameterValue returns a string representing the value, and a boolean indicating
	// if the value is set.
	GetParameterValue() (string, bool)
}

// ModuleNameGetter is the interface for types that implement GetModuleName.
type ModuleNameGetter interface {
	GetModuleName() string
}

// SCADEncoder is the interface for types that implement EncodeSCAD.
type SCADEncoder interface {
	// EncodeSCAD returns a new interface that should be passed to scad.EncodeSCAD, enabling
	// overriding of encoding functionality, or more generally to permit an arbitrary type to
	// define what object should be encoded in its place.
	EncodeSCAD() (interface{}, error)
}

// Encode encodes an interface into a scad.Function. The given interface must be a struct.
//
// The resulting Function will have values applied based on the struct fields implementing
// one or more of these interfaces:
//
// FunctionNameGetter fields will determine the Function's Name value as the first non-empty
// value of these options:
//
// •Return value of the field's GetFunctionName method
//
// •The value of the field's "scad" tag
//
// •The lowercased field name
//
// •The lowercased type name
//
// ParameterValueGetter fields will set Parameter values for the Function. The Parameter's
// key will be the last non-empty value of these options:
//
// •The value of the field's "scad" tag
//
// •The lowercased field name
//
// Slice fields set the Children values for the Function.
//
// An error will be returned if:
//
// • Name is empty after encoding
//
// • Multiple field set a value for the same Parameter key (multiple fields can have the same key name,
// as long no more than one sets a value)
//
// • Multiple Children fields are found
func Encode(i interface{}) (Function, error) {
	var fn Function

	if i == nil {
		return Function{}, fmt.Errorf("scad: attempted to encode null value to Function")
	}

	// be nice and dereference pointers
	iV := reflect.ValueOf(i)
	if iV.Kind() == reflect.Ptr {
		iV = iV.Elem()
	}
	iT := iV.Type()

	if iV.Kind() != reflect.Struct {
		return Function{}, fmt.Errorf("scad: attempted to encode non-struct (%T) to Function", i)
	}

	for i := 0; i < iT.NumField(); i++ {
		field := iT.Field(i)

		scadName := field.Tag.Get("scad")
		if scadName == "" {
			scadName = strings.ToLower(field.Name)
		}

		fieldV := iV.Field(i)
		fieldT := fieldV.Type()

		// ModuleName
		if fieldT.Implements(reflect.TypeOf((*ModuleNameGetter)(nil)).Elem()) {
			fn.ModuleName = scadName

			if field.IsExported() {
				gotMN := fieldV.Interface().(ModuleNameGetter).GetModuleName()
				if gotMN != "" {
					fn.ModuleName = gotMN
				}

			}
		}

		// Name
		if fieldT.Implements(reflect.TypeOf((*FunctionNameGetter)(nil)).Elem()) {
			if fn.Name != "" {
				return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple FunctionNameGetter fields", i)
			}

			fnName := scadName

			if field.IsExported() {
				gotFnName := fieldV.Interface().(FunctionNameGetter).GetFunctionName()
				if gotFnName != "" {
					fnName = gotFnName
				}
			}

			fn.Name = fnName
		}

		// Parameters
		if fieldT.Implements(reflect.TypeOf((*ParameterValueGetter)(nil)).Elem()) {
			// silently ignore unexported ParameterValueGetter fields
			if !field.IsExported() {
				continue
			}

			if gotValue, ok := fieldV.Interface().(ParameterValueGetter).GetParameterValue(); ok {
				if replaced := fn.SetParameter(scadName, gotValue); replaced {
					return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple ParameterValueGetter fields with the same name: %s", i, scadName)
				}
			}
		}

		// Children
		if fieldV.Kind() == reflect.Slice {
			// silently ignore unexported slices
			if !field.IsExported() {
				continue
			}

			if fn.Children != nil {
				return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with multiple slice fields", i)
			}

			children := make([]Function, fieldV.Len())

			for i := 0; i < fieldV.Len(); i++ {
				child, err := Encode(fieldV.Index(i).Interface())
				if err != nil {
					return Function{}, err
				}

				children[i] = child
			}
			fn.Children = children
		}
	}

	// after all that, if the given interface is a FunctionEncoder, undo everything except module name
	if iT.Implements(reflect.TypeOf((*SCADEncoder)(nil)).Elem()) {
		encodeFn, err := iV.Interface().(SCADEncoder).EncodeSCAD()
		if err != nil {
			return Function{}, err
		}

		encoderFn, err := Encode(encodeFn)
		if err != nil {
			return Function{}, err
		}

		encoderFn.ModuleName = fn.ModuleName

		return encoderFn, nil
	}

	if fn.Name == "" {
		fn.Name = strings.ToLower(iT.Name())
	}

	if fn.Name == "" {
		return Function{}, fmt.Errorf("scad: attempted to encode type (%T) with empty name", i)
	}

	return fn, nil
}
