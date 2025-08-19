package main

import "fmt"

// var ch = make(chan string)     // unbuffered channels
// var ch2 = make(chan string, 1) // buffered channel
// unbuffered channels
//func sender() {
//	ch <- "message"
//}

//func receive() {
//	msg := <-ch
//	fmt.Println(msg)
//}

func main() {
	// syncronization elided for clarity
	//go sender()
	//go receive()
	ch := make(chan string)

	ch <- "message"

	fmt.Println(<-ch)
}
