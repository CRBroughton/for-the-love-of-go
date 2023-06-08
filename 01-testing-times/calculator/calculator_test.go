package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
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

func TestSubtract(t *testing.T) {
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

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if !closeEnough(testCase.want, got, 0.001) {
			t.Errorf("Add(%f, %f): want %f, got %f",
				testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}
