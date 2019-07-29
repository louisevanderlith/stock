package core

import (
	"github.com/louisevanderlith/husk"
)

type Property struct {
	StockItem
	Address string
}

func (p Property) Valid() (bool, error) {
	return husk.ValidateStruct(&p)
}
