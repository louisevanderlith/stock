package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type StockItem struct {
	EntityKey husk.Key
	Expires   time.Time
	Price     float32
	Tags      []Tag
	Location  string `hsk:"size(128)"`
}
