package main

import (
	"HuarongPuzzle/constant"
	"HuarongPuzzle/models"
	"fmt"
	"strings"
)

var (
	CmdFlag bool
	O       *models.Order
)

func init() {
	CmdFlag = false
	models.PieceInit()
	models.CheckerboardInit("", nil)
	models.DisplayInit()
}

func main() {
Loop:
	for {
		cmd := getorder()
		switch cmd {
		case "H":
			models.DisplayHelp()
			continue
		case "N":
			models.NewGame()
			CmdFlag = false
			fmt.Println("")
			models.PieceInit()
			models.CheckerboardInit("", nil)
			models.DisplayInit()
			models.FileDelete("data")
		case "EXIT":
			break Loop
		case "":
			fmt.Println(constant.ERR_INPUT_EMPTY)
			continue
		default:
			switch {
			case !CmdFlag && models.CMDInfo["Persons"][cmd] != "":
				O = &models.Order{TypeName: cmd}
			case CmdFlag && models.CMDInfo["Moves"][cmd] != "":
				fmt.Println("")
				O.Order = cmd
				err := Move(O)
				if err != nil {
					fmt.Println(err)
				}
				models.DisplayCheckerboard()
				if models.SuccessCheckerboard() {
					fmt.Println(constant.SUCCESS)
					models.FileDelete("data")
					break Loop
				}
			default:
				if !CmdFlag {
					fmt.Println(constant.ERR_INPUT_GENERALS)
				} else {
					fmt.Println(constant.ERR_INPUT_CMD)
				}
				continue
			}
			CmdFlag = !CmdFlag
			fmt.Println("")
		}
	}
}

func getorder() string {
	var cmd string
	if !CmdFlag {
		fmt.Printf(constant.INPUT_GENERALS)
	} else {
		fmt.Printf(constant.INPUT_CMDS)
	}
	fmt.Scanf("%s", &cmd)
	return strings.ToUpper(cmd)
}
