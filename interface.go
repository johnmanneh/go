package main

import(
	"fmt"
)

type John struct{
	Name string
}

func main(){

	var l Labeh = Number{}
		l.PrintVal([]byte("hi"))

		myInt := IntCounter(0)
		var inc Incrementer = &myInt
		for i := 0;i<10;i++{
			fmt.Println(inc.Increment())
		}
}

type Labeh interface{
	PrintVal([]byte) (int,error)
}

type Number struct{}

func (l Number) PrintVal(data []byte)(int,error){
	n,err := fmt.Println(string(data))
	return n,err
}
type Incrementer interface{
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int{
	*ic++
	return int(*ic)
}