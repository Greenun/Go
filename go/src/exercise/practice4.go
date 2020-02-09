package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main(){
	//test()
	//fmt.Println("ohohoh..")
	runtime.GOMAXPROCS(2)
	var wait sync.WaitGroup
	wait.Add(2)

	go func(){
		defer wait.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine", i)
			time.Sleep(time.Second)
		}
	}()

	go func(msg string){
		defer wait.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Second)
		}
	}("Fucking")

	wait.Wait()
}

func test(){
	f, err := os.Open("fucking_invalid.txt")

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("fucking Error: ", r)
		}
	}()
	if err != nil{
		panic(err)
	}
	fmt.Println(f)
	defer func(){
		fmt.Println("END")
	}()
}