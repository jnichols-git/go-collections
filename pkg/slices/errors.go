package slices

import "fmt"

type errIndexOutOfRange struct {
	index int
	max   int
}

func (err *errIndexOutOfRange) Error() string {
	return fmt.Sprintf("index %d out of range in slice size %d", err.index, err.max)
}

func indexOutOfRange(index, max int) *errIndexOutOfRange {
	return &errIndexOutOfRange{
		index: index,
		max:   max,
	}
}
