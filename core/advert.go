package core

import (
	"time"

	"github.com/louisevanderlith/husk"
)

type Advert struct {
	UserKey     husk.Key
	Expires time.Time
	Price      float32
	Tags       []Tag
	Location   string `hsk:"size(128)"`
}