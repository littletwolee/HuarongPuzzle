package models

import (
	"fmt"
)

var Checkerboards map[string]interface{}

type Checkerboard struct {
	TypeName string
	Name     string
}

func init() {
	Checkerboards = make(map[string]interface{}, 80)
}

func CheckerboardInit(typename string, pieces *Piece) {
	if pieces != nil {
		pushCheckerBoard(typename, pieces)
	}
	for t, v := range PersonTable {
		pushCheckerBoard(t, v)
	}
}

func pushCheckerBoard(t string, v *Piece) {
	func(p *Piece) {
		if v.Places != nil {
			l := len(v.Places)
			for k, n := range v.Places {
				c := Checkerboard{TypeName: t}
				switch k {
				case 0:
					c.Name = p.Name[0]
				case l - 1:
					c.Name = p.Name[1]
				default:
					c.Name = "  "
				}
				Checkerboards[KeyFormat(n)] = c
			}
		}

	}(v)
}

func KeyFormat(arr []int) string {
	str := ""
	for _, v := range arr {
		if v < 10 {
			str += fmt.Sprintf("0%d", v)
		} else {
			str += fmt.Sprintln(v)
		}
		str += ","
	}
	return str
}
