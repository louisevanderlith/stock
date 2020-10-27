package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Property struct {
	StockItem
	Address string
}

func (p Property) Valid() error {
	return validation.Struct(p)
}
