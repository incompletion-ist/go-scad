package scad

import (
	"reflect"
	"testing"
)

func TestFunction_parametersString(t *testing.T) {
	tests := []struct {
		name  string
		input Function
		want  string
	}{
		{
			name:  "empty",
			input: Function{},
			want:  "",
		},
		{
			name: "many",
			input: Function{
				Parameters: map[string]string{
					// intentionally misordered, but maps have no order anyway
					"paramB": "valueB",
					"paramA": "valueA",
					"paramC": "valueC",
				},
			},
			want: "paramA=valueA, paramB=valueB, paramC=valueC",
		},
	}

	for _, test := range tests {
		got := test.input.parametersString()

		if got != test.want {
			t.Errorf("%q parametersString() got\n%s, want\n%s", test.name, got, test.want)
		}
	}
}
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
			name: "untagged module name",
			input: struct {
				myModule   ModuleName
				myFunction AutoFunctionName
			}{},
			wantFunction: Function{
				ModuleName: "mymodule",
				Name:       "myfunction",
			},
		},
		{
			name: "tagged module name",
			input: struct {
				myModule   ModuleName `scad:"my_module"`
				myFunction AutoFunctionName
			}{},
			wantFunction: Function{
				ModuleName: "my_module",
				Name:       "myfunction",
			},
		},
		{
			name: "overridden module name",
			input: struct {
				// MyModule must be exported to be overridden
				MyModule   ModuleName `scad:"my_module"`
				myFunction AutoFunctionName
			}{
				MyModule: "my_customized_module_name",
			},
			wantFunction: Function{
				ModuleName: "my_customized_module_name",
				Name:       "myfunction",
			},
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
		{
			name: "children",
			input: struct {
				translate AutoFunctionName
				Children  []interface{}
			}{
				Children: []interface{}{
					testFunction{
						Length: testParameterValueGetter{value: "20", explicit: true},
					},
					testFunction{
						Length: testParameterValueGetter{value: "10", explicit: true},
						Center: testParameterValueGetter{value: "true", explicit: true},
					},
				},
			},
			wantFunction: Function{
				Name: "translate",
				Children: []Function{
					{
						Name: "cube",
						Parameters: map[string]string{
							"x": "20",
						},
					},
					{
						Name: "cube",
						Parameters: map[string]string{
							"x":      "10",
							"center": "true",
						},
					},
				},
			},
		},
		{
			name: "multiple children fields",
			input: struct {
				translate    AutoFunctionName
				Children     []interface{}
				MoreChildren []interface{}
			}{},
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
