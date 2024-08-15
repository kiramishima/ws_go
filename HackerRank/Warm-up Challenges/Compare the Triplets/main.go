package main

import "fmt"

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	var results = []int32{0, 0}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			results[0]++
		} else if a[i] < b[i] {
			results[1]++
		}
	}
	return results
}

func main() {
	var a = []int32{5, 6, 7}
	var b = []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}
