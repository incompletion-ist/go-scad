package transforms

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/values"
)

// Translate is a translate operation.
type Translate struct {
	translate scad.AutoFunctionName

	V        values.FloatXYZ `scad:"v"`
	Children []interface{}
}

// TranslateTo applies a translate operation.
func TranslateTo(x, y, z float64) scad.ChildWrapper {
	return func(i interface{}) interface{} {
		return Translate{
			V:        values.NewFloatXYZ(x, y, z),
			Children: []interface{}{i},
		}
	}
}
