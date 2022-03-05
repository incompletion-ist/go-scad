package scad

import (
	"reflect"
	"testing"
)

type testParameterValueGetter struct {
	value    string
	explicit bool
}

func (testGetter testParameterValueGetter) GetParameterValue() (string, bool) {
	return testGetter.value, testGetter.explicit
}

type testFunction struct {
	Length testParameterValueGetter `scad:"x"`
	Center testParameterValueGetter

	cube AutoFunctionName
}

func Test_EncodeFunction(t *testing.T) {
	tests := []struct {
		name         string
		input        interface{}
		wantFunction Function
		wantError    bool
	}{
		{
			name:      "nil",
			wantError: true,
		},
		{
			name:      "empty struct",
			input:     struct{}{},
			wantError: true,
		},
		{
			name:  "unset empty",
			input: testFunction{},
			wantFunction: Function{
				Name: "cube",
			},
		},
		{
			name: "set empty",
			input: testFunction{
				Length: testParameterValueGetter{
					explicit: true,
				},
				Center: testParameterValueGetter{
					explicit: true,
				},
			},
			wantFunction: Function{
				Name: "cube",
				Parameters: map[string]string{
					"x":      "",
					"center": "",
				},
			},
		},
		{
			name: "set value",
			input: testFunction{
				Length: testParameterValueGetter{
					value:    "xValue",
					explicit: true,
				},
				Center: testParameterValueGetter{
					value:    "centerValue",
					explicit: true,
				},
			},
			wantFunction: Function{
				Name: "cube",
				Parameters: map[string]string{
					"x":      "xValue",
					"center": "centerValue",
				},
			},
		},
		{
			name: "multiple FunctionNameGetter",
			input: struct {
				cube     AutoFunctionName
				cylinder AutoFunctionName
			}{},
			wantError: true,
		},
		{
			name: "parameter collision",
			input: struct {
				cube AutoFunctionName
				X    testParameterValueGetter `scad:"size"`
				Size testParameterValueGetter `scad:"size"`
			}{
				X:    testParameterValueGetter{value: "xValue", explicit: true},
				Size: testParameterValueGetter{value: "sizeValue", explicit: true},
			},
			wantError: true,
		},
	}

	for _, test := range tests {
		gotFunction, err := EncodeFunction(test.input)
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%q EncodeFunction returned error? %v (%s)", test.name, gotError, err)
		}

		if !reflect.DeepEqual(gotFunction, test.wantFunction) {
			t.Errorf("%q EncodeFunction got\n%#v, want\n%#v", test.name, gotFunction, test.wantFunction)
		}
	}
}
