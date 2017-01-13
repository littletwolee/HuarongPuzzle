package main

import (
	"HuarongPuzzle/models"
	"fmt"
)

func init() {
	fmt.Println(1)

}

func main() {
	models.PieceInit()
	models.CheckerboardInit()
	for i := 0; i < 10; i++ {
		fmt.Println(models.Checkerboard[i*8 : (i+1)*8])
	}
	// for _, v := range models.Checkerboard {

	// }

}
