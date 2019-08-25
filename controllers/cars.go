package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/stock/core"
)

type Cars struct {
}

// /v1/car/:key
func (req *Cars) GetByKey(ctx context.Contexer) (int, interface{}) {
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetCar(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// @router /all/:pagesize [get]
func (req *Cars) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetLatestCars(page, size)

	return http.StatusOK, results
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Cars) Post(ctx context.Contexer) (int, interface{}) {
	var obj core.Car
	err := ctx.Body(&obj)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec := obj.Create()

	return http.StatusOK, rec
}

// @Title Update Car advert
// @Description Updates a Advert
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Cars) Put(ctx context.Contexer) (int, interface{}) {
	body := &core.Car{}
	key, err := ctx.GetKeyedRequest(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
