package main

import (
	"fmt"
)

func main(){
	c := make(chan bool)
	go func() {
		fmt.Println("Hello !")
		<- c

	}()
	c <- true

	ch := make(chan int)

	// only read channel
	go func(ch <- chan int) {
		for v := range ch{
			fmt.Println(v)
		}

	}(ch)

	// only write channel
	go func(ch chan <- int) {
		ch <- 100
		ch <- 101
		ch <- 102
	}(ch)
	//time.Sleep(1 * time.Second)
	ch <- 104
	ch <- 105
}
