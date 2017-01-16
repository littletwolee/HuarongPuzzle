package main

import (
	"HuarongPuzzle/constant"
	"HuarongPuzzle/models"
	"fmt"
	"github.com/qianlnk/terminal"
	"strings"
)

var (
	CmdFlag bool
	Term    *terminal.Terminal
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
		case "EXIT":
			break Loop
		case "":
			fmt.Println(constant.ERR_INPUT_EMPTY)
			continue
		default:
			o := models.GetOrder()
			switch {
			case !CmdFlag && models.CMDInfo["Persons"][cmd] != "":
				o.TypeName = cmd
			case CmdFlag && models.CMDInfo["Moves"][cmd] != "":
				fmt.Println("")
				o.Order = cmd
				err := Move()
				if err != nil {
					fmt.Println(err)
				}
				models.DisplayCheckerboard()
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

	if Term == nil {
		Term = terminal.NewTerminal(">")
		Term.SetSystemCommand(models.CMDList)
	}
	if !CmdFlag {
		fmt.Printf("Input Generals Shorthand > ")
	} else {
		fmt.Printf("Input cmd order > ")
	}
	return strings.ToUpper(Term.GetCommand())
}
