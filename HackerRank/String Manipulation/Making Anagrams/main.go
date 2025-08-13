package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	// Write your code here
	count := 0
	arrA := [26]int{}
	arrB := [26]int{}

	for _, v := range a {
		position := v - 97
		arrA[position]++
	}
	fmt.Println(arrA)
	for _, v := range b {
		position := v - 97
		arrB[position]++
	}
	fmt.Println(arrB)
	for i := 0; i < 26; i++ {
		difference := math.Abs(float64(arrA[i]) - float64(arrB[i]))

		count += int(difference)
	}
	return int32(count)
}

func main() {
	var a = "cde"
	var b = "abc"
	fmt.Println(makeAnagram(a, b))
	a = "fcrxzwscanmligyxyvym"
	b = "jxwtrhvujlmrpdoqbisbwhmgpmeoke"
	fmt.Println(makeAnagram(a, b))

	switch string(a) {
	case "a":
	case "e":
	case "i":
	case "o":
	case "u":
	}
}
