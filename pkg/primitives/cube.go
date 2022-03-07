package primitives

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/values"
)

// Cube is a cube.
type Cube struct {
	cube scad.AutoFunctionName

	// Only one of these size values may be set.
	Size    values.Float    `scad:"size"`
	SizeXYZ values.FloatXYZ `scad:"size"`

	Center values.Bool
}
