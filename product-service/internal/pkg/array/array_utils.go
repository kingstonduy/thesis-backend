package array_utils

import (
	"fmt"
	"reflect"
)

func IsContained(arr interface{}, target interface{}) bool {
	// Ensure the input is a slice
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		fmt.Println("Provided input is not a slice")
		return false
	}

	// Iterate through the slice
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == target {
			fmt.Println("Element is in the array")
			return true
		}
	}

	return false
}
