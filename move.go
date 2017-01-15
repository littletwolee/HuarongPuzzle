package main

import (
	"HuarongPuzzle/models"
	//"fmt"
	"strconv"
	"strings"
)

func Move(typename string, order int) error {
	typename = strings.ToUpper(typename)
	o := &models.Order{Order: order}
	o.Piece = models.PersonTable[typename]
	o.TypeName = typename
	err := o.CheckCheckerboard(o.Piece)
	if err != nil {
		return err
	}
	o.Piece.PlaceStart = o.Point_Start
	err = pointNil(o)
	if err != nil {
		return err
	}
	return nil
}

func pointNil(o *models.Order) error {
	typenames := []string{o.TypeName}
	if len(o.Point_End) > 0 {
		typenames = append(typenames, "N1", "N2")
	} else {
		typenames = append(typenames, models.Checkerboards[models.KeyFormat(o.Point_Start)].(models.Checkerboard).TypeName)
	}
	for _, n := range typenames {
		piece := models.PersonTable[n]
		for _, p := range piece.Places {
			models.Checkerboards[models.KeyFormat(p)] = nil
		}
	}
	return pushpoint(o, typenames[1:])
}

func pushpoint(o *models.Order, typenames []string) error {
	models.PersonTable[o.TypeName].Places = models.PushData(o.Piece)
	models.CheckerboardInit(o.TypeName, o.Piece)
	nilarr := map[string][]int{}
	for k, v := range models.Checkerboards {
		if v == nil {
			strs := strings.Split(k, ",")
			r, err := strconv.Atoi(strs[0])
			if err != nil {
				return err
			}
			c, err := strconv.Atoi(strs[1])
			if err != nil {
				return err
			}
			nilarr[models.KeyFormat([]int{r, c})] = []int{r, c}
		}
	}
	pushN(nilarr, o, typenames[0])
	if len(typenames) > 1 {
		for _, v := range models.PersonTable[typenames[0]].Places {
			delete(nilarr, models.KeyFormat(v))
		}

		pushN(nilarr, o, typenames[1])
	}
	return nil
}

func pushN(arr map[string][]int, o *models.Order, typename string) {
	s := []int{}
	for _, v := range arr {
		if len(s) == 0 {
			s = v
			continue
		}
		if s[0] > v[0] || s[1] > v[1] {
			s = v
		}
	}
	models.PersonTable[typename].PlaceStart = s
	models.PersonTable[typename].Places = models.PushData(models.PersonTable[typename])
	models.CheckerboardInit(typename, models.PersonTable[typename])
}
