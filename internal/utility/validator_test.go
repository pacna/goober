package utility

import (
	"testing"
)

// TestIsEmpty checks to see if string is empty
func TestIsEmpty(t *testing.T) {
	// ARRANGE
	var emptyString string = ""
	var expectedResult bool = true

	// ACT
	result := IsEmpty(emptyString)

	// ASSERT
	assert(expectedResult, result, t)
}

func assert(expected bool, actual bool, t *testing.T) {
	if expected != actual {
		t.Errorf("Test failed: expected %t but got %t", expected, actual)
	}
}