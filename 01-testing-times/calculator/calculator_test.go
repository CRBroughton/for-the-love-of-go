package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)

		if testCase.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtrat(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 4, want: -3},
		{a: 5, b: 2, want: 3},
	}

	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)

		if testCase.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 4, want: 4},
		{a: 5, b: 2, want: 10},
	}

	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)

		if testCase.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				testCase.a, testCase.b, testCase.want, got)
		}
	}
}
