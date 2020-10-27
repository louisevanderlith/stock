package handles

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"

	"github.com/louisevanderlith/stock/core"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	results, err := core.Context().FindLatestCars(1, 10, claims["azp"].(string))

	if err != nil {
		log.Println("Find Cars Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ViewCar(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := keys.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().GetCar(key)

	if err != nil {
		log.Println("Get Service Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchCars(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	results, err := core.Context().FindLatestCars(page, size, claims["azp"].(string))

	if err != nil {
		log.Println("Find Parts Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var obj core.Car
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().CreateCar(obj)

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Car{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateCar(key, obj)

	if err != nil {
		log.Println("Update Service Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
