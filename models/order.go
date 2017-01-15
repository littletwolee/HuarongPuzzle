package models

import (
	"HuarongPuzzle/constant"
	"errors"
	"strings"
)

type Order struct {
	Order       int
	Piece       *Piece
	TypeName    string
	Point_Start []int
	Point_End   []int
}

func (o *Order) CheckCheckerboard(piece *Piece) error {
	switch o.Order {
	case constant.LEFT:
		o.Point_Start = []int{piece.PlaceStart[0], piece.PlaceStart[1] - 2}
		if o.Point_Start[1] < 0 {
			return errors.New(constant.MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0] + 2, o.Point_Start[1]}
		}
	case constant.RIGHT:
		o.Point_Start = []int{piece.PlaceStart[0], piece.PlaceStart[1] + 2}
		if o.Point_Start[1] > 7 {
			return errors.New(constant.MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0] + 2, o.Point_Start[1]}
		}
	case constant.UP:
		o.Point_Start = []int{piece.PlaceStart[0] - 2, piece.PlaceStart[1]}
		if o.Point_Start[0] < 0 {
			return errors.New(constant.MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0], o.Point_Start[1] + 2}
		}
	case constant.DOWN:
		o.Point_Start = []int{piece.PlaceStart[0] + 2, piece.PlaceStart[1]}
		if o.Point_Start[0] > 9 {
			return errors.New(constant.MAX_OUT)
		}
		if piece.Weight[1] == 2 {
			o.Point_End = []int{o.Point_Start[0], o.Point_Start[1] + 2}
		}
	}
	if !IsNPoint(o.Point_Start) && !IsNPoint(o.Point_End) {
		return errors.New(constant.OTHER_PIECE)
	}
	return nil
}

func IsNPoint(point []int) bool {
	if len(point) != 2 {
		return true
	}
	return strings.Contains("N1N2", Checkerboards[KeyFormat((point))].(Checkerboard).TypeName)
}
