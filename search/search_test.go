package search

import (
	"testing"
)

type ArraySearchTest struct {
	haystack []int
	needle   int
	expected bool
}

var SearchTestCases = []ArraySearchTest{
	{[]int{1, 2, 3, 4, 5}, 2, true},
	{[]int{1, 2, 3, 4, 5}, 10, false},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range SearchTestCases {
		if output := LinearSearch(test.haystack, test.needle); output != test.expected {
			t.Errorf("Haystack %v containing %d expected to be %t, but %t is returned", test.haystack, test.needle, test.expected, output)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range SearchTestCases {
		if output := BinarySearch(test.haystack, test.needle); output != test.expected {
			t.Errorf("Haystack %v containing %d expected to be %t, but %t is returned", test.haystack, test.needle, test.expected, output)
		}
	}
}

type TwoCrystallBallsCase struct {
	breaks   []bool
	expected int
}

var TwoCrystallBallsCases = []TwoCrystallBallsCase{
	{[]bool{false, false, false, false, false, true, true, true}, 5},
	{[]bool{false, false, false, false, false, false, false, false}, -1},
	{[]bool{true, true, true, true, true, true, true, true}, 0},
}

func TestTwoCrystallBalls(t *testing.T) {
	for _, test := range TwoCrystallBallsCases {
		if output := TwoCrystallBalls(test.breaks); output != test.expected {
			t.Errorf("TwoCrystalBalls returned %d, but expected to be %d, for input %v", output, test.expected, test.breaks)
		}
	}
}
