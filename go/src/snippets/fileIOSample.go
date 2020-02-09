package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

func main(){
	base := "/home/wessup"
	f := OpenFile("test.txt", base)
	defer f.Close()

	cf := MakeFile("test2.txt", base)
	defer cf.Close()

	readBuffer := make([]byte, 1024)
	if f == nil {
		os.Exit(1)
	}
	for {
		cnt, err := f.Read(readBuffer)
		if err != nil && err != io.EOF{
			panic(err)
		} else if err == io.EOF || cnt == 0{
			break
		}
		fmt.Println(string(readBuffer[:cnt]))

		cf.Write(readBuffer[:cnt])

	}
}

func OpenFile(filename string, basedir ...string) *os.File{
	if basedir != nil {
		filename = path.Join(basedir[0], filename)
	}
	f, err := os.Open(filename)
	defer func(){
		r := recover()
		if r != nil {
			fmt.Println("File Open Error:", r)
		}
	}()
	if err != nil {
		panic(err)
	}
	return f
}

func MakeFile(filename string, basedir ...string) *os.File{
	if basedir != nil {
		filename = path.Join(basedir[0], filename)
	}
	f, err := os.Create(filename)
	defer func(){
		r := recover()
		if r != nil {
			fmt.Println("File Open Error:", r)
		}
	}()
	if err != nil {
		panic(err)
	}
	return f
}