package routers

import (
	"github.com/louisevanderlith/stock/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	carCtrl := &controllers.Cars{}
	partCtrl := &controllers.Parts{}
	propCtrl := &controllers.Properties{}
	srvCtrl := &controllers.Services{}
	e.JoinBundle("/", roletype.Owner, mix.JSON, carCtrl, partCtrl, propCtrl, srvCtrl)
}
