package models

import (
	"fmt"
)

var Checkerboards map[string]Checkerboard

type Checkerboard struct {
	TypeName string
	Name     string
}

func init() {
	Checkerboards = make(map[string]Checkerboard, 80)
}

func CheckerboardInit() {
	for t, v := range PersonTable {
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
					Checkerboards[fmt.Sprintf("%s%s", n[0], n[1])] = c
				}
			}

		}(v)
	}
}

func format(arr []int) string {
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
