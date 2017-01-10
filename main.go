package main

import (
	"fmt"
	"os"
	"unsafe"
)

//var I chan []byte

func main() {
	gamelogo()
	b := make([]byte, 100)
	f := os.Stdin
	w := os.Stdout
	defer f.Close()
	defer w.Close()
	for {
		w.WriteString("input:")
		c, err := f.Read(b)
		if err != nil {
			continue
		}
		bb := b[:c-2]
		b = make([]byte, 100)
		str := *(*string)(unsafe.Pointer(&bb))
		fmt.Println(str)
		if str == "exit" {
			break
		}
	}

}

func gamelogo() {
	fmt.Println("***********************************************************************************")
	fmt.Println("*      **    **  **    **     **       ******     ****    ****      **   ******   *")
	fmt.Println("*     **    **  **    **   **  **     **    **  **   **  ** **     **  **         *")
	fmt.Println("*    ********  **    **  **    **    **    **  **   **  **   **   **  **  *****   *")
	fmt.Println("*   **    **   **   **  *********   **  ***   **   **  **     ** **  **     **    *")
	fmt.Println("*  **    **     ****   **      **  **    **    ****   **      ****     *****      *")
	fmt.Println("*                                                                                 *")
	fmt.Println("*      *****     **    **                                                                    *")
	fmt.Println("*     **    **  **    **                                                           *")
	fmt.Println("*    **    **  **    **                                                          *")
	fmt.Println("*   ******    **   **                                                         *")
	fmt.Println("*  **          ****                                                        *")
	fmt.Println("*************************************************************************")
}

// fun Handle(){

// }
