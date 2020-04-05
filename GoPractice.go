package main

import (
	"fmt"
	"sync"
)

const (
	a = iota + 2
	b
	c
)

type myStruct struct {
	name string
	age  int
}

func (s myStruct) Error() string {
	return fmt.Sprintf("name:%s age:%d", s.name, s.age)
}

func switchTest(i int)  {
	switch i {
	case 0,1:
		fmt.Println("i:", i)
	}
}
func init(){
	fmt.Println("init run !")
}
func main() {

	wg := sync.WaitGroup{}
	wg.Add(10)
	fmt.Printf("ori wg:%p\n", &wg)
	for i:=0;i<10;i++{
		go func(i int, wg *sync.WaitGroup){
			fmt.Println("routine:", i)
			fmt.Printf("wg:%p\n", wg)
			fmt.Printf("&wg:%p\n", &wg)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
