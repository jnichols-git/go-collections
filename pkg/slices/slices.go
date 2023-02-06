package slices

// Find finds the index of a value in a slice.
//
// val can be of any comparable type T, and slice must be a slice of the same type.
func Find[T comparable](slice []T, val T) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

// Contains checks if a value is in a slice.
// See documentation for Find[comparable].
func Contains[T comparable](slice []T, val T) bool {
	return Find(slice, val) != -1
}

// Equal checks if two ordered slices are equal.
//
// Both slices must contain the same comparable type T.
func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualUnordered checks if two slices contain the same values.
//
// Both slices must contain the same comparable type T.
func EqualUnordered[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !Contains(a, b[i]) {
			return false
		}
	}
	return true
}
