package channels

import (
	"fmt"
	"time"
)

func sendMessage(msgChan chan<- string, num int) {
	fmt.Println("message is in process!! %d", num)

	time.Sleep(time.Second * time.Duration(num))

	mess := fmt.Sprintf("message completed: %d", num)

	msgChan <- mess
}

func recieveMessage(msgChan <-chan string) {
	fmt.Println("waiting for messages!!")

	for msg := range msgChan {
		fmt.Println("Recived: ", msg)
	}

}

func main() {
	msgChan := make(chan string)

	go sendMessage(msgChan, 2)
	go sendMessage(msgChan, 3)

	recieveMessage(msgChan)

	close(msgChan)
}
