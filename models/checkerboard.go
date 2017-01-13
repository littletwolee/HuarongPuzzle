package models

import ()

var Checkerboard []string

func init() {
	Checkerboard = make([]string, 80)
}

func CheckerboardInit() {
	for _, v := range PersonTable {
		func(p *Piece) {
			if v.Places != nil {
				for k, n := range v.Places {
					switch k {
					case 0:
						Checkerboard[n] = p.Name[0]
					case len(v.Places) - 1:
						Checkerboard[n] = p.Name[1]
					default:
						Checkerboard[n] = "  "
					}
				}

			}
		}(v)
	}
}
