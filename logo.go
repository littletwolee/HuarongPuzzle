package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	// "os"
	// "unsafe"
)

func LogoInit() {
	//gamelogo()
	// b := make([]byte, 100)
	// f := os.Stdin
	// w := os.Stdout
	// defer f.Close()
	// defer w.Close()
	// for {
	// 	w.WriteString("input:")
	// 	c, err := f.Read(b)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	bb := b[:c-1]
	// 	b = make([]byte, 100)
	// 	str := *(*string)(unsafe.Pointer(&bb))
	// 	fmt.Println(str)
	// 	if str == "exit" {
	// 		break
	// 	}
	// }
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				fmt.Println("You press 1")
			case termbox.KeyArrowDown:
				fmt.Println("You press F1")

			default:
				fmt.Println(ev.Key)
				//break Loop
				break Loop
			}
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
