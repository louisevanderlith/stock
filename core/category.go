package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/stock/core/categories"
)

type Category struct {
	Name         string          `hsk:"size(128)"`
	Text         string          `hsk:"size(256)"`
	Description  string          `hsk:"size(512)"`
	PageURL      string          `hsk:"size(128)"`
	BaseCategory categories.Enum `hsk:"null"`
	ClientID     string
	ImageKey     *keys.TimeKey
	OwnerKey     hsk.Key
}

func (c Category) Valid() error {
	return validation.Struct(c)
}
