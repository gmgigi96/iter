package iter_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/gmgigi96/iter"
)

func TestFromMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []iter.MapEntry[int, string]
	}{
		{"FromMap with multiple key-value pairs", map[int]string{1: "a", 2: "b", 3: "c"}, []iter.MapEntry[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}},
		{"FromMap with empty map", map[int]string{}, []iter.MapEntry[int, string]{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.FromMap(tt.input))
			sort.Slice(result, func(i, j int) bool {
				return result[i].Key < result[j].Key
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []int
	}{
		{"Keys from map with multiple key-value pairs", map[int]string{1: "a", 2: "b", 3: "c"}, []int{1, 2, 3}},
		{"Keys from empty map", map[int]string{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.Keys(tt.input))
			sort.Ints(result)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []string
	}{
		{"Values from map with multiple key-value pairs", map[int]string{1: "a", 2: "b", 3: "c"}, []string{"a", "b", "c"}},
		{"Values from empty map", map[int]string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := iter.Slice(iter.Values(tt.input))
			sort.Strings(result)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
