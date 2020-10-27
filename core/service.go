package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type Service struct {
	StockItem
	Url string
}

func (o Service) Valid() error {
	return validation.Struct(o)
}
