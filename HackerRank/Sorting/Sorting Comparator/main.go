package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(aSlice)

	fmt.Println(aSlice[:5])

	t := aSlice[0:5:10]
	fmt.Println(t, len(t), cap(t))

	t = aSlice[:5:6]
	fmt.Println(t, len(t), cap(t))
}
