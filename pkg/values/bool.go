package values

import "strconv"

// Bool represents an explicitly settable bool.
type Bool struct {
	value bool
	set   bool
}

// Set explicitly sets the given value.
func (b *Bool) Set(value bool) {
	b.value = value
	b.set = true
}

// IsSet returns a boolean indicating if Bool has been explicitly set.
func (b Bool) IsSet() bool {
	return b.set
}

// Value returns a boolean value representing the stored value, implicit
// or explicit.
func (b Bool) Value() bool {
	return b.value
}

// GetParameterValue returns a string value representing Bool, and a boolean
// indicating if it was explicitly set.
func (b Bool) GetParameterValue() (string, bool) {
	return strconv.FormatBool(b.value), b.set
}

// NewBool returns a new Bool with the given value explicitly set.
func NewBool(value bool) Bool {
	var b Bool
	b.Set(value)

	return b
}
