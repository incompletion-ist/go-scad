package booleans

import "github.com/micahkemp/scad/pkg/scad"

// Union is a union boolean operation.
type Union struct {
	union scad.AutoFunctionName //nolint:golint,structcheck,unused

	Children []interface{}
}

// UnionedWith applies a Union operation.
func UnionedWith(additions ...interface{}) scad.ChildWrapper {
	return func(i interface{}) interface{} {
		children := make([]interface{}, len(additions)+1)
		children = append(children, i)
		children = append(children, additions)

		return Union{
			Children: children,
		}
	}
}
