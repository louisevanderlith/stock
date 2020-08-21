package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type StockItem struct {
	Profile      string
	ImageKey     husk.Key
	OwnerKey     husk.Key //Hero
	Expires      time.Time
	Price        int64 //coins can't be divided
	Tags         []string
	Location     string `hsk:"size(128)"`
	OwnerHistory map[time.Time]husk.Key
}

func (s StockItem) Valid() error {
	return husk.ValidateStruct(&s)
}
