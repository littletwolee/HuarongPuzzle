package models

import (
	"fmt"
)

var (
	CMDList []string
	CMDInfo map[string]map[string]string
	Helps   []string
)

func DisplayInit() {
	CMDInfo = make(map[string]map[string]string, 10)
	Persons := make(map[string]string, 10)
	Persons["CC"] = "Generals CaoCao"
	Persons["GY"] = "Generals GuanYu"
	Persons["ZY"] = "Generals ZhaoYun"
	Persons["HZ"] = "Generals HuangZhong"
	Persons["MC"] = "Generals MaoChao"
	Persons["ZF"] = "Generals ZhangFei"
	Persons["B1"] = "The first Shi"
	Persons["B2"] = "The Second Shi"
	Persons["B3"] = "The third Shi"
	Persons["B4"] = "The fourth Shi"
	CMDInfo["Persons"] = Persons
	Moves := make(map[string]string, 4)
	Moves["U"] = "Move Up"
	Moves["D"] = "Move Down"
	Moves["L"] = "Move Left"
	Moves["R"] = "Move Right"
	CMDInfo["Moves"] = Moves
	Cmds := make(map[string]string, 2)
	Cmds["H"] = "Game help"
	Cmds["EXIT"] = "Exit game"
	CMDInfo["Cmds"] = Cmds
	Helps = []string{}
	Helps = append(Helps, "\n")
	for t, vs := range CMDInfo {
		Helps = append(Helps, fmt.Sprintf("    %s -", t))
		for k, v := range vs {
			Helps = append(Helps, fmt.Sprintf("      %s - %s", k, v))
			CMDList = append(CMDList, k)
		}
	}
	Helps = append(Helps, "\n")
	DisplayCheckerboard()
	DisplayHelp()
}

func DisplayHelp() {
	for _, v := range Helps {
		fmt.Println(v)
	}
}

func DisplayCheckerboard() {
	for i := 0; i < 10; i++ {
		str := "      "
		for j := 0; j < 8; j++ {
			str += Checkerboards[KeyFormat([]int{i, j})].(Checkerboard).Name
		}
		fmt.Println(str)

	}
}
