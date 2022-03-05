package scad

// AutoFunctionName is an empty type that is used to use a struct field's name
// or "scad" tag value as the Name value when encoding to a Function. It enables
// this by always returning an empty string for GetFunctionName.
type AutoFunctionName struct{}

// GetFunctionName returns an empty string.
func (fnName AutoFunctionName) GetFunctionName() string {
	return ""
}
