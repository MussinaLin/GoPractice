package main

import "fmt"

func main(){
	c := make(chan bool)
	go func() {
		fmt.Println("Hello !")
		<- c

	}()

	c <- true
}
