package xjoin

import (
	"log"
	"testing"
)

func equals(a [][]interface{}, b [][]interface{}) bool {
	// Return false if the lengths of the arrays mismatch
	if len(a) != len(b) {
		return false
	}

	// Return false if the lengths of the slices in an array
	// mismatches
	for i, slice := range a {
		if len(slice) != len(b[i]) {
			return false
		}
	}

	// Test if the element values match
	for i, slice := range a {
		for j, elem := range slice {
			if elem != b[i][j] {
				return false
			}
		}
	}

	return true
}

// test that no slice returns an empty array
func TestEmptySlice(t *testing.T) {
	s1 := []interface{}{}
	expected := append([][]interface{}{}, []interface{}{})
	result := Join(s1)

	if equals(expected, result) == false {
		t.Fatalf("\nError for empty slice:\n\tExpected: %v\n\tResult: %v", expected, result)
	}
}

// test that a single slice returns that slice
func TestOneSlice(t *testing.T) {
	s1 := []interface{}{"a", "b"}
	expected := append([][]interface{}{}, []interface{}{"a", "b"})
	result := Join(s1)

	if equals(expected, result) == false {
		t.Fatalf("\nError for 1 slice:\n\tExpected: %v\n\tResult: %v", expected, result)
	}
}

func TestTwoSlices(t *testing.T) {
	s1 := []interface{}{"a", "b"}
	s2 := []interface{}{1, 2, 3}
	expected := append([][]interface{}{}, []interface{}{"a", 1}, []interface{}{"a", 2}, []interface{}{"a", 3}, []interface{}{"b", 1}, []interface{}{"b", 2}, []interface{}{"b", 3})
	result := Join(s1, s2)
	log.Println(result)

	if equals(expected, result) == false {
		t.Fatalf("\nError for empty slice:\n\tExpected: %v\n\tResult: %v", expected, result)
	}
}
