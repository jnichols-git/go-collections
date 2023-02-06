package slices

import (
	"errors"
	"math/rand"
	"testing"
)

// Generate a test slice of type T, with all values in mustContain and otherwise random values from mayContain.
// Values can be in mustContain and not in mayContain, but mayContain cannot be empty.
func genTestSlice[T any](size int, mayContain []T, mustContain []T) (slice []T, err error) {
	if len(mayContain) == 0 {
		return nil, errors.New("test slices must be able to contain values")
	}
	if len(mustContain) > size {
		return nil, indexOutOfRange(len(mustContain), size)
	}
	// temp map mustContain to random indices
	mci := make(map[int]T)
	for _, mcv := range mustContain {
		for {
			idx := rand.Int() % size
			if _, exists := mci[idx]; !exists {
				mci[idx] = mcv
				break
			}
		}
	}
	// fill slice with random values from mayContain, substituting values in mci
	slice = make([]T, size)
	for i := 0; i < size; i++ {
		if mcv, ok := mci[i]; ok {
			slice[i] = mcv
		} else {
			slice[i] = mayContain[rand.Int()%len(mayContain)]
		}
	}
	return
}

var ts_simple []int = []int{0, 1, 2, 3, 4, 5}

func TestFindSimple(t *testing.T) {
	for i, v := range ts_simple {
		if Find(ts_simple, v) != i {
			t.Errorf("Expected to find %d at %d in ts_simple", v, i)
		}
	}
	for v := 6; v < 12; v++ {
		if Find(ts_simple, v) != -1 {
			t.Errorf("Expected to not find %d in ts_simple", v)
		}
	}
}

func TestFindFuncSimple(t *testing.T) {
	for i, v := range ts_simple {
		if FindFunc(ts_simple, v, func(a, b int) bool {
			return a == b
		}) != i {
			t.Errorf("Expected to find %d at %d in ts_simple", v, i)
		}
	}
	for v := 6; v < 12; v++ {
		if FindFunc(ts_simple, v, func(a, b int) bool {
			return a == b
		}) != -1 {
			t.Errorf("Expected to not find %d in ts_simple", v)
		}
	}
}

func TestContainsSimple(t *testing.T) {
	for _, v := range ts_simple {
		if !Contains(ts_simple, v) {
			t.Errorf("Expected ts_simple to contain %d", v)
		}
	}
	for v := 6; v < 12; v++ {
		if Contains(ts_simple, v) {
			t.Errorf("Expected ts_simple to not contain %d", v)
		}
	}
}

func TestContainsFuncSimple(t *testing.T) {
	for _, v := range ts_simple {
		if !ContainsFunc(ts_simple, v, func(a, b int) bool {
			return a == b
		}) {
			t.Errorf("Expected ts_simple to contain %d", v)
		}
	}
	for v := 6; v < 12; v++ {
		if ContainsFunc(ts_simple, v, func(a, b int) bool {
			return a == b
		}) {
			t.Errorf("Expected ts_simple to not contain %d", v)
		}
	}
}

func TestEqualSimple(t *testing.T) {
	if !Equal(ts_simple, ts_simple) {
		t.Errorf("Expected ts_simple to Equal itself")
	}
}

func TestEqualFuncSimple(t *testing.T) {
	if !EqualFunc(ts_simple, ts_simple, func(a int, b int) bool {
		return a == b
	}) {
		t.Errorf("Expected ts_simple to Equal itself")
	}
}

func TestEqualUnorderedSimple(t *testing.T) {
	ts_simple_shuffled := []int{1, 5, 3, 4, 2, 0}
	if !EqualUnordered(ts_simple, ts_simple_shuffled) {
		t.Errorf("expected ts_simple to EqualUnordered a shuffled version of itself")
	}
}

func TestEqualUnorderedFuncSimple(t *testing.T) {
	ts_simple_shuffled := []int{1, 5, 3, 4, 2, 0}
	if !EqualUnorderedFunc(ts_simple, ts_simple_shuffled, func(a int, b int) bool {
		return a == b
	}) {
		t.Errorf("expected ts_simple to EqualUnordered a shuffled version of itself")
	}
}

func TestFilterFuncSimple(t *testing.T) {
	ts_simple := []int{0, 1, 2, 3, 4, 5}
	ts_simple, ts_simple_out := FilterFunc(ts_simple, func(a int) bool { return a < 3 })
	if !Equal(ts_simple, []int{3, 4, 5}) {
		t.Errorf("expected ts_simple to contain exactly 3, 4, 5 after filter; got %v", ts_simple)
	}
	if !Equal(ts_simple_out, []int{0, 1, 2}) {
		t.Errorf("expected ts_simple_out to contain exactly 0, 1, 2 after filter; got %v", ts_simple)
	}
}
