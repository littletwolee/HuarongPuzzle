package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	b := make([]byte, 100)
	f := os.Stdin
	w := os.Stdout
	defer f.Close()
	defer w.Close()
	for {
		w.WriteString("input:")
		c, _ := f.Read(b)
		bb := b[:c-2]
		str := *(*string)(unsafe.Pointer(&bb))
		fmt.Println(str)
		if str == "exit" {
			break
		}
	}
}
