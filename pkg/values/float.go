package values

import "strconv"

// Float represents a float64 value that can be explicitly set.
type Float struct {
	value float64
	set   bool
}

// Set explicitly sets the given value.
func (f *Float) Set(value float64) {
	f.value = value
	f.set = true
}

// IsSet returns a boolean indicating if Float has been explicitly set.
func (f Float) IsSet() bool {
	return f.set
}

// Value returns the Float's stored value.
func (f Float) Value() float64 {
	return f.value
}

// GetParameterValue returns the string representation of Float, and a boolean
// indicating if it was explicitly set.
func (f Float) GetParameterValue() (string, bool) {
	return strconv.FormatFloat(f.value, 'f', -1, 64), f.set
}

// NewFloat returns a new Float with the given value explicitly set.
func NewFloat(value float64) Float {
	var f Float
	f.Set(value)

	return f
}
