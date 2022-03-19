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

import (
	"fmt"
	"os"
	"path"
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

// uniqueChildModules returns the deduplicated output of childModules.
func (fn Function) uniqueChildModules() ([]Function, error) {
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

	modules := make([]Function, len(seenModuleNames))
	for i, moduleName := range seenModuleNames {
		modules[i] = seenModules[moduleName]
	}

	return modules, nil
}

// childUseStrings returns a slice of content strings containing the "use" directives
// needed by the Function. An error is returned if there are multiple non-identical
// modules with the same name.
func (fn Function) childUseStrings() ([]string, error) {
	modules, err := fn.uniqueChildModules()
	if err != nil {
		return nil, err
	}

	chUseStrings := make([]string, len(modules))
	for i, module := range modules {
		chUseStrings[i] = module.moduleUseString()
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
	mdStrings = append(mdStrings, "}")

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

// callStrings returns the content for a Function, which will either be calling the module
// by name (if ModuleName is set) or the Function's call contents.
func (fn Function) callStrings() []string {
	if fn.ModuleName != "" {
		return []string{fn.moduleCallString()}
	}

	return fn.functionCallStrings()
}

// fileContentStrings returns all content lines for the Function that are needed
// when writing it to a file. This will include the "use" directives at the top,
// the module content, and calling the module at the end (so any module file can
// be opened in OpenSCAD and viewed properly on its own).
func (fn Function) fileContentStrings() ([]string, error) {
	fStrings := []string{}

	chUseStrings, err := fn.childUseStrings()
	if err != nil {
		return nil, err
	}
	fStrings = append(fStrings, chUseStrings...)

	if fn.ModuleName == "" {
		fStrings = append(fStrings, fn.functionCallStrings()...)
	} else {
		fStrings = append(fStrings, fn.moduleContentStrings()...)
	}

	// trailing newline
	fStrings = append(fStrings, "")

	return fStrings, nil
}

// content returns the Function's content as one large string, suitable to be
// written to a file.
func (fn Function) content() (string, error) {
	contentStrings, err := fn.fileContentStrings()
	if err != nil {
		return "", err
	}

	return strings.Join(contentStrings, "\n"), nil
}

// Write writes the module at the given path. Any nested modules will be written
// at paths relative to this one. For a Function to be successfully able to Write,
// it must be a module (non-empty ModuleName).
func (fn Function) Write(p string) error {
	if fn.ModuleName == "" {
		return fmt.Errorf("attempted Write on non-Module")
	}

	content, err := fn.content()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(p, 0777); err != nil {
		return err
	}

	writeP := path.Join(p, fn.moduleFilename())
	writeF, err := os.OpenFile(writeP, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	if _, err := writeF.Write([]byte(content)); err != nil {
		return err
	}

	for _, childModule := range fn.childModules() {
		childPath := path.Join(p, childModule.ModuleName)

		if err := childModule.Write(childPath); err != nil {
			return err
		}
	}

	return nil
}
