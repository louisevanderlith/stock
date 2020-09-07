package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Part struct {
	StockItem
	Number string
}

func (o Part) Valid() error {
	return validation.Struct(o)
}

func (c Part) Create() (hsk.Key, error) {
	return ctx.Parts.Create(c)
}

func (c Part) Update(key hsk.Key) error {
	return ctx.Parts.Update(key, c)
}
