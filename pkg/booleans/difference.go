package booleans

import "github.com/micahkemp/scad/pkg/scad"

// Difference is a difference boolean operation.
type Difference struct {
	difference scad.AutoFunctionName

	Children []interface{}
}

// DifferenceWith applies a difference operation.
func DifferenceWith(with ...interface{}) scad.ChildWrapper {
	return func(i interface{}) interface{} {
		children := make([]interface{}, 0, len(with)+1)
		children = append(children, i)
		children = append(children, with...)

		return Difference{
			Children: children,
		}
	}
}
