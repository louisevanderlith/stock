package core

import "github.com/louisevanderlith/husk"

type Part struct {
	StockItem
}

func (o Part) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
