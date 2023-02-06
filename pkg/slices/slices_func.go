package slices

// FindFunc finds the index of a value in a slice using a comparison function.
//
// val can be of any type U, and slice is a slice of any type T.
//
// This is particularly useful in cases where Find[comparable] is not useful, typically for data structures with non-comparable
// field types.
func FindFunc[T any, U any](slice []T, val U, compare func(T, U) bool) int {
	for i, v := range slice {
		if compare(v, val) {
			return i
		}
	}
	return -1
}

// ContainsFunc checks if a value is in a slice using a comparison function.
// See documentation for FindFunc[any].
func ContainsFunc[T any, U any](slice []T, val U, compare func(T, U) bool) bool {
	return FindFunc(slice, val, compare) != -1
}

// EqualFunc checks if two ordered slices are equal, according to some comparison function.
//
// Slice a must contain type T, and slice b must contain type U. compare must be a boolean comparison between T and U.
func EqualFunc[T any, U any](a []T, b []U, compare func(T, U) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !compare(a[i], b[i]) {
			return false
		}
	}
	return true
}

// EqualFunc checks if two slices contain the same values, according to some comparison function.
//
// Slice a must contain type T, and slice b must contain type U. compare must be a boolean comparison between T and U.
func EqualUnorderedFunc[T any, U any](a []T, b []U, compare func(T, U) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range b {
		if !ContainsFunc(a, b[i], compare) {
			return false
		}
	}
	return true
}

// FilterFunc filters a slice using a function, returning two slices with the unfiltered and filtered values, respectively.
//
// Slice a must contain type T. filter must be a function taking a value of type T and returning true if it should be filtered, or false otherwise.
func FilterFunc[T any](a []T, filter func(T) bool) ([]T, []T) {
	unfiltered := make([]T, 0)
	filtered := make([]T, 0)
	for _, v := range a {
		if filter(v) {
			filtered = append(filtered, v)
		} else {
			unfiltered = append(unfiltered, v)
		}
	}
	return unfiltered, filtered
}
