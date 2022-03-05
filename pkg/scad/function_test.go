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
			name:  "empty struct",
			input: struct{}{},
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
