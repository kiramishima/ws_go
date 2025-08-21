package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for range time.Tick(500 * time.Millisecond) {
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			fmt.Println("tick!")
		}
		wg.Done()
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}
