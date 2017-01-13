package main

import (
	"HuarongPuzzle/models"
	"fmt"
)

func init() {
	models.PieceInit()
	models.CheckerboardInit()

}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(models.Checkerboard[i*8 : (i+1)*8])
	}
	// for _, v := range models.Checkerboard {

	// }

}
