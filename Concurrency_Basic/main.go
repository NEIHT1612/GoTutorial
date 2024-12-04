package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool){
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main(){
	dones := make([]chan bool, 4)

	dones[0] = make(chan bool)
	go greet("Nice to meet u", dones[0])
	dones[1] = make(chan bool)
	go greet("How r u", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... r ... u", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope u r liking the course", dones[3])

	for _, done := range dones{
		<- done
	}
}