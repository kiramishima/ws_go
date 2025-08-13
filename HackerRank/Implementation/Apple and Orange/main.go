package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func MakeChange(amount int) map[byte]int {
	if amount <= 0 {
		return map[byte]int{}
	}
	coins := map[byte]int{'H': 50, 'Q': 25, 'D': 10, 'N': 5, 'P': 1}
	change := map[byte]int{}
	// H - 72 - 50c
	// Q - 81 - 25c
	// D - 68 - 10c
	// N - 78 - 5c
	// P  - 80 - 1c
	for coin, val := range coins {
		quot, rem := divmod(amount, val)
		change[coin] = quot
		amount = rem

		if change[coin] == 0 {
			delete(change, coin)
		}
	}
	// 72 - 1, 81 - 1, 68 - 1, 78 -1, 80 - 1
	return change
}

func main() {
	fmt.Println([]byte("P"))
}
