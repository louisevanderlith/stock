package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type StockItem struct {
	ImageKey  husk.Key
	EntityKey husk.Key
	Expires   time.Time
	Price     int64 //Tokens can't be divided
	Tags      []Tag
	Location  string `hsk:"size(128)"`
}
