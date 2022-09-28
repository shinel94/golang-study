package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func (f *File) Close() {

}

func Convert(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {
	var f Reader = &File{}
	c := f.(Closer)
	fmt.Println(f, c)
	Convert(f)
}
