package core

import (
	"github.com/louisevanderlith/husk"
)

type Property struct {
	StockItem
	Address string
}

func (p Property) Valid() error {
	return husk.ValidateStruct(p)
}
