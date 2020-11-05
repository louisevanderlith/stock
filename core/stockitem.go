package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type StockItem struct {
	ItemKey       *keys.TimeKey
	ShortName     string `hsk:"size(128)"`
	ImageKey      *keys.TimeKey
	OwnerKey      *keys.TimeKey //Hero
	Expires       time.Time
	Currency      string
	Price         float32
	EstimateValue float32
	Tags          []string
	Location      string `hsk:"size(128)"`
	OwnerHistory  map[time.Time]hsk.Key
	Views         int64 `hsk:"null"`
}

func (s StockItem) Valid() error {
	err := validation.Struct(s)

	if err != nil {
		return err
	}

	return PriceInBounds(s.Price, s.EstimateValue)
}
