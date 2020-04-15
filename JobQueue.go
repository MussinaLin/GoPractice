package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan int, 1)
	go worker(queue)

	fmt.Println(enqueue(queue,1))
	fmt.Println(enqueue(queue,2))
	fmt.Println(enqueue(queue,3))
	fmt.Println(enqueue(queue,4))
	fmt.Println(enqueue(queue,5))
	fmt.Println("finish enqueue")
	time.Sleep( 5 * time.Second)
}
func worker(q <- chan int)  {
	for i := range q {
		fmt.Println("doing job:", i)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("finish job:", i)
	}
}
func enqueue(q chan <- int, job int) bool{
	select {
	case q <- job:
			return true
	default:
		return false
	}
}