package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Cars husk.Tabler
}

var ctx context

func CreateContext() {

	ctx = context{
		Cars: husk.NewTable(new(CarAdvert)),
	}
}

func Shutdown() {
	ctx.Cars.Save()
}
