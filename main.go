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
		var str string
		for j := 0; j < 8; j++ {
			str += models.Checkerboards[fmt.Sprintf("%s%s", i, j)].Name
		}
		fmt.Println(str)
	}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(models.Checkerboard[i*8 : (i+1)*8])
	// }
	// for _, v := range models.PersonTable {
	// 	fmt.Println(v)
	// }
	// m := map[int]int{123: 456}
	// for i, j := range m {
	// 	fmt.Println(i)
	// 	fmt.Println(j)
	// }
}
