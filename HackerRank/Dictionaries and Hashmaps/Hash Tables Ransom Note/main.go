package main

import "fmt"

/*
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) {
	// Write your code here
	var magWords = make(map[string]int)
	var notesCount = 0

	for _, word := range magazine {
		magWords[word]++
	}

	for _, word := range note {
		if _, ok := magWords[word]; ok {
			if magWords[word] > 0 {
				notesCount++
				magWords[word]--
			}
		}
	}

	if len(note) != notesCount {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}

func main() {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	checkMagazine(magazine, note) // Yes

	magazine = []string{"two", "times", "three", "is", "not", "four"}
	note = []string{"two", "times", "two", "is", "four"}
	checkMagazine(magazine, note) // No
}
