package main

import (
	"fmt"
	"time"
)

func main(){
	//c := make(chan int)
	//c <- 1
	//fmt.Println(<- c)

	//c2 := make(chan int, 4)
	//c2 <- 1
	//fmt.Println(<- c2)

	//var c chan string = make(chan string, 3)
	//
	//defer close(c)
	//
	//sendChannel(c, "sg")
	//fmt.Println(receiveChannel(c))
	//
	//for v := range c {
	//	fmt.Println(v)
	//}

	//done := make(chan bool)
	//done2 := make(chan bool)
	//
	//go func(msg string){
	//	for i := 0; i < 3; i++ {
	//		fmt.Println(msg)
	//		time.Sleep(time.Second)
	//	}
	//	done <- true
	//}("Hello1")
	//
	//go func(msg string){
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(msg)
	//		time.Sleep(time.Second)
	//	}
	//	done2 <- true
	//}("Hello2")
	//
	//EXIT:
	//	for {
	//		select {
	//		case <- done:
	//			fmt.Println("done 1")
	//		case <- done2:
	//			fmt.Println("done 2")
	//			break EXIT
	//		}
	//	}

	requestChan := make(chan chan string)
	go goroutineA(requestChan)
	go goroutineB(requestChan)
	time.Sleep(time.Second)

	done := make(chan int)
	d2 := make(chan bool)
	go func(){
		done <- 1
		d2 <- true
	}()

	go func(){
		done <- 2
		d2 <- true
	}()

	go func(){
		done <- 3
		d2 <- true
	}()

	go func(){
		for i := range done {
			fmt.Println(i)
		}
	}()
	for i := 0; i < 3; i++ {
		<- done
	}

}

func sendChannel(channel chan <- string, msg string){
	channel <- msg
	return
}

func receiveChannel(channel <- chan string) (string, bool){
	value, err := <- channel
	return value, err
}

func goroutineA(requestChan chan chan string){
	responseChan := make(chan string)
	requestChan <- responseChan
	response := <- responseChan
	fmt.Println("Async goroutineA",response)
}

func goroutineB(requestChan chan chan string){
	responseChan := <- requestChan
	responseChan <- "wessup"
}