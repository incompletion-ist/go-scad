package transforms

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/values"
)

// Rotate is a rotate transform.
type Rotate struct {
	rotate scad.AutoFunctionName

	// Only one of A, Axyz may be set.
	A    values.Float    `scad:"a"`
	Axyz values.FloatXYZ `scad:"a"`

	V values.FloatXYZ `scad:"v"`

	Children []interface{}
}

// RotateAround performs a Rotate.
func RotateAround(a, vX, vY, vZ float64) scad.ChildWrapper {
	return func(i interface{}) interface{} {
		return Rotate{
			A:        values.NewFloat(a),
			V:        values.NewFloatXYZ(vX, vY, vZ),
			Children: []interface{}{i},
		}
	}
}
