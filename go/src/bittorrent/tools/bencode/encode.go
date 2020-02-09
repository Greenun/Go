package bencode

import (
	"fmt"
	"reflect"
	"sort"
)

func Encode(target interface{}) string{
	// keys must be string
	rv := reflect.ValueOf(target)
	str := ""
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = fmt.Sprintf("i%de", rv.Int())
	case reflect.String:
		str = fmt.Sprintf("%d:%s", len(rv.String()), rv.String())
	case reflect.Bool:
		var v int
		if rv.Bool() {
			v = 1
		} else { v = 0 }
		str = fmt.Sprintf("i%de", v)
	case reflect.Map:
		str += "d"
		d, ok := target.(map[string]interface{})
		if !ok { panic("Map Error?") }
		mk := rv.MapKeys()
		var keys sortedKey = make([]string, len(mk))
		for i, v := range mk {
			keys[i] = v.String()
		}
		sort.Sort(keys)
		for _, k := range keys {
			str += Encode(k)
			str += Encode(d[k])
		}
		str += "e"
	case reflect.Struct:
		str += "d"
		d, ok := target.(Dictionary)
		if !ok {
			panic("Not Dictionary Type")
		}
		for _, k := range d.Keys() {
			str += Encode(k)
			str += Encode(d.data[k])
		}
		str += "e"
	case reflect.Slice, reflect.Array:
		str += "l"
		for i := 0; i < rv.Len(); i++ {
			str += Encode(rv.Index(i))
		}
		str += "e"
	default:
		panic("what..?")
	}
	return str
}


func EncodeDict(dict Dictionary) []byte{
	keySet := dict.Keys()
	encoded := "d"
	for _, k := range keySet {
		encoded += Encode(k)
		encoded += Encode(dict.getValue(k))
	}
	encoded += "e"
	fmt.Println(encoded)
	return ConvertToBytes(encoded)
}