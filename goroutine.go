package main

import(
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}

func main(){
	sayHello()
	var msg = "message"
	wg.Add(1)
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}
func sayHello(){
	fmt.Println("hello")
}