package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
	// Write your code here
	var ss1 = strings.Split(s1, "")
	var ss2 = strings.Split(s2, "")
	var subs = make(map[string]int)
	for _, word := range ss1 {
		subs[word]++
	}

	for _, word := range ss2 {
		if subs[word] > 0 {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	var s1 = "hello"
	var s2 = "world"
	fmt.Println(twoStrings(s1, s2))
	s1 = "hi"
	s2 = "world"
	fmt.Println(twoStrings(s1, s2))
}
