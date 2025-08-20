package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}
	var wg sync.WaitGroup

	const iterations = 1_000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			s = append(s, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(s))
}

// usaremos el comando:
//
