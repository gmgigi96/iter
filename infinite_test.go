package iter_test

import (
	"testing"

	"github.com/gmgigi96/iter"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		step     int
		n        int
		expected []int
	}{
		{"Count from 5 with step 2", 5, 2, 3, []int{5, 7, 9}},
		{"Count from 0 with step 1", 0, 1, 3, []int{0, 1, 2}},
		{"Count from -3 with step -1", -3, -1, 3, []int{-3, -4, -5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := iter.Count(tt.start, tt.step)
			var result []int
			for i := 0; i < tt.n; i++ {
				val, _ := c.Next()
				result = append(result, val)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: Expected %v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestCycle(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{"Cycle through [1,2]", []int{1, 2}, 5, []int{1, 2, 1, 2, 1}},
		{"Cycle through [3,4,5]", []int{3, 4, 5}, 4, []int{3, 4, 5, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			cycled := iter.Cycle(it)
			var result []int
			for i := 0; i < tt.n; i++ {
				val, _ := cycled.Next()
				result = append(result, val)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: Expected %v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		times    int
		n        int
		expected []int
	}{
		{"Repeat value 5 three times", 5, 3, 3, []int{5, 5, 5}},
		{"Repeat value 7 indefinitely", 7, -1, 5, []int{7, 7, 7, 7, 7}},
		{"Repeat value 3 zero times", 3, 0, 2, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repeated := iter.Repeat(tt.value, tt.times)
			var result []int
			for i := 0; i < tt.n; i++ {
				val, ok := repeated.Next()
				if ok {
					result = append(result, val)
				}
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: Expected %v, got %v", i, tt.expected[i], v)
				}
			}
		})
	}
}
