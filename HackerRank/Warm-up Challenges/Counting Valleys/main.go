package main

import "fmt"

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var m int32 = 0
	l := 0
	for _, v := range path {
		if v == 'D' {
			l--
		} else {
			if l == -1 {
				m++
			}
			l++
		}
	}

	return m
}

func main() {
	var steps int32 = 8
	var path = "UDDDUDUU"

	fmt.Println(countingValleys(steps, path))
}
