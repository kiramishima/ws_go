package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	t, _ := time.Parse("03:04:05PM", s)
	return t.Format("15:04:05")
}

func main() {
	var s = "07:05:45PM"
	fmt.Println(timeConversion(s)) // 19:05:45

	s = "12:01:00AM"
	fmt.Println(timeConversion(s)) // 00:01:00
}
