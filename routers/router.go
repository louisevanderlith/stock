package routers

import (
	"github.com/louisevanderlith/stock/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Car
	carCtrl := &controllers.Cars{}
	carGroup := routing.NewRouteGroup("Car", mix.JSON)
	carGroup.AddRoute("Car by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, carCtrl.GetByKey)
	carGroup.AddRoute("Create Car", "", "POST", roletype.Owner, carCtrl.Post)
	carGroup.AddRoute("All Cars", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, carCtrl.Get)
	poxy.AddGroup(carGroup)

	//Part
	partCtrl := &controllers.Parts{}
	partGroup := routing.NewRouteGroup("part", mix.JSON)
	partGroup.AddRoute("Part by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, partCtrl.GetByKey)
	partGroup.AddRoute("Create Part", "", "POST", roletype.Owner, partCtrl.Post)
	partGroup.AddRoute("All Parts", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, partCtrl.Get)
	poxy.AddGroup(partGroup)

	//Property
	//propCtrl := &controllers.Properties{}
	propGroup := routing.NewRouteGroup("property", mix.JSON)
	//propGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, propCtrl.GetByKey)
	//propGroup.AddRoute("/", "POST", roletype.Owner, propCtrl.Post)
	//propGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, propCtrl.Get)
	poxy.AddGroup(propGroup)

	//Service
	srvCtrl := &controllers.Services{}
	srvGroup := routing.NewRouteGroup("service", mix.JSON)
	srvGroup.AddRoute("Service by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, srvCtrl.GetByKey)
	srvGroup.AddRoute("Create Service", "", "POST", roletype.Owner, srvCtrl.Post)
	srvGroup.AddRoute("All Services", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, srvCtrl.Get)
	poxy.AddGroup(srvGroup)
}
