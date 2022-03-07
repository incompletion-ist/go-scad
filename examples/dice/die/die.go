package die

import (
	"github.com/micahkemp/scad/examples/dice/dimples"
	"github.com/micahkemp/scad/pkg/booleans"
	"github.com/micahkemp/scad/pkg/primitives"
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/transforms"
	"github.com/micahkemp/scad/pkg/values"
)

// Die is a die.
type Die struct {
	Name scad.ModuleName `scad:"die"`

	// Dimple defines the Die's Dimple characteristics.
	Dimple dimples.Dimple

	// Width is the width of the Die's face.
	Width float64
}

// EncodeFunction implements custom encoding for scad.EncodeFunction.
func (d Die) EncodeFunction() (interface{}, error) {
	return scad.Wrap(
		primitives.Cube{Size: values.NewFloat(d.Width)},
		transforms.TranslateTo(-d.Width/2, -d.Width/2, -d.Width/2),
		booleans.DifferenceWith(Dimples{
			Dimple: d.Dimple,
			Width:  d.Width,
		}),
	), nil
}
