package main

import "fmt"

/*
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func countSwaps(a []int32) {
	n := len(a)
	var count = 0
	for i := 0; i < n; i++ {

		for j := 0; j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				count++
			}
		}
	}
	fmt.Printf(`Array is sorted in %d swaps.
	First Element: %d
	Last Element: %d`, count, a[0], a[n-1])
}

func main() {
	a := []int32{6, 4, 1}
	countSwaps(a)
}
