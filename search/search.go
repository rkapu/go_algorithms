package search

import (
	"math"
)

func LinearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		m := low + (high-low)/2
		v := haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			high = m
		} else {
			low = m + 1
		}
	}

	return false
}

func TwoCrystallBalls(breaks []bool) int {
	var i int
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	for i = jumpAmount; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for ; i < len(breaks); i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}
