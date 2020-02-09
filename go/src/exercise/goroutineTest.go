package main

import (
	"fmt"
	"runtime"
	"sync"
)
func main(){
	//stringList := []string{"1234", "5678", "9101112", "13141516"}
	//groups := sync.WaitGroup{}
	//groups.Add(1)
	//groups.Add(1)
	//groups.Add(1)
	//groups.Add(1)
	//for _, s := range stringList {
	//	fmt.Println("outside goroutine:", s)
	//	go func() {
	//		fmt.Println("inside goroutine:", s)
	//		groups.Done()
	//	}()
	//}
	//groups.Wait()
	group := sync.WaitGroup{}
	testSlice := make([]int, 100000)
	for i := 0; i < len(testSlice); i++ {
		testSlice[i] = i + 130
		group.Add(1)
	}
	runtime.GOMAXPROCS(4)
	for _, i := range testSlice {
		go func() {
			fmt.Print(i, " ")
			group.Done()
		}()
	}
	group.Wait()

}
