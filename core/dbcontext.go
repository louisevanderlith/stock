package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Cars       husk.Tabler
	Services   husk.Tabler
	Parts      husk.Tabler
	Properties husk.Tabler
}

var ctx context

func CreateContext() {

	ctx = context{
		Cars:       husk.NewTable(new(Car)),
		Services:   husk.NewTable(new(Service)),
		Parts:      husk.NewTable(new(Part)),
		Properties: husk.NewTable(new(Property)),
	}
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
	ctx.Properties.Save()
}
