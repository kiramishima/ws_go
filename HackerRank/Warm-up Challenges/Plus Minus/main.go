package main

import "fmt"

func plusMinus(arr []int32) {
	var positive float32 = 0
	var negative float32 = 0
	var zero float32 = 0
	n := float32(len(arr))
	for _, item := range arr {
		if item < 0 {
			negative++
		} else if item > 0 {
			positive++
		} else if item == 0 {
			zero++
		}
	}

	fmt.Println(positive / n)
	fmt.Println(negative / n)
	fmt.Println(zero / n)
}

func main() {
	var arr = []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
