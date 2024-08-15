package main

import "fmt"

func minimumBribes(q []int32) {
	// Write your code here
	var n = int32(0)
	var count = 0
	fmt.Println(q)
	var idx = 0
	// buscamos el indice donde esta el numero problematico
	for i := 0; i < len(q)-1; i++ {
		if q[i] > q[i+1] {
			idx = i
			break
		}
	}
	n = q[idx]
	// Comenzamos desde el indice mal acomodado
	for i := idx; i < len(q)-1; i++ {
		if n > q[i] {
			count++
		}
	}
	fmt.Println(count)

}

func main() {
	var q = []int32{1, 2, 3, 5, 4, 6, 7, 8}
	minimumBribes(q) // 1

	q = []int32{4, 1, 2, 3}
	minimumBribes(q) // 3 -> Too chaotic
}
