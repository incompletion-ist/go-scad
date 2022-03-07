package dimples

import (
	"fmt"

	"github.com/micahkemp/scad/pkg/booleans"
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/transforms"
)

// dimplePositions defines the placement of dimples from 0 to 7 dimples on a 2-dimensional face.
// The integer values determine which direction the dimple is shifted on the X and Y axes.
var dimplePositions = [7][][2]int{
	{},
	{
		{0, 0},
	},
	{
		{-1, -1},
		{1, 1},
	},
	{
		{-1, -1},
		{0, 0},
		{1, 1},
	},
	{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	},
	{
		{-1, -1},
		{-1, 1},
		{0, 0},
		{1, -1},
		{1, 1},
	},
	{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	},
}

// Dimples defines a collection of Dimple objects.
type Dimples struct {
	Name scad.ModuleName `scad:"dimples"`

	// Dimple defines the Dimple characteristics.
	Dimple Dimple

	// Width is the width of the 2-dimensional face the Dimples will occupy.
	Width float64

	// Count is how many Dimples will be placed. Count must be between 0 and 6.
	Count int
}

// EncodeFunction implements custom encoding for scad.EncodeFunction.
func (d Dimples) EncodeFunction() (interface{}, error) {
	if d.Count < 0 || d.Count > 6 {
		return nil, fmt.Errorf("dimples: Dimples Count is %d, must be between 0 and 6", d.Count)
	}

	dimpleLayout := dimplePositions[d.Count]
	dimples := make([]interface{}, d.Count)

	for i, dimplePosition := range dimpleLayout {
		dimples[i] = scad.Wrap(
			d.Dimple,
			transforms.TranslateTo(
				float64(dimplePosition[0])*(d.Width/4),
				float64(dimplePosition[1])*(d.Width/4),
				0,
			),
		)
	}

	return booleans.Union{
		Children: dimples,
	}, nil
}
