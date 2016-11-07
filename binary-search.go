package main

import (
	"fmt"
)

func binarySearch(a []int, value int) int {
	first := 0
	last := len(a) - 1
	middle := (first + last) / 2

	for first <= last {
		if a[middle] < value {
			first = middle + 1
		} else if a[middle] > value {
			last = middle - 1
		} else {
			return middle
		}

		middle = (first + last) / 2
	}

	return -1
}

func main() {
	a := []int{1, 4, 10, 12, 19, 23, 31, 32, 40, 42, 48, 49, 50}
	fmt.Println(binarySearch(a, 42))
}