package main

import (
	"fmt"
	"sync"
	"time"
)

type taskType struct {
	id        int
	timeTakes int
}

func worker(task taskType, msgchan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("wok started: %d\n", task.id)

	time.Sleep(time.Second * time.Duration(task.timeTakes))

	msg := fmt.Sprintf("work done for task: %d\n", task.id)

	msgchan <- msg

}

func main() {
	msgchan := make(chan string)
	wg := sync.WaitGroup{}

	tasks := []taskType{
		{id: 1, timeTakes: 2},
		{id: 2, timeTakes: 3},
	}

	for _, work := range tasks {
		wg.Add(1)
		go worker(work, msgchan, &wg)
	}

	go func() {
		wg.Wait()
		close(msgchan)
	}()

	for ms := range msgchan {
		fmt.Println(ms)
	}

}
