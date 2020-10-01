package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
	"time"
)

type StockItem struct {
	ShortName    string `hsk:"size(128)"`
	Profile      string
	ImageKey     keys.TimeKey
	OwnerKey     keys.TimeKey //Hero
	Expires      time.Time
	Price        int64 //coins can't be divided, OR LESS THAN ZERO
	Tags         []string
	Location     string `hsk:"size(128)"`
	OwnerHistory map[time.Time]hsk.Key
	Views        int64
}

func (s StockItem) Valid() error {
	return validation.Struct(s)
}
