package main

import (
	"fmt"
	"strconv"
)

func sockMerchant(n int32, ar []int32) int32 {
	socks := make(map[string]int32)
	for _, sock := range ar {
		socks[strconv.FormatInt(int64(sock), 10)] += 1
	}

	var result int32
	for _, v := range socks {
		result += v / 2
	}

	return result
}

func main() {
	var n = int32(9)
	var ar = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(n, ar))
}
