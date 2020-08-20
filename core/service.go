package core

import "github.com/louisevanderlith/husk"

type Service struct {
	StockItem
	Url string
}

func (o Service) Valid() error {
	return husk.ValidateStruct(&o)
}
