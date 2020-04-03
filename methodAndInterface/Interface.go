package main

import "fmt"

type Point struct {
	x,y float64
}

type IPAddr [4]byte

func main() {

	emptyInterface()
	fmt.Println("")
	typeAssertion()
	fmt.Println("")

	typeSwitch(123)
	typeSwitch("mussina")
	fmt.Println("")

	fmt.Println(Point{8.7, 7.8})

	fmt.Println(IPAddr{127,0,0,1})
}

func emptyInterface(){
	var i interface{}
	describe(i)

	i = 25
	describe(i)

	i = Point{2.5, 6.4}
	describe(i)
}

func describe( v interface{}){
	fmt.Printf("%v  %T\n", v , v)
}

func typeAssertion(){
	var i interface{} = "hi"
	v,ok := i.(string)
	fmt.Println(v , ok)

	d,ok := i.(int32)
	fmt.Println(d , ok)
}

func typeSwitch(i interface{}){
	switch t := i.(type) {
	case int:
		fmt.Printf("%T value:%v \n", t,t)
	case string:
		fmt.Printf("%T value:%v is %d bytes \n", t,t, len(t))
	}
}

//One of the most ubiquitous interfaces is Stringer defined by the fmt package
func (p Point)String() string{
	return fmt.Sprintf("x:%f y:%f", p.x, p.y)
}

func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ip[0],ip[1],ip[2],ip[3])
}


