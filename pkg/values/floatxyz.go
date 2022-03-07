package values

import (
	"fmt"
	"strconv"
	"strings"
)

// FloatXYZ represents a tuple of XYZ float64 values that can be explicitly
// set.
type FloatXYZ struct {
	valueX float64
	valueY float64
	valueZ float64
	set    bool
}

// Set explicitly sets the given values.
func (xyz *FloatXYZ) Set(x, y, z float64) {
	xyz.valueX = x
	xyz.valueY = y
	xyz.valueZ = z
	xyz.set = true
}

// IsSet returns a boolean indicating if FloatXYZ has been explicitly set.
func (xyz FloatXYZ) IsSet() bool {
	return xyz.set
}

// ValueX returns the stored value for X.
func (xyz FloatXYZ) ValueX() float64 {
	return xyz.valueX
}

// ValueX returns the stored value for Y.
func (xyz FloatXYZ) ValueY() float64 {
	return xyz.valueY
}

// ValueX returns the stored value for Z.
func (xyz FloatXYZ) ValueZ() float64 {
	return xyz.valueZ
}

// GetParameterValue returns a string value for the FloatXYZ, and a boolean
// indicating if its value was explicity set.
func (xyz FloatXYZ) GetParameterValue() (string, bool) {
	values := strings.Join([]string{
		strconv.FormatFloat(xyz.valueX, 'f', -1, 64),
		strconv.FormatFloat(xyz.valueY, 'f', -1, 64),
		strconv.FormatFloat(xyz.valueZ, 'f', -1, 64),
	}, ", ")

	value := fmt.Sprintf("[%s]", values)

	return value, xyz.set
}

// NewFloatXYZ creates a new FloatXYZ with its value explicitly set.
func NewFloatXYZ(x, y, z float64) FloatXYZ {
	var xyz FloatXYZ
	xyz.Set(x, y, z)

	return xyz
}
