package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/stock/controllers"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Car
	carCtrl := &controllers.CarController{}
	carGroup := droxolite.NewRouteGroup("Car", carCtrl)
	carGroup.AddRoute("Car by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, carCtrl.GetByKey)
	carGroup.AddRoute("Create Car", "", "POST", roletype.Owner, carCtrl.Post)
	carGroup.AddRoute("All Cars", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, carCtrl.Get)
	poxy.AddGroup(carGroup)

	//Part
	partCtrl := &controllers.PartController{}
	partGroup := droxolite.NewRouteGroup("part", partCtrl)
	partGroup.AddRoute("Part by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, partCtrl.GetByKey)
	partGroup.AddRoute("Create Part", "", "POST", roletype.Owner, partCtrl.Post)
	partGroup.AddRoute("All Parts", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, partCtrl.Get)
	poxy.AddGroup(partGroup)

	//Property
	propCtrl := &controllers.PropertyController{}
	propGroup := droxolite.NewRouteGroup("property", propCtrl)
	//propGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, propCtrl.GetByKey)
	//propGroup.AddRoute("/", "POST", roletype.Owner, propCtrl.Post)
	//propGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, propCtrl.Get)
	poxy.AddGroup(propGroup)

	//Service
	srvCtrl := &controllers.ServiceController{}
	srvGroup := droxolite.NewRouteGroup("service", srvCtrl)
	srvGroup.AddRoute("Service by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, srvCtrl.GetByKey)
	srvGroup.AddRoute("Create Service", "", "POST", roletype.Owner, srvCtrl.Post)
	srvGroup.AddRoute("All Services", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, srvCtrl.Get)
	poxy.AddGroup(srvGroup)

	/*
		ctrlmap := EnableFilters(s, host)

		carsCtrl := controllers.NewCarCtrl(ctrlmap)
		//beego.Router("/v1/car", uplCtrl, "post:Post")
		beego.Router("/v1/car/:key", carsCtrl, "get:GetByKey")
		beego.Router("/v1/car/all/:pagesize", carsCtrl, "get:Get")

		partsCtrl := controllers.NewPartCtrl(ctrlmap)
		//beego.Router("/v1/car", uplCtrl, "post:Post")
		beego.Router("/v1/part/:key", partsCtrl, "get:GetByKey")
		beego.Router("/v1/part/all/:pagesize", partsCtrl, "get:Get")

		servCtrl := controllers.NewServiceCtrl(ctrlmap)
		//beego.Router("/v1/car", uplCtrl, "post:Post")
		beego.Router("/v1/service/:key", servCtrl, "get:GetByKey")
		beego.Router("/v1/service/all/:pagesize", servCtrl, "get:Get")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner

	ctrlmap.Add("/v1/car", emptyMap)
	ctrlmap.Add("/v1/part", emptyMap)
	ctrlmap.Add("/v1/service", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
