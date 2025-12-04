package context

import (
	"context"
	"fmt"
	"time"
)

func main() {

	strings := make(chan string)
	nums := make(chan int)

	go func() {
		for {
			strings <- "hello world"
			time.Sleep(time.Millisecond * 300)
		}
	}()

	go func() {
		for {
			nums <- 1
			time.Sleep(time.Millisecond * 100)
		}
	}()

	cxt, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	for {
		select {
		case msg := <-strings:
			fmt.Println(msg)
		case msg := <-nums:
			fmt.Println(msg)

		case <-cxt.Done():
			fmt.Println("error incurred: ", cxt.Err())
		}

	}

}
