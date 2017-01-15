package models

import (
	"fmt"
)

const (
	ROWLENGTH = 8
)

var (
	ZF, MC, GY, ZY, HZ, CC, B1, B2, B3, B4, N1, N2 *Piece
	PersonTable                                    map[string]*Piece
	Weight                                         int
)

type Piece struct {
	TypeName   string
	Name       []string
	PlaceStart []int
	Places     [][]int
	Weight     int
	IsVertical int
}

func init() {
	Weight = 2
	PersonTable = make(map[string]*Piece, 12)
	PersonTable["ZF"] = &Piece{Name: []string{"张", "飞"}, PlaceStart: []int{0, 0}, Weight: 2, IsVertical: 1}
	PersonTable["HZ"] = &Piece{Name: []string{"黄", "忠"}, PlaceStart: []int{0, 6}, Weight: 2, IsVertical: 1}
	PersonTable["ZY"] = &Piece{Name: []string{"赵", "云"}, PlaceStart: []int{4, 0}, Weight: 2, IsVertical: 1}
	PersonTable["MC"] = &Piece{Name: []string{"马", "超"}, PlaceStart: []int{4, 6}, Weight: 2, IsVertical: 1}
	PersonTable["GY"] = &Piece{Name: []string{"关", "羽"}, PlaceStart: []int{4, 2}, Weight: 2, IsVertical: 0}
	PersonTable["CC"] = &Piece{Name: []string{"曹", "操"}, PlaceStart: []int{0, 2}, Weight: 2, IsVertical: 2}

	PersonTable["B1"] = &Piece{Name: []string{"士", "一"}, PlaceStart: []int{6, 2}, Weight: 1, IsVertical: 2}
	PersonTable["B2"] = &Piece{Name: []string{"士", "二"}, PlaceStart: []int{6, 4}, Weight: 1, IsVertical: 2}
	PersonTable["B3"] = &Piece{Name: []string{"士", "三"}, PlaceStart: []int{8, 0}, Weight: 1, IsVertical: 2}
	PersonTable["B4"] = &Piece{Name: []string{"士", "四"}, PlaceStart: []int{8, 6}, Weight: 1, IsVertical: 2}

	PersonTable["N1"] = &Piece{Name: []string{"  ", "  "}, PlaceStart: []int{8, 2}, Weight: 1, IsVertical: 2}
	PersonTable["N2"] = &Piece{Name: []string{"  ", "  "}, PlaceStart: []int{8, 4}, Weight: 1, IsVertical: 2}
}

func PieceInit() {
	for _, v := range PersonTable {
		i := 1
		j := 1
		switch v.IsVertical {
		case 0:
			j = 2
		case 1:
			i = 2
		case 2:
			i = 2
			j = 2
		default:
			return
		}
		v.Places = pushdata(i, j, v)
		fmt.Println(v)
	}
}
func pushdata(in, jn int, p *Piece) [][]int {
	places := [][]int{}
	for i := 0; i < in*p.Weight; i++ {
		for j := 0; j < jn*p.Weight; j++ {
			//places = append(places, s+(i*ROWLENGTH)+j)
			places = append(places, []int{(p.PlaceStart[0] + i), p.PlaceStart[1] + j})

		}
	}
	return places
}
