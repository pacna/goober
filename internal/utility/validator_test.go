package utility

import (
	"testing"
)

// TestIsStringEmpty checks to see if string is empty
func TestIsStringEmpty(t *testing.T) {
	// ARRANGE
	var emptyString string
	var expectedResult bool = true

	// ACT
	result := IsEmpty(emptyString)

	// ASSERT
	assert(expectedResult, result, t)
}

// TestIsArrayEmpty check to see if array is empty
func TestIsArrayEmpty(t *testing.T) {
	// ARRANGE
	var emptyArray [0]string
	var expectedResult bool = true

	// ACT
	result := IsEmpty(emptyArray)

	// ASSERT
	assert(expectedResult, result, t)

}

// TestIsValidHTTPURL check to url is a valid http url
func TestIsValidHTTPURL(t *testing.T) {
	// ARRANGE
	var validURL string = "https://www.google.com"
	var expectedResult bool = true

	// ACT
	result := IsHttpURL(validURL)

	// ASSERT
	assert(expectedResult, result, t)
}

func assert(expected bool, actual bool, t *testing.T) {
	if expected != actual {
		t.Errorf("Test failed: expected %t but got %t", expected, actual)
	}
}