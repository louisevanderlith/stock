package core

import "github.com/louisevanderlith/husk"

type Service struct {
	StockItem
}

func (o Service) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
