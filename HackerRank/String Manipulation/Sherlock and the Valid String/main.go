package main

import (
	"fmt"
	"sort"
	"strings"
)

func isValid(s string) string {
	// Para no repetir el YES / NO
	const YES = "YES"
	const NO = "NO"
	// validamos que no venga con espacios vacio
	if len(strings.TrimSpace(s)) == 0 {
		return NO
	}
	var arr2 = make([]int, 0)
	arr := [26]int{}
	for _, l := range s {
		arr[int(l)-int('a')]++
	}

	for _, item := range arr {
		if item != 0 {
			arr2 = append(arr2, item)
		}
	}

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	fmt.Println(arr2)

	elem0 := arr2[0]
	elem1 := arr2[len(arr2)-1]

	result := NO
	if elem0 == elem1 {
		result = YES
	} else {
		if (elem1-elem0 == 1) && (elem1 > arr2[len(arr2)-2]) {
			result = YES
		} else if (elem0 == 1) && arr2[1] == elem1 {
			result = YES
		}
	}
	return result
}

func main() {
	var sample = "aabbcd"

	fmt.Println(isValid(sample))

	sample = "aabbccddeefghi"

	fmt.Println(isValid(sample))
}
