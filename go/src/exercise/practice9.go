package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	//hashfunc := crypto.SHA256.New()

	hashf := sha1.New()
	hashf.Write([]byte("abcde"))
	h := hashf.Sum(nil)
	fmt.Println(h)
}
