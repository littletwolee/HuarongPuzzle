package main

import (
	"HuarongPuzzle/models"
	"fmt"
)

func main() {
	models.DataInit()
	for i := 0; i < 10; i++ {
		fmt.Println(models.Checkerboard[i*8 : (i+1)*8])
	}
	// for _, v := range models.Checkerboard {

	// }

}
