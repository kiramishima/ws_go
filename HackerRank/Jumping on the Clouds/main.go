package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var jumps int32
	size := len(c)
	if size == 0 {
		return jumps
	} else {
		for i := 0; i < size-1; i++ {
			if c[i] == 0 {
				i++
			}
			jumps += 1
		}
		return jumps
	}
}

func main() {
	// 4
	c := []int32{0, 0, 0, 0, 1, 0}
	var res = jumpingOnClouds(c)
	fmt.Println(res)
}
