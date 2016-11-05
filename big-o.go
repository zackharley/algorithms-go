package main

import (
	"fmt"
)

// VI.1 - O(b) linear
func product(a, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		sum += a
	}
	return sum
}

// VI.2 - O(b) linear
func power(a, b int) int {
	if b < 0 {
		return 0
	} else if b == 0 {
		return 1
	} else {
		return a * power(a, b - 1)
	}
}

// VI.3 - O(1) - constant time
func mod(a, b int) int {
	if b <= 0 {
		return -1
	}
	div := a / b
	return a - (div * b)
}

// VI.4 - O(a/b)
func div(a, b int) int {
	count := 0
	sum := b

	for sum <= a {
		sum += b
		count++
	}

	return count
}

// VI.5 - O(log n)
func sqrt(n int) int {
	return sqrtHelper(n, 1, n);
}

func sqrtHelper(n, min, max int) int {
	if max < min {
		return -1;
	}

	guess := (min + max) / 2
	if guess * guess == n {
		return guess
	} else if guess * guess < n {
		return sqrtHelper(n, guess + 1, max)
	} else {
		return sqrtHelper(n, min, guess - 1)
	}
}

// VI.6 - O(sqrt(n))
func sqrtAlso(n int) int {
	for guess := 1; guess * guess <= n; guess++ {
		if guess * guess == n {
			return guess
		}
	}
	return -1
}

// VI.7 - O(n)
// If a binary search tree is not balanced, how long might it take (worst case) to find an element in it?
// (Visit each node once)

// VI.8 - O(n)
// You are looking for a specific value in a binary tree, but the tree is not a binary search tree. What is
// the time complexity of this? (Visit each node once)

// VI.9 - O(n^2)
func copyArray(array []int) []int {
	cp := make([]int, 0)
	fmt.Println(len(array))
	for i := 0; i < len(array); i++ { 
		cp = appendToNew(cp, array[i])
	}
	return cp
}

func appendToNew(array []int, value int) []int {
	bigger := make([]int, len(array) + 1)
	for i := 0; i < len(array); i++ {
		bigger[i] = array[i]
	}

	bigger[len(bigger) - 1] = value

	return bigger
}

// VI.10 - O(log n)
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

var numChars int = 26

func main() {
	fmt.Println("3 * 4 =", product(3, 4))
	fmt.Println("2^4 =", power(2, 4))
	fmt.Println("5 %% 3 =", mod(5, 3))
	fmt.Println("4 / 2 =", div(4, 2))
	fmt.Println("sqrt(16) =", sqrt(16))

	array := make([]int, 6)

	for i := 0; i < len(array); i++ {
		array[i] = len(array) - i
	}

	fmt.Println("I'm copying an array", array, "->", copyArray(array))
	fmt.Println("The digits in 365 sum to", sumDigits(365))
	fmt.Println("Sorted strings:")
	printSortedStrings(6, "")
}