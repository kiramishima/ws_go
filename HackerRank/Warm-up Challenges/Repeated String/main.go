package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	searchLetter := "a"
	strlen := int64(len(s))
	multiply := n / int64(strlen)
	remanent := n % int64(strlen)
	if strlen == 0 {
		return 0
	} else if strlen == 1 && s == searchLetter {
		return n
	} else {
		return int64(strings.Count(s, searchLetter)*int(multiply) + strings.Count(s[:remanent], searchLetter))
	}
}

func main() {
	s := "bcbccacaacbbacabcabccacbccbababbbbabcccbbcbcaccababccbcbcaabbbaabbcaabbbbbbabcbcbbcaccbccaabacbbacbc"
	n := int64(649606239668)
	res := repeatedString(s, n)
	fmt.Println(res)
}
