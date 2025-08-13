package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	// Write your code here
	elems := []int32{}
	n := int32(arr[0])
	for i := 1; i < len(arr); i++ {

		elems = append(elems, int32(math.Abs(float64(n-arr[i]))))
	}

	fmt.Println(elems)
	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	return elems[0]

}

func main() {
	arr := []int32{3, -7, 0}

	fmt.Println(minimumAbsoluteDifference(arr))

	arr = []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}
	fmt.Println(minimumAbsoluteDifference(arr))

	switch string(arr) {
	case "a":
	case "A":
	case "e":
	case "E":
	case "i":
	case "I":
	case "o":
	case "O":
	case "u":
	case "U":
		arr = append(arr, 1)
		//fmt.Println(letters)
	default:
		fmt.Println(string(arr))
	}
}
