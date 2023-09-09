package iter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gmgigi96/iter"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected []int
	}{
		{"Filter even numbers", []int{1, 2, 3, 4, 5}, func(e int) bool { return e%2 == 0 }, []int{2, 4}},
		{"Filter numbers greater than 3", []int{1, 2, 3, 4, 5}, func(e int) bool { return e > 3 }, []int{4, 5}},
		{"Filter from empty slice", []int{}, func(e int) bool { return e%2 == 0 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			filtered := iter.Filter(tt.fn, it)
			result := iter.Slice(filtered)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %+v, got %+v", tt.expected, result)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) string
		expected []string
	}{
		{"Map int to string", []int{1, 2, 3}, func(e int) string { return fmt.Sprintf("%d", e) }, []string{"1", "2", "3"}},
		{"Map empty slice", []int{}, func(e int) string { return fmt.Sprintf("%d", e) }, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			mapped := iter.Map(tt.fn, it)
			result := iter.Slice(mapped)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
