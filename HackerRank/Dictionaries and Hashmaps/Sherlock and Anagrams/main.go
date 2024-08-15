package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	var letters = strings.Split(s, "")
	fmt.Println(letters)
	// var m = make(map[string]int)
	for i := 0; i < len(letters); i++ {
		for j := i + 1; j <= len(letters); j++ {
			subChar := s[i:j]
			fmt.Println(subChar)
		}
	}

	var count int32
	return count
}

func main() {
	var s = "abba"
	fmt.Println(sherlockAndAnagrams(s)) // 4
	s = "abcd"
	fmt.Println(sherlockAndAnagrams(s)) // 0
}
