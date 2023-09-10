package iter_test

import (
	"reflect"
	"testing"

	"github.com/gmgigi96/iter"
)

func TestAccumulate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int, int) int
		init     int
		expected []int
	}{
		{"Accumulate with sum", []int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 0, []int{1, 3, 6, 10}},
		{"Accumulate with multiplication", []int{1, 2, 3, 4}, func(a, b int) int { return a * b }, 1, []int{1, 2, 6, 24}},
		{"Accumulate empty slice", []int{}, func(a, b int) int { return a + b }, 1, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Slice(iter.Accumulate(it, tt.fn, tt.init))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int, int) int
		init     int
		expected int
	}{
		{"Reduce with sum", []int{1, 2, 3, 4}, func(a, b int) int { return a + b }, 0, 10},
		{"Reduce with multiplication", []int{1, 2, 3, 4}, func(a, b int) int { return a * b }, 1, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Reduce(it, tt.fn, tt.init)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestChain(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []int
		expected []int
	}{
		{"Chain two slices", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"Chain with an empty slice", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it1 := iter.FromSlice(tt.input1)
			it2 := iter.FromSlice(tt.input2)
			result := iter.Slice(iter.Chain(it1, it2))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDropWhile(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		pred     func(int) bool
		expected []int
	}{
		{"Drop while less than 3", []int{1, 2, 3, 4, 5}, func(e int) bool { return e < 3 }, []int{3, 4, 5}},
		{"Drop while even", []int{2, 4, 6, 7, 8}, func(e int) bool { return e%2 == 0 }, []int{7, 8}},
		{"Drop from empty slice", []int{}, func(e int) bool { return e < 3 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Slice(iter.DropWhile(it, tt.pred))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFilterFalse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		pred     func(int) bool
		expected []int
	}{
		{"FilterFalse even numbers", []int{1, 2, 3, 4, 5}, func(e int) bool { return e%2 == 0 }, []int{1, 3, 5}},
		{"FilterFalse numbers greater than 3", []int{1, 2, 3, 4, 5}, func(e int) bool { return e > 3 }, []int{1, 2, 3}},
		{"FilterFalse from empty slice", []int{}, func(e int) bool { return e%2 == 0 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Slice(iter.FilterFalse(it, tt.pred))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		pred     func(int) bool
		expected []int
	}{
		{"TakeWhile less than 3", []int{1, 2, 3, 4, 5}, func(e int) bool { return e < 3 }, []int{1, 2}},
		{"TakeWhile even numbers", []int{2, 4, 6, 7, 8}, func(e int) bool { return e%2 == 0 }, []int{2, 4, 6}},
		{"TakeWhile from empty slice", []int{}, func(e int) bool { return e < 3 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Slice(iter.TakeWhile(it, tt.pred))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{"ForEach through [1, 2, 3, 4, 5]", []int{1, 2, 3, 4, 5}},
		{"ForEach from empty slice", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := []int{}
			iter.ForEach(it, func(n int) {
				result = append(result, n)
			})
			if !reflect.DeepEqual(result, tt.input) {
				t.Errorf("Expected %v, got %v", tt.input, result)
			}
		})
	}
}
