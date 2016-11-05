package main

import (
	"fmt"
	"math"
)

func mergeSort(a1, a2 []int) []int {
	length := len(a1) + len(a2)
	a := make([]int, 0, length)

	a1Count := 0
	a2Count := 0

	for a1Count < len(a1) || a2Count < len(a2) {
		if a1Count == len(a1) {
			a = append(a, a2[a2Count])
			a2Count++
		} else if a2Count == len(a2) {
			a = append(a, a1[a1Count])
			a1Count++
		} else {
			if a1[a1Count] <= a2[a2Count] {
				a = append(a, a1[a1Count])
				a1Count++
			} else {
				a = append(a, a2[a2Count])
				a2Count++
			}
		}
	}
	
	return a
}

func heapsort(a []int) []int {
	a = buildMaxHeap(a)
	for i := len(a); i >= 2; i-- {
		temp := a[0]
		a[0] = a[i]
		a[i] = temp

	}
}

func insertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		valueToInsert := a[i]
		holePosition := i

		for holePosition > 0 && a[holePosition - 1] > valueToInsert {
			a[holePosition] = a[holePosition - 1]
			holePosition = holePosition - 1
		}

		a[holePosition] = valueToInsert
	}

	return a
}

func main() {
	array1 := []int{4, 15, 16, 50}
	array2 := []int{8, 23, 42, 108}
	array3 := []int{23, 42, 4, 16, 8, 15} 

	fmt.Println(mergeSort(array1[0:], array2[0:]))
	fmt.Println(insertionSort(array3[0:]))
}