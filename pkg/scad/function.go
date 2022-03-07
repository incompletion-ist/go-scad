package scad

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Function describes a SCAD function call.
type Function struct {
	// ModuleName is the name of the module to create for the Function. If set
	// the Function will be created as its own SCAD module.
	ModuleName string

	// Name is the function name, such as "cube", "cylinder", "translate".
	Name string

	// Parameters is a map of parameter names to parameter values in string form.
	Parameters map[string]string

	// Children is a slice of Function objects that are children of the Function.
	Children []Function
}

// SetParameter sets the parameter with the given key to the given value. A boolean
// is returned indicating if the parameter was replaced (already had a value).
func (fn *Function) SetParameter(key string, value string) bool {
	if fn.Parameters == nil {
		fn.Parameters = map[string]string{}
	}

	_, replaced := fn.Parameters[key]

	fn.Parameters[key] = value

	return replaced
}

// parametersString returns the Function's Parameters as a string suitable
// to use when calling the function in .scad.
func (fn Function) parametersString() string {
	paramKeys := make([]string, 0, len(fn.Parameters))
	for key := range fn.Parameters {
		paramKeys = append(paramKeys, key)
	}
	sort.Strings(paramKeys)

	params := make([]string, len(fn.Parameters))
	for i, key := range paramKeys {
		params[i] = fmt.Sprintf("%s=%s", key, fn.Parameters[key])
	}

	return strings.Join(params, ", ")
}

// moduleFilename returns the module filename.
func (fn Function) moduleFilename() string {
	return fmt.Sprintf("%s.scad", fn.ModuleName)
}

// modulePath returns the relative path to the module.
func (fn Function) modulePath() string {
	return path.Join(fn.ModuleName, fn.moduleFilename())
}

// moduleUseString returns the "use" directive for the module.
func (fn Function) moduleUseString() string {
	return fmt.Sprintf("use <%s>", fn.modulePath())
}

// childModules returns a slice of Functions for all direct descendent Functions that
// are modules (non-empty ModuleName). A direct descendent is one that is not below
// another module.
func (fn Function) childModules() []Function {
	var modules []Function

	for _, child := range fn.Children {
		if child.ModuleName != "" {
			modules = append(modules, child)
		} else {
			modules = append(modules, child.childModules()...)
		}
	}

	return modules
}

// childUseStrings returns a slice of content strings containing the "use" directives
// needed by the Function. An error is returned if there are multiple non-identical
// modules with the same name.
func (fn Function) childUseStrings() ([]string, error) {
	seenModules := map[string]Function{}

	for _, module := range fn.childModules() {
		if seenModule, ok := seenModules[module.ModuleName]; ok {
			if !reflect.DeepEqual(module, seenModule) {
				return nil, fmt.Errorf("conflicting module name: %s", seenModule.ModuleName)
			}
		} else {
			seenModules[module.ModuleName] = module
		}
	}

	seenModuleNames := make([]string, 0, len(seenModules))
	for moduleName := range seenModules {
		seenModuleNames = append(seenModuleNames, moduleName)
	}
	sort.Strings(seenModuleNames)

	chUseStrings := make([]string, len(seenModuleNames))
	for i, moduleName := range seenModuleNames {
		chUseStrings[i] = seenModules[moduleName].moduleUseString()
	}

	return chUseStrings, nil
}

// moduleCallString returns the directive to call the module.
func (fn Function) moduleCallString() string {
	return fmt.Sprintf("%s();", fn.ModuleName)
}

// moduleContentStrings returns the content of a module, surrounding the function's
// content by the module() { } syntax.
func (fn Function) moduleContentStrings() []string {
	var mdStrings []string

	mdStrings = append(mdStrings, fmt.Sprintf("module %s() {", fn.ModuleName))
	for _, fnString := range fn.functionCallStrings() {
		mdStrings = append(mdStrings, fmt.Sprintf("  %s", fnString))
	}
	mdStrings = append(mdStrings, fmt.Sprintf("}"))

	// add module call at the end so any module can be opened directly in OpenSCAD
	mdStrings = append(mdStrings, fn.moduleCallString())

	return mdStrings
}

// functionCallStrings returns the content of a Function, including its children content.
func (fn Function) functionCallStrings() []string {
	var fnClose string = ";"
	var childrenClose []string
	if len(fn.Children) > 0 {
		fnClose = " {"
		childrenClose = []string{"}"}
	}

	fnCallStrings := []string{
		fmt.Sprintf("%s(%s)%s", fn.Name, fn.parametersString(), fnClose),
	}

	fnStrings := []string{}
	fnStrings = append(fnStrings, fnCallStrings...)

	for _, child := range fn.Children {
		for _, childString := range child.callStrings() {
			fnStrings = append(fnStrings, fmt.Sprintf("  %s", childString))
		}
	}
	fnStrings = append(fnStrings, childrenClose...)

	return fnStrings
}

func (fn Function) callStrings() []string {
	if fn.ModuleName != "" {
		return []string{fn.moduleCallString()}
	}

	return fn.functionCallStrings()
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

// EncodeFunction encodes an interface into a Function. The given interface must be a struct.
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
// • Exactly one FunctionNameGetter field is not found
//
// • Multiple field set a value for the same Parameter key (multiple fields can have the same key name,
//   as long no more than one sets a value)
//
// • Multiple Children fields are found
func EncodeFunction(i interface{}) (Function, error) {
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
				child, err := EncodeFunction(fieldV.Index(i).Interface())
				if err != nil {
					return Function{}, err
				}

				children[i] = child
			}
			fn.Children = children
		}
	}

	if fn.Name == "" {
		return Function{}, fmt.Errorf("scad: attempted to encode type (%T) without FunctionNameGetter field", i)
	}

	return fn, nil
}
