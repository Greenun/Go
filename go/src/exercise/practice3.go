package main

import (
	"fmt"
	"math"
)

type sample struct {
	name string
	id int
	data map[string]string
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func main() {
	//testSample := sample{"abc", 10, make(map[string]string)}
	//var test2 sample = sample{"abc2", 10, make(map[string]string)}
	//fmt.Println(test2)
	//fmt.Println(testSample)
	//testSample.makeName("fuck")
	//testSample.dummyName("shit")
	//fmt.Println(testSample)
	//everything(testSample)
	//everything("test")
	r1 := Rectangle{10, 30}
	c1 := Circle{5}
	area(r1)
	//area(c1)
	fmt.Println(c1)

	var c2 interface{} = Circle{10}
	fmt.Println(c2.(Rectangle))

}

func (s *sample) makeName(name string) string{
	s.name = name
	return s.name
}

func (s sample) dummyName(name string) string{
	s.name = name
	return s.name
}

func everything(value interface{}){
	fmt.Println(value)
}

func (c Circle) area() float64{
	return c.radius*c.radius*math.Pi
}
func (c Circle) perimeter() float64 {
	return c.radius * 2 * math.Pi

}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return (r.height + r.width) * 2
}

func area(shape Shape) interface {} {
	println(shape.area())
	println(uint(shape.area()))
	return 0
}