package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/louisevanderlith/ads/core"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/control"
)

type CarController struct {
	control.APIController
}

func NewCarCtrl(ctrlmap *control.ControllerMap) *CarController {
	result := &CarController{}
	result.SetInstanceMap(ctrlmap)

	return result
}

// /v1/car/:key
func (req *CarController) GetByKey() {
	k := req.Ctx.Input.Param(":key")
	key, err := husk.ParseKey(k)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := core.GetCarAdvert(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @router /all/:pagesize [get]
func (req *CarController) Get() {
	page, size := req.GetPageData()
	results := core.GetLatestCars(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *CarController) Post() {
	var obj core.CarAdvert
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
func (req *CarController) Put() {
	body := &core.CarAdvert{}
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
