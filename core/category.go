package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/stock/core/categories"
)

type Category struct {
	Name         string
	Text         string
	Description  string
	BaseCategory categories.Enum
	ClientID     string
	ImageKey     *keys.TimeKey
	Items        []StockItem
}

func (c Category) Valid() error {
	return validation.Struct(c)
}

func (c Category) GetItem(itemKey hsk.Key) (StockItem, int, error) {
	for i, itm := range c.Items {
		if itm.ItemKey == itemKey {
			return itm, i, nil
		}
	}

	return StockItem{}, -1, errors.New("no such item")
}
