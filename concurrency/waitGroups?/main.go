package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(order Order) {
	fmt.Printf("Preparing order: %d...\n", order.TableNumber)
	time.Sleep(order.PrepTime)
	fmt.Printf("Order completed: %d...\n", order.TableNumber)
}

func main() {

	wg := sync.WaitGroup{}

	order := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 3 * time.Second},
		{TableNumber: 3, PrepTime: 2 * time.Second},
	}

	for _, o := range order {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processOrder(o)
		}()
	}

	wg.Wait()
	fmt.Println("Completed all products")

}
