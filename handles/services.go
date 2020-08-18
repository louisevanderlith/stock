package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/stock/core"
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	idn := drx.GetUserIdentity(r)
	k, err := husk.ParseKey(idn.GetUserID())

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	results, err := core.GetServices(1, 10, k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

// /v1/service/:key
func ViewServices(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := husk.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetService(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

// @router /all/:pagesize [get]
func SearchServices(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	idn := drx.GetUserIdentity(r)
	k, err := husk.ParseKey(idn.GetUserID())

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	results, err := core.GetServices(page, size, k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateServices(w http.ResponseWriter, r *http.Request) {
	var obj core.Service
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

// @Title Update Car advert
// @Description Updates a Advert
// @Param	body		body 	core.CarAdvert	true		"body for ad content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func UpdateServices(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := &core.Service{}
	err = drx.JSONBody(r, body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = body.Update(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println(err)
	}
}
