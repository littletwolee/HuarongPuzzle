package models

import ()

var (
	ZF, MC, GY, ZY, HZ, CC, B1, B2, B3, B4, N1, N2 *Piece
	PersonTable                                    []*Piece
	Weight                                         int
)

type Piece struct {
	Name       string
	PlaceStart int
	Places     []int
	Weight     int
	IsVertical int
}

func init() {
	Weight = 2
	PersonTable = make([]*Piece, 12)
	PersonTable[0] = &Piece{Name: "ZF", PlaceStart: 0, Weight: 2, IsVertical: 1}
	PersonTable[1] = &Piece{Name: "HZ", PlaceStart: 6, Weight: 2, IsVertical: 1}
	PersonTable[2] = &Piece{Name: "ZY", PlaceStart: 32, Weight: 2, IsVertical: 1}
	PersonTable[3] = &Piece{Name: "MC", PlaceStart: 38, Weight: 2, IsVertical: 1}
	PersonTable[4] = &Piece{Name: "GY", PlaceStart: 34, Weight: 2, IsVertical: 0}
	PersonTable[5] = &Piece{Name: "CC", PlaceStart: 2, Weight: 2, IsVertical: 2}

	PersonTable[6] = &Piece{Name: "B1", PlaceStart: 50, Weight: 1, IsVertical: 2}
	PersonTable[7] = &Piece{Name: "B2", PlaceStart: 52, Weight: 1, IsVertical: 2}
	PersonTable[8] = &Piece{Name: "B3", PlaceStart: 64, Weight: 1, IsVertical: 2}
	PersonTable[9] = &Piece{Name: "B4", PlaceStart: 70, Weight: 1, IsVertical: 2}

	PersonTable[10] = &Piece{Name: "N1", PlaceStart: 66, Weight: 1, IsVertical: 2}
	PersonTable[11] = &Piece{Name: "N2", PlaceStart: 68, Weight: 1, IsVertical: 2}
	PersonTable[4] = dataplace(PersonTable[4])
}

func DataInit() {
	for _, v := range PersonTable {
		func(p *Piece) {
			if v.Places != nil {
				for _, v := range v.Places {
					Checkerboard[v] = p.Name
				}
			}
		}(v)
	}
}

func dataplace(p *Piece) *Piece {
	places := make([]int, p.Weight*2*2)
	switch p.IsVertical {
	case 0:
		n := 0
		for i := 0; i < 2; i++ {
			for j := 0; j < 2*p.Weight; j++ {
				places[(j+1)*(i+1)] = p.PlaceStart + (j+1)*(i+1) + n
			}
			n += 8
		}
		return p
	case 1:
		return nil
	case 2:
		return nil
	default:
		return nil
	}
}
