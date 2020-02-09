package main

import (
	"fmt"
	"reflect"
)

type testst struct {
	name string
	id int16
}

type Dictionary struct {
	data map[string]interface{}
}

func (d Dictionary) Keys() []string{
	keySet := make([]string, len(d.data))
	idx := 0
	for k, _ := range d.data {
		keySet[idx] = k
		idx++
	}
	return keySet
}

func (d Dictionary) Values() []interface{} {
	valueSet := make([]interface{}, len(d.data))
	idx := 0
	for _, v := range d.data {
		valueSet[idx] = v
		idx++
	}
	return valueSet
}
func (d Dictionary) getValue(key string) interface{}{
	return d.data[key]
}

func main() {
	//t := &testst{"abc", 14}
	//testReflect(t)
	//d := map[string]interface{}{"fuck": "you", "shit": []string{"a", "b", "c"}}
	//dict := Dictionary{d}
	//k := dict.Keys()
	//fmt.Println(dict.Keys())
	//fmt.Println(dict.data[k[0]])
	//fmt.Println(dict.getValue(k[0]))
	//fmt.Println(dict.Values())
	//
	//fmt.Println("a" < "b")
	//t := 14
	//r := reflect.ValueOf(t)
	//fmt.Println(r)
	//fmt.Println(reflect.TypeOf(r))
	//var t interface{}
	//t = "abc"
	//r := reflect.ValueOf(t)
	//fmt.Println(r, reflect.TypeOf(r), r.Kind())
	//x := "10:abci15e"
	//idx := strings.IndexByte(x, ':')
	//fmt.Println(x[:idx])
	var x = "a"
	fmt.Println(x[1:])
	abc := []byte("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe")
	fmt.Println(abc)

}

func testReflect(x interface{}){
	e := reflect.ValueOf(x).Elem()
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.Indirect(e))
	fmt.Println(reflect.TypeOf(reflect.Indirect(e)))
	fmt.Println(e)
}