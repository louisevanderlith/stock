package handles

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/core"
	"log"
	"net/http"
)

func GetClientCategories(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("user").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)

	result, err := core.Context().FindCategoriesByClient(1, 10, claims["clientId"].(string))

	if err != nil {
		log.Println("Get Categories Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchClientCategories(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	usr := r.Context().Value("user").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)

	result, err := core.Context().FindCategoriesByClient(page, size, claims["clientId"].(string))

	if err != nil {
		log.Println("Search Categories Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchCategoryNameProducts(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	category := drx.FindParam(r, "category")
	result, err := core.Context().FindProductsByCategoryName(page, size, category)

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchCategoryProducts(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	category, err := keys.ParseKey(drx.FindParam(r, "categoryKey"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	result, err := core.Context().FindProductsByCategory(page, size, category)

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchItemProducts(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	itemKey, err := keys.ParseKey(drx.FindParam(r, "itemKey"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Context().FindProductsByItem(page, size, itemKey)

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
