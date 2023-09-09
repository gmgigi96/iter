package iter_test

import (
	"reflect"
	"testing"

	"github.com/gmgigi96/iter"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		stop     int
		expected []int
	}{
		{"Range up to 3", 3, []int{0, 1, 2}},
		{"Range up to 0", 0, []int{}},
		{"Range up to -3", -3, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.Range(tt.stop))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRange2(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		stop     int
		expected []int
	}{
		{"Range from 1 to 4", 1, 4, []int{1, 2, 3}},
		{"Range from 4 to 1", 4, 1, []int{}},
		{"Range from -2 to 2", -2, 2, []int{-2, -1, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.Range2(tt.start, tt.stop))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRange3(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		stop     int
		step     int
		expected []int
	}{
		{"Range from 1 to 10 with step 3", 1, 10, 3, []int{1, 4, 7}},
		{"Range from 10 to 1 with negative step -3", 10, 1, -3, []int{10, 7, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.Range3(tt.start, tt.stop, tt.step))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []iter.Enum[string]
	}{
		{"Enumerate ['a', 'b', 'c']", []string{"a", "b", "c"}, []iter.Enum[string]{{0, "a"}, {1, "b"}, {2, "c"}}},
		{"Enumerate empty slice", []string{}, []iter.Enum[string]{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			result := iter.Slice(iter.Enumerate(it))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestZip(t *testing.T) {
	tests := []struct {
		name     string
		input1   []int
		input2   []string
		expected []iter.Pair[int, string]
	}{
		{"Zip [1, 2, 3] and ['a', 'b', 'c']", []int{1, 2, 3}, []string{"a", "b", "c"}, []iter.Pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}},
		{"Zip mismatched lengths", []int{1, 2}, []string{"a", "b", "c"}, []iter.Pair[int, string]{{1, "a"}, {2, "b"}}},
		{"Zip with empty slice", []int{}, []string{"a", "b", "c"}, []iter.Pair[int, string]{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it1 := iter.FromSlice(tt.input1)
			it2 := iter.FromSlice(tt.input2)
			result := iter.Slice(iter.Zip(it1, it2))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestChan(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Chan with multiple elements", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Chan with empty slice", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			ch := iter.Chan(it)
			result := []int{}
			for val := range ch {
				result = append(result, val)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestChanBuff(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		buffSize int
		expected []int
	}{
		{"ChanBuff with multiple elements and buffer size 3", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 4, 5}},
		{"ChanBuff with empty slice and buffer size 3", []int{}, 3, []int{}},
		{"ChanBuff with multiple elements and buffer size 0", []int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := iter.FromSlice(tt.input)
			ch := iter.ChanBuff(it, tt.buffSize)
			result := []int{}
			for val := range ch {
				result = append(result, val)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
