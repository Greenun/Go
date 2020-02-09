package bencode

import (
	"log"
	"strconv"
	"strings"
)

func DecodeString(target string) (string, string) {
	idx := strings.IndexByte(target, ':')
	readCount, err := strconv.Atoi(target[:idx])
	if err != nil {
		panic("String Decode Error")
	}
	return target[idx+1:idx+1+readCount], target[idx+1+readCount:]
}

func DecodeInt(target string) (int, string) {
	idx := strings.IndexByte(target, 'e')
	num, err := strconv.Atoi(target[1:idx])
	if err != nil {
		panic("Int Decode Error")
	}
	return num, target[idx+1:]
}

func DecodeData(encoded string) (interface{}, string){
	remain := encoded[:]
	log.Println(remain)
	switch remain[0]{
	case 'd':
		//to dict
		remain = remain[1:]
		temp := make(map[string]interface{})
		for {
			if remain[0] == 'e' {
				remain = remain[1:]
				break
			}
			var (
				key interface{}
				value interface{}
			)
			key, remain = DecodeData(remain)
			value, remain = DecodeData(remain)

			k, ok := key.(string)
			if !ok { panic("key has to be string")}
			temp[k] = value
		}
		return temp, remain
	case 'i':
		// to int
		return DecodeInt(remain)
	case 'l':
		// to list
		// may change it to linked list
		var temp []interface{}
		for {
			if remain[0] == 'e' { break }
			var element interface{}
			element, remain = DecodeData(remain)
			temp = append(temp, element)
		}
		return temp, remain
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		// to str
		return DecodeString(remain)
	default:
		panic("Decoding Error")
		return nil, ""
	}
}

func DecodeStream(encoded []byte) interface{}{
	// Decode Byte Data To Dict
	encodedString := ConvertToString(encoded)
	decoded, remain := DecodeData(encodedString)
	if remain != "" {
		panic("Stream remains after decode")
	}
	// todo : interface --> Dictionary Type / more Test
	return decoded
}