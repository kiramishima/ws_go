package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	template := fmt.Sprintf("%s%d%s\n", "%", n, "s")
	// fmt.Println(template)
	for i := int32(1); i <= n; i++ {
		fmt.Printf(template, strings.Repeat("#", int(i)))
	}
}

func main() {
	var n = int32(6)
	//     #
	//    ##
	//   ###
	//  ####
	// #####
	//######
	staircase(n)
}
