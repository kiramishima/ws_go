package main

import (
	"fmt"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {

	var max int32
	for i := 0; i < 4; i++ {
		// 3x3
		if i > 3 {
			break
		}
		for j := 0; j < 4; j++ {
			var total int32
			total = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			// fmt.Println(total)
			if total > max || (i == 0 && j == 0) {
				max = total
			}
		}

	}
	return max
}

func main() {
	var arr = [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	var arr2 = [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	var arr3 = [][]int32{
		{-1, 1, -1, 0, 0, 0},
		{0, -1, 0, 0, 0, 0},
		{-1, -1, -1, 0, 0, 0},
		{0, -9, 2, -4, -4, 0},
		{-7, 0, 0, -2, 0, 0},
		{0, 0, -1, -2, -4, 0},
	}
	res := hourglassSum(arr)
	fmt.Println("Result", res)

	res = hourglassSum(arr2)
	fmt.Println("Result2", res)

	res = hourglassSum(arr3)
	fmt.Println("Result3", res)
}
