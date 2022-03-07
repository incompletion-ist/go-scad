package values

// String represents an explicitly settable string.
type String struct {
	value string
	set   bool
}

// Set explicitly sets the given value.
func (s *String) Set(value string) {
	s.value = value
	s.set = true
}

// IsSet returns a boolean indicating if String has been explicitly set.
func (s String) IsSet() bool {
	return s.set
}

// Value returns a boolean value representing the stored value, implicit
// or explicit.
func (s String) Value() string {
	return s.value
}

// GetParameterValue returns a string value representing String, and a boolean
// indicating if it was explicitly set.
func (s String) GetParameterValue() (string, bool) {
	return s.value, s.set
}

// NewString returns a new String with the given value explicitly set.
func NewString(value string) String {
	var s String
	s.Set(value)

	return s
}
