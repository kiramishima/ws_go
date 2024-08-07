package main

import "fmt"

func minimumBribes(q []int32) {
	// Write your code here
	var n = q[0]
	var count = 0
	fmt.Println(q)
	for i := 1; i < len(q); i++ {
		if n < q[i] {
			n = q[i]
			continue
		} else {
			var old = int32(n)
			q[i-1] = old
			q[i] = n
			// fmt.Println(n, old, q[i-1])
			if count > 2 {
				fmt.Println("Too chaotic")
				return
			} else {
				fmt.Println(count)
			}
			count++
		}
	}
	fmt.Println(q)

}

func main() {
	var q = []int32{1, 2, 3, 5, 4, 6, 7, 8}
	minimumBribes(q) // 1

	q = []int32{4, 1, 2, 3}
	minimumBribes(q) // 3 -> Too chaotic
}
