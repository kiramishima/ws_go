package main

import "fmt"

func main()  {
	println("Lord of the Code")

	hrStart := []string{"10:00", "14:00", "10:00", "15:00", "10:00", "13:00"}
	fmt.Printf("%s", hrStart)
	hrEnd := []string{"17:00", "22:00", "14:00", "21:00", "17:00", "22:00"}
	fmt.Printf("%s", hrEnd)
	var hrTuples = [][]string{}

	for i, _ := range(hrStart) {
		fmt.Println(i)
		hrTuples = append(hrTuples, []string{hrStart[i], hrEnd[i]})
	}

	fmt.Printf("%s\n", hrTuples)
}