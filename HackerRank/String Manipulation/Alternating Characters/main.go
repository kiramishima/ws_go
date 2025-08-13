package main

import "fmt"

func alternatingCharacters(s string) int32 {
	// Write your code here
	var result int32 = 0
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == s[idx-1] {
			result++
		}
	}
	return result
}

func main() {
	var input = []string{"AAAA", "BBBBB", "ABABABAB", "BABABA", "AAABBB"}

	for _, s := range input {
		fmt.Println(alternatingCharacters(s))
	}
}
