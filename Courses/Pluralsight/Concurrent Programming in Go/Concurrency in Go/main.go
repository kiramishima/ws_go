package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan order)
	var validOrdrCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrdrCh, invalidOrderCh)
	wg.Add(1)
	go func(validOrdrCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrdrCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrdrCh, invalidOrderCh)

	wg.Wait()
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	// order := <-in
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
	close(out)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
