package primitives

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/values"
)

// Sphere is a sphere.
type Sphere struct {
	sphere scad.AutoFunctionName //nolint:golint,structcheck,unused

	// Only one of R or D should be set.
	R values.Float `scad:"r"`
	D values.Float `scad:"d"`
}
