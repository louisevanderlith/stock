package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Cars     husk.Tabler
	Services husk.Tabler
	Parts    husk.Tabler
}

var ctx context

func CreateContext() {

	ctx = context{
		Cars:     husk.NewTable(new(Car)),
		Services: husk.NewTable(new(Service)),
		Parts:    husk.NewTable(new(Part)),
	}
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
}
