package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	var min_sum int32 = 0
	fmt.Println(arr[:4])
	fmt.Println(arr[len(arr)-4:])
	for _, n := range arr[:4] {
		min_sum += n
	}
	var max_sum int32 = 0
	for _, n := range arr[len(arr)-4:] {
		max_sum += n
	}
	fmt.Println(min_sum, max_sum)

}

func miniMaxSum2(arr []int32) {
	var sum, smallest, biggest int64
	smallest = int64(arr[0])
	for i := 0; i < 5; i++ {
		sum += int64(arr[i])
		if biggest < int64(arr[i]) {
			biggest = int64(arr[i])
		}
		if smallest > int64(arr[i]) {
			smallest = int64(arr[i])
		}
	}
	fmt.Println(sum-biggest, sum-smallest)
}

func main() {

	var arr = []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)

	miniMaxSum2(arr)

	arr = []int32{156873294, 719583602, 581240736, 605827319, 895647130}
	// 2063524951 2802298787
	miniMaxSum2(arr)
}
