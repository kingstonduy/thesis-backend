package array_utils

import (
	"testing"
)

func TestIsContained(t *testing.T) {
	tests := []struct {
		name     string
		arr      interface{}
		target   interface{}
		expected bool
	}{
		{"Int element in array", []int{1, 2, 3, 4, 5}, 3, true},
		{"Int element not in array", []int{1, 2, 3, 4, 5}, 6, false},
		{"String element in array", []string{"apple", "banana", "cherry"}, "banana", true},
		{"String element not in array", []string{"apple", "banana", "cherry"}, "grape", false},
		{"Empty array", []int{}, 3, false},
		{"Nil array", nil, 3, false},
		{"Non-slice input", 123, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsContained(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("IsContained(%v, %v) = %v; expected %v", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
