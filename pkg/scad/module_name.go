package scad

// ModuleName represents an OpenSCAD module's name. Its presence in a
// struct causes that type to be split into its own module.
type ModuleName string

// GetModuleName returns the value of the ModuleName. It satisfies the
// ModuleNameGetter interface.
func (name ModuleName) GetModuleName() string {
	return string(name)
}
