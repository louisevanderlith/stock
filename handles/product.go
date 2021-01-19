package handles

import (
	"encoding/base64"
	"encoding/json"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/core"
	"log"
	"net/http"
)

func ViewProduct(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().GetProduct(key)

	if err != nil {
		log.Println("Get Category Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(rec.GetValue()))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	hsh, err := base64.URLEncoding.DecodeString(drx.FindParam(r, "hash"))

	if err != nil {
		log.Println("Hash Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	in := core.Product{}
	err = json.Unmarshal(hsh, &in)

	if err != nil {
		log.Println("Hash Bind Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	result, err := core.Context().SearchProducts(page, size, in)

	if err != nil {
		log.Println("Search Products Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	obj := core.Product{}

	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Context().CreateProduct(obj)

	if err != nil {
		log.Println("Create Product Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Product{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateProduct(key, obj)

	if err != nil {
		log.Println("Update Product Error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
