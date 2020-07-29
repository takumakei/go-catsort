package catsort

// Values is a pair of the string and Cats created from the string.
type Value struct {
	String string
	Cats   Cats
}

// NewValue creates a new *Value.
func NewValue(s string) *Value {
	return &Value{String: s, Cats: NewCats(s)}
}

// CompareValue compares lhs.Cats to rhs.Cats.
func CompareValue(lhs, rhs *Value) int {
	return CompareCats(lhs.Cats, rhs.Cats)
}
