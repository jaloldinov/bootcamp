package test

import (
	"testing"

	"github.com/test-go/testify/assert"
)

// unit test
func TestSum(t *testing.T) {
	a, b := 4, 9

	res := sum(a, b)

	assert.Equal(t, 13, res)

}

func TestMultiply(t *testing.T) {
	a, b := 4, 5

	res := multiply(a, b)

	assert.Equal(t, 20, res)
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

type argsAndExpected struct {
	a, b, expected int
}

// table driven test
func TestSumTD(t *testing.T) {
	cases := []argsAndExpected{
		{a: 6, b: 3, expected: 9},
		{a: 0, b: 8, expected: 8},
		{a: -4, b: -7, expected: -11},
		{a: -4, b: 0, expected: -4},
	}
	for _, c := range cases {
		res := sum(c.a, c.b)
		assert.Equal(t, c.expected, res)
	}
}

func TestMultiplyTD(t *testing.T) {
	cases := []argsAndExpected{
		{a: 5, b: 10, expected: 50},
		{a: 2, b: 10, expected: 20},
		{a: 7, b: 8, expected: 56},
		{a: 8, b: 11, expected: 88},
	}

	for _, c := range cases {
		res := multiply(c.a, c.b)

		assert.Equal(t, c.expected, res)
	}
}
