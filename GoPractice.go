package main

import (
	"fmt"
	"time"
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

	for i:=0;i<10;i++{
		go func(i int){
			fmt.Println("routine:", i)
		}(i)
		fmt.Println(i)
	}

	time.Sleep(1 * time.Second)
}
