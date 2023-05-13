/**
 * Eric Schmidt
 * CSS 508 Spring 2023
 * 2023-05-13
 *
 * This test file demonstrates a TDD approach to developing software. The
 * tests use the stretchr/testify package for asserts. (Go does not have a
 * fully maintained, recognized xUnit library--Testify comes closest.)
 *
 */
package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputOne(t *testing.T) {
	input := 1
	want := "one"
	got, err := IntegerToWordedString(uint(input))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, got)
}
