package models

import (
	"fmt"
)

const (
	ROWLENGTH = 8
)

var (
	ZF, MC, GY, ZY, HZ, CC, B1, B2, B3, B4, N1, N2 *Piece
	PersonTable                                    []*Piece
	Weight                                         int
)

type Piece struct {
	TypeName   string
	Name       []string
	PlaceStart int
	Places     []int
	Weight     int
	IsVertical int
}

func init() {
	fmt.Println(2)
	Weight = 2
	PersonTable = make([]*Piece, 12)
	PersonTable[0] = &Piece{TypeName: "ZF", Name: []string{"张", "飞"}, PlaceStart: 0, Weight: 2, IsVertical: 1}
	PersonTable[1] = &Piece{TypeName: "HZ", Name: []string{"黄", "忠"}, PlaceStart: 6, Weight: 2, IsVertical: 1}
	PersonTable[2] = &Piece{TypeName: "ZY", Name: []string{"赵", "云"}, PlaceStart: 32, Weight: 2, IsVertical: 1}
	PersonTable[3] = &Piece{TypeName: "MC", Name: []string{"马", "超"}, PlaceStart: 38, Weight: 2, IsVertical: 1}
	PersonTable[4] = &Piece{TypeName: "GY", Name: []string{"关", "羽"}, PlaceStart: 34, Weight: 2, IsVertical: 0}
	PersonTable[5] = &Piece{TypeName: "CC", Name: []string{"曹", "操"}, PlaceStart: 2, Weight: 2, IsVertical: 2}

	PersonTable[6] = &Piece{TypeName: "B1", Name: []string{"士", "兵"}, PlaceStart: 50, Weight: 1, IsVertical: 2}
	PersonTable[7] = &Piece{TypeName: "B2", Name: []string{"士", "兵"}, PlaceStart: 52, Weight: 1, IsVertical: 2}
	PersonTable[8] = &Piece{TypeName: "B3", Name: []string{"士", "兵"}, PlaceStart: 64, Weight: 1, IsVertical: 2}
	PersonTable[9] = &Piece{TypeName: "B4", Name: []string{"士", "兵"}, PlaceStart: 70, Weight: 1, IsVertical: 2}

	PersonTable[10] = &Piece{TypeName: "N1", Name: []string{"  ", "  "}, PlaceStart: 66, Weight: 1, IsVertical: 2}
	PersonTable[11] = &Piece{TypeName: "N2", Name: []string{"  ", "  "}, PlaceStart: 68, Weight: 1, IsVertical: 2}
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
		v.Places = pushdata(i, j, v.Weight, v.PlaceStart)
	}
}
func pushdata(in, jn, w, s int) []int {
	places := []int{}
	for i := 0; i < in*w; i++ {
		for j := 0; j < jn*w; j++ {
			places = append(places, s+(i*ROWLENGTH)+j)
		}
	}
	return places
}
