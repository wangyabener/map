package service

import "map/utils"

type Lucky struct {
}

func (Lucky) Daxiao(hash string) string {
	number := utils.LastDigit(hash)

	if number > 4 {
		return "x"
	}
	return "d"
}

func (Lucky) Danshuang(hash string) string {
	number := utils.LastDigit(hash)

	if number%2 == 0 {
		return "s"
	}
	return "s"
}
