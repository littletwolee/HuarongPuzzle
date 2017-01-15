package main

import (
	"HuarongPuzzle/constant"
	"HuarongPuzzle/models"
	"fmt"
)

func init() {
	models.PieceInit()
	models.CheckerboardInit("", nil)

}

func main() {
	for i := 0; i < 10; i++ {
		var str string
		for j := 0; j < 8; j++ {
			str += models.Checkerboards[models.KeyFormat([]int{i, j})].(models.Checkerboard).Name
		}
		fmt.Println(str)

	}
	err := Move("GY", constant.DOWN)
	fmt.Println(models.PersonTable["GY"])
	fmt.Println(models.PersonTable["N1"])
	fmt.Println(models.PersonTable["N2"])
	fmt.Println(models.Checkerboards)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < 10; i++ {
			var str string
			for j := 0; j < 8; j++ {
				str += models.Checkerboards[models.KeyFormat([]int{i, j})].(models.Checkerboard).Name
			}
			fmt.Println(str)
		}
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
