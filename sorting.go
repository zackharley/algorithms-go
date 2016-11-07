package main

import (
	"fmt"
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

func quicksort(a []int, left, right int) {
	pivot := a[int(left + (right - left) / 2)]
	i := left
	j  := right

	for i <= j {
		for a[i] < pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}

		if i <= j {
			temp := a[i]
			a[i] = a[j]
			a[j] = temp
			i++
			j--
		}
	}

	if left < j {
		quicksort(a, left, j)
	}
	fmt.Println(a)
	if i < right {
		quicksort(a, i, right)
	}
	fmt.Println(a)
}

func main() {
	array1 := []int{4, 15, 16, 50}
	array2 := []int{8, 23, 42, 108}
	array3 := []int{23, 42, 4, 16, 8, 15} 
	array4 := []int{2, 8, 7, 1, 3, 5, 6, 4}

	fmt.Println(mergeSort(array1[0:], array2[0:]))
	fmt.Println(insertionSort(array3))
	fmt.Println(array4, "---->");
	quicksort(array4, 0, len(array4) - 1)
	fmt.Println(array4);
}