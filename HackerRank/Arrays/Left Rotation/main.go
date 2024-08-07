package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	// Write your code here
	if len(a) == 0 {
		return a
	}

	res := make([]int32, 0)
	res = append(res, a[d:]...)
	res = append(res, a[:d]...)
	return res
}

func main() {
	var arr = []int32{1, 2, 3, 4, 5}
	var res = rotLeft(arr, 4)
	fmt.Println(res) // 5 1 2 3 4

}
