package main

import (
	"fmt"
	"log"
	"reflect"
)

func main(){
	var r rune = '씨'
	t := '\ud55c'

	var ra []rune = []rune{}
	var ra2 [4]rune
	fmt.Println(reflect.TypeOf(ra))
	fmt.Println(reflect.TypeOf(ra).Kind())
	fmt.Println(reflect.TypeOf(ra2))
	fmt.Println(reflect.TypeOf(ra2).Kind())
	ra = append(ra, '퍽')
	fmt.Println(reflect.TypeOf(ra))
	fmt.Println(reflect.TypeOf(ra).Kind())

	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(r, t)
	log.Println(r, t)
}