package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/stock/core"
)

type ServiceController struct {
	control.APIController
}

func NewServiceCtrl(ctrlmap *control.ControllerMap) *ServiceController {
	result := &ServiceController{}
	result.SetInstanceMap(ctrlmap)

	return result
}

// /v1/service/:key
func (req *ServiceController) GetByKey() {
	k := req.Ctx.Input.Param(":key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetCar(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @router /all/:pagesize [get]
func (req *ServiceController) Get() {
	page, size := req.GetPageData()
	results := core.GetServices(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ServiceController) Post() {
	var obj core.Service
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &obj)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec := obj.Create()

	req.Serve(http.StatusOK, nil, rec)
}

// @Title Update Car advert
// @Description Updates a Advert
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ServiceController) Put() {
	body := &core.Service{}
	key, err := req.GetKeyedRequest(body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = body.Update(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, nil)
}
