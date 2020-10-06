package core

import "github.com/louisevanderlith/husk/validation"

type Clothing struct {
	StockItem
	Category string
	Type     string
	Size     string
	Colour   string
}

func (c Clothing) Validate() error {
	return validation.Struct(c)
}
