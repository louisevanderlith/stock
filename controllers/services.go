package controllers

import (
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/stock/core"
)

type Services struct {
}

func (req *Services) Get(c *gin.Context) {
	results := core.GetServices(1, 10)

	return http.StatusOK, results
}

// /v1/service/:key
func (req *Services) View(c *gin.Context) {
	k := c.Param("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := core.GetService(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// @router /all/:pagesize [get]
func (req *Services) Search(c *gin.Context) {
	page, size := ctx.GetPageData()
	results := core.GetServices(page, size)

	return http.StatusOK, results
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Services) Create(c *gin.Context) {
	var obj core.Service
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
func (req *Services) Update(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	body := &core.Service{}
	err = ctx.Body(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
