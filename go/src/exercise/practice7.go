package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx := context.WithDeadline(context.Background(), time.After(time.Second*2))
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second * 3)
	go func(){
		time.Sleep(time.Second)
		cancel()
	}()
	ret := <- ctx.Done()
	fmt.Println(ret)
	fmt.Println(ctx.Err())

	time.Sleep(time.Second)
	ret = <- ctx.Done()
	fmt.Println(ret)
	fmt.Println(ctx.Err())

}
