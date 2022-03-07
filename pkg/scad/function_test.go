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

func TestFunction_callStrings(t *testing.T) {
	tests := []struct {
		name  string
		input Function
		want  []string
	}{
		{
			name:  "empty",
			input: Function{},
			want: []string{
				"();",
			},
		},
		{
			name: "name only",
			input: Function{
				Name: "cube",
			},
			want: []string{
				"cube();",
			},
		},
		{
			name: "name and params",
			input: Function{
				Name: "cube",
				Parameters: map[string]string{
					"center": "true",
					"size":   "10",
				},
			},
			want: []string{
				"cube(center=true, size=10);",
			},
		},
		{
			name: "name, params, and children",
			input: Function{
				Name: "translate",
				Parameters: map[string]string{
					"v": "[1, 2, 3]",
				},
				Children: []Function{
					{
						Name: "cube",
						Parameters: map[string]string{
							"size": "10",
						},
					},
				},
			},
			want: []string{
				"translate(v=[1, 2, 3]) {",
				"  cube(size=10);",
				"}",
			},
		},
	}

	for _, test := range tests {
		got := test.input.functionCallStrings()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q contentStrings() got\n%#v, want\n%#v", test.name, got, test.want)
		}
	}
}

func TestFunction_moduleContentStrings(t *testing.T) {
	tests := []struct {
		name  string
		input Function
		want  []string
	}{
		{
			// the empty Function's moduleContentStrings is nonsense, but test it for consistency
			name:  "empty",
			input: Function{},
			want: []string{
				"module () {",
				"  ();",
				"}",
				"();",
			},
		},
		{
			name: "module and function names",
			input: Function{
				ModuleName: "testModule",
				Name:       "testFunction",
			},
			want: []string{
				"module testModule() {",
				"  testFunction();",
				"}",
				"testModule();",
			},
		},
	}

	for _, test := range tests {
		got := test.input.moduleContentStrings()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q moduleContentStrings() got\n%#v, want\n%#v", test.name, got, test.want)
		}
	}
}

func TestFunction_childModules(t *testing.T) {
	tests := []struct {
		name  string
		input Function
		want  []Function
	}{
		{
			name:  "empty",
			input: Function{},
			want:  nil,
		},
		{
			name:  "no children",
			input: Function{ModuleName: "topModule"},
			want:  nil,
		},
		{
			name: "one level",
			input: Function{
				ModuleName: "topModule",
				Children: []Function{
					{ModuleName: "child1Module"},
					{ModuleName: "child2Module"},
				},
			},
			want: []Function{
				{ModuleName: "child1Module"},
				{ModuleName: "child2Module"},
			},
		},
		{
			name: "child modules but under a module",
			input: Function{
				Children: []Function{
					{ModuleName: "child1Module", Children: []Function{{ModuleName: "child1AModule"}}},
				},
			},
			want: []Function{
				{ModuleName: "child1Module", Children: []Function{{ModuleName: "child1AModule"}}},
			},
		},
		{
			name: "child modules",
			input: Function{
				Children: []Function{
					{Children: []Function{{ModuleName: "child1AModule"}}},
				},
			},
			want: []Function{
				{ModuleName: "child1AModule"},
			},
		},
		{
			name: "multi-level nested modules",
			input: Function{
				Children: []Function{
					{Children: []Function{
						{Children: []Function{
							{ModuleName: "child1AModule"},
						}},
					}},
				},
			},
			want: []Function{
				{ModuleName: "child1AModule"},
			},
		},
	}

	for _, test := range tests {
		got := test.input.childModules()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q, childModules() got\n%#v, want\n%#v", test.name, got, test.want)
		}
	}
}

func TestFunction_childUseStrings(t *testing.T) {
	tests := []struct {
		name      string
		input     Function
		want      []string
		wantError bool
	}{
		{
			name:  "empty",
			input: Function{},
			want:  []string{},
		},
		{
			name:  "no children",
			input: Function{ModuleName: "topModule"},
			want:  []string{},
		},
		{
			name: "one level",
			input: Function{
				ModuleName: "topModule",
				Children: []Function{
					{ModuleName: "child1Module"},
					{ModuleName: "child2Module"},
				},
			},
			want: []string{
				"use <child1Module/child1Module.scad>",
				"use <child2Module/child2Module.scad>",
			},
		},
		{
			name: "child modules but under a module",
			input: Function{
				ModuleName: "topModule",
				Children: []Function{
					{ModuleName: "child1Module", Children: []Function{{ModuleName: "child1AModule"}}},
				},
			},
			want: []string{
				"use <child1Module/child1Module.scad>",
			},
		},
		{
			name: "child modules",
			input: Function{
				Children: []Function{
					{Children: []Function{{ModuleName: "child1AModule"}}},
				},
			},
			want: []string{
				"use <child1AModule/child1AModule.scad>",
			},
		},
		{
			name: "multi-level nested modules",
			input: Function{
				Children: []Function{
					{Children: []Function{
						{Children: []Function{
							{ModuleName: "child1AModule"},
						}},
					}},
				},
			},
			want: []string{
				"use <child1AModule/child1AModule.scad>",
			},
		},
		{
			name: "module name conflicts",
			input: Function{
				Children: []Function{
					{ModuleName: "my_module", Name: "cube"},
					// the offending ModuleName is detected even when nested
					{Children: []Function{
						{ModuleName: "my_module", Name: "cylinder"},
					}},
				},
			},
			wantError: true,
		},
	}

	for _, test := range tests {
		got, err := test.input.childUseStrings()
		gotError := err != nil

		if gotError != test.wantError {
			t.Errorf("%q childUseStrings returned error? %v (%s)", test.name, gotError, err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%q childUseStrings got\n%#v, want\n%#v", test.name, got, test.want)
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

	cube AutoFunctionName //nolint:golint,structcheck,unused
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
			t.Errorf("%q EncodeFunction() returned error? %v (%s)", test.name, gotError, err)
		}

		if !reflect.DeepEqual(gotFunction, test.wantFunction) {
			t.Errorf("%q EncodeFunction() got\n%#v, want\n%#v", test.name, gotFunction, test.wantFunction)
		}
	}
}
