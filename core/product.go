package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type Product struct {
	CategoryKey   hsk.Key
	ItemKeys      []hsk.Key
	ShortName     string `hsk:"size(128)"`
	ImageKey      hsk.Key
	Expires       time.Time
	Currency      string
	Price         float32 //User set Value
	EstimateValue float32 //Total of all Products
	Tags          []string
	Location      string `hsk:"size(128)"`
	Views         int64  `hsk:"null"`
}

func (s Product) Valid() error {
	err := validation.Struct(s)

	if err != nil {
		return err
	}

	return PriceInBounds(s.Price, s.EstimateValue)
}
