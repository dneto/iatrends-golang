package main

import (
	"fmt"
	"time"
)

func main() {
	intChan, quit := make(chan int, 10), make(chan bool)

	go func() {
		time.Sleep(1000 * time.Millisecond)
		for i := 0; i < 100; i++ {
			intChan <- i
		}
	}()

	go func() {
		fmt.Println("Waiting input")
		for {
			select { //HLaa
			case i := <-intChan:
				fmt.Println(i)
			case <-quit:
				break
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	quit <- true

}
