// O(n^2)

package main

import (
	"fmt"
)

func bubbleSort(a []int) {
	for i := 0; i < len(a) - 1; i++ {
		for j := 0; j < len(a) - i - 1; j++ {
			if a[j] > a[j + 1] {
				temp := a[j]
				a[j] = a[j + 1]
				a[j + 1] = temp
			}
		}
	}
}

func main() {
	a := []int{5, 1, 3, 2, 7, 42, 10, 3}
	fmt.Println(a, "-->")
	bubbleSort(a)
	fmt.Println(a)
}