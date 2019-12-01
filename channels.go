package main

import(
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
//	sender(2,2,3)
//	senderReceiver()
//	sendBothWays()
	routineBuffer()

	
}

func sender(msg ...int)  {
	ch := make(chan []int)
	wg.Add(2)
	go func(){
		x := <- ch
		fmt.Println(x)
		wg.Done()
	}()
	go func(){
		ch <- msg
		wg.Done()
	}()
	wg.Wait()

}

func senderONLYReceiver(){
	ch := make(chan int)
	wg.Add(2)
	//this channel is a receiver
	go func(ch <- chan int){
		i:= <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	//this channel is a sender
	go func(ch chan <- int){
		ch <- 43
		wg.Done()
	}(ch)
	wg.Wait()
}

//this goroutine sends in both ways.. in the same channel
func sendBothWays(){
	ch := make(chan int)
	wg.Add(2)
	//receiving and sending to channel
	go func(){
		val := <- ch
		fmt.Println(val)
		ch <- 44
		wg.Done()
	}()
	//sending to channel and receiving from channel
	go func(){
		ch <- 50
		fmt.Println(<- ch)
		wg.Done()
	}()
	wg.Wait()
}

//buffer allows us to send multiple values to channel 
func routineBuffer(){
	channel := make(chan int,4)//make channel and add buffer
	wg.Add(2)
	//receive from channel
	go func(channel <- chan int){
		//loop throug the channel
		for i := range channel{
			fmt.Println(i)
		}
		wg.Done()
	}(channel)
	//sending to channel
	go func(channel chan <- int){
		channel <- 100
		channel <- 200
		close(channel) //after sending close the channel 
		wg.Done()
	}(channel)
	wg.Wait()
}