package main

import (
	"fmt"
)

type add func(first int, second int) (result int)

type student struct {
	name string
	id int
}

func main(){
	//var msg string = "fuck you"
	//printer(&msg)
	//multiPrinter("what", "the", "fuck")
	//println(msg)

	//testSumFunction := func(n ...int) (result int){
	//	result = 0
	//	for _, value := range n {
	//		result += value
	//	}
	//	return
	//}
	//
	//println(testSumFunction(1,2,3,4,5))
	//
	//sum := func(one int, two int) (result int) {
	//	result = one + two
	//	return
	//}
	//t := executor(sum, 10, 30)
	//println(t)
	//t2 := executorTwo(sum, 50, 100)
	//println(t2)

	//g := generator()
	//for i := g(); i < 10; i = g() {
	//	println(i)
	//}

	//var testArr = [...]int{1,2,3,}
	//testArr[0] = 10
	//testArr[1] = 50
	//
	//println(testArr[0])
	//for i, v := range testArr {
	//	println(i, v)
	//}
	//
	//var slice []int
	//slice = []int{1,2,3}
	//for i, g := range slice {
	//	println(i, g)
	//}
	//
	//var slice2 []int = make([]int, 5, 7)
	//println(len(slice2), cap(slice2))
	////for i, v := range slice2 {
	////	println(i, v)
	////}
	//slice2 = append(slice2, 10)
	//slice2 = append(slice2, 150)
	//slice2 = append(slice2, 1000)
	//for i, v := range slice2 {
	//	println(i, v)
	//}
	//slice2 = append(slice2, slice...)
	//print(slice2)

	var hashmap = make(map[int]string)
	hashmap[1] = "fuck"
	println(hashmap[1])
	println(hashmap[5100])
	val, exists := hashmap[100]
	println(val, exists)

	p := student{name: "name", id: 100}
	var pp *student = &p

	fmt.Println(hashmap)
	fmt.Println(p)
	fmt.Println(*pp)
	println(pp.name)
}

func printer(msg *string) int{
	println(*msg)
	*msg = "Shit.."
	return 0
}

func multiPrinter(msg ...string) (something1 int, something2 int){
	println(msg)
	for idx, s := range msg {
		println(idx, s)
	}
	return 0, 1
}

func executor(f func (x int, y int) int, a int, b int) int{
	println("Executed")
	return f(a, b)
}

func executorTwo(f add, a int, b int) int{
	println("Executed Second")
	return f(a, b)
}

func generator() func() int {
	// closure
	ret := 0
	return func() int {
		ret++
		return ret
	}
}