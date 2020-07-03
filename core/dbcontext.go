package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
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
		Cars:       husk.NewTable(Car{}, serials.GobSerial{}),
		Services:   husk.NewTable(Service{}, serials.GobSerial{}),
		Parts:      husk.NewTable(Part{}, serials.GobSerial{}),
		Properties: husk.NewTable(Property{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Cars.Save()
	ctx.Services.Save()
	ctx.Parts.Save()
	ctx.Properties.Save()
}
