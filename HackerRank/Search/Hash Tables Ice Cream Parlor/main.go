package main

import "fmt"

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
	var comb = make(map[int]int)
	for idx, item := range cost {
		comb[int(item)] = idx
	}
	// fmt.Println(comb)
	for idx, item := range cost {
		var diff = money - item
		// fmt.Println(diff)
		if _, ok := comb[int(diff)]; ok && comb[int(diff)] > idx {
			fmt.Println(idx+1, comb[int(diff)]+1)
			return
		}
	}
}

func main() {
	var money = int32(4)
	var cost = []int32{1, 4, 5, 3, 2}

	whatFlavors(cost, money) // [1 , 4]
	cost = []int32{2, 2, 4, 3}

	whatFlavors(cost, money) // [1 , 2]
}
