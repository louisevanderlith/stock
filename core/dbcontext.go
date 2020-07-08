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
		Cars:       husk.NewTable(Car{}),
		Services:   husk.NewTable(Service{}),
		Parts:      husk.NewTable(Part{}),
		Properties: husk.NewTable(Property{}),
	}
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
	ctx.Properties.Save()
}
