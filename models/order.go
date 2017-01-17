package models

import (
	"HuarongPuzzle/constant"
	"errors"
	"strings"
)

var (
	TheOrder *Order
)

type Order struct {
	Order           string
	Piece           *Piece
	TypeName        string
	Point_Start     []int
	Point_End       []int
	Point_Nil_Start []int
	Point_Nil_End   []int
}

func GetOrder() *Order {
	if TheOrder == nil {
		TheOrder = &Order{}
	}
	return TheOrder
}

func (o *Order) CheckCheckerboard(piece *Piece) error {
	switch o.Order {
	case constant.LEFT:
		o.Point_Start = []int{piece.PlaceStart[0], piece.PlaceStart[1] - 2}
		o.Point_Nil_Start = o.Point_Start
		if o.Point_Nil_Start[1] < 0 {
			return errors.New(constant.ERR_MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0] + 2, o.Point_Start[1]}
			o.Point_Nil_End = o.Point_End
		}
	case constant.RIGHT:
		o.Point_Start = []int{piece.PlaceStart[0], piece.PlaceStart[1] + 2}
		o.Point_Nil_Start = []int{piece.PlaceStart[0], piece.PlaceStart[1] + (2 * piece.Weight[0])}
		if o.Point_Nil_Start[1] > 7 {
			return errors.New(constant.ERR_MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0] + 2, o.Point_Start[1]}
			o.Point_Nil_End = []int{o.Point_Nil_Start[0] + 2, o.Point_Nil_Start[1]}
		}
	case constant.UP:
		o.Point_Start = []int{piece.PlaceStart[0] - 2, piece.PlaceStart[1]}
		o.Point_Nil_Start = o.Point_Start
		if o.Point_Nil_Start[0] < 0 {
			return errors.New(constant.ERR_MAX_OUT)
		}
		if piece.Weight[0] == 2 {
			o.Point_End = []int{o.Point_Start[0], o.Point_Start[1] + 2}
			o.Point_Nil_End = o.Point_End
		}
	case constant.DOWN:
		o.Point_Start = []int{piece.PlaceStart[0] + 2, piece.PlaceStart[1]}
		o.Point_Nil_Start = []int{piece.PlaceStart[0] + (2 * piece.Weight[1]), piece.PlaceStart[1]}
		if o.Point_Nil_Start[0] > 9 {
			return errors.New(constant.ERR_MAX_OUT)
		}
		if piece.Weight[0] == 2 {
			o.Point_End = []int{o.Point_Start[0], o.Point_Start[1] + 2}
			o.Point_Nil_End = []int{o.Point_Nil_Start[0], o.Point_Nil_Start[1] + 2}
		}
	}
	if !IsNPoint(o.Point_Nil_Start) || !IsNPoint(o.Point_Nil_End) {
		return errors.New(constant.ERR_OTHER_PIECE)
	}
	return nil
}

func IsNPoint(point []int) bool {
	if len(point) != 2 {
		return true
	}
	return strings.Contains("N1N2", Checkerboards[KeyFormat((point))].(Checkerboard).TypeName)
}
