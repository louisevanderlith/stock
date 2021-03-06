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

func GetCategories(w http.ResponseWriter, r *http.Request) {
	result, err := core.Context().ListCategories(1, 10)

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

func SearchCategories(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	result, err := core.Context().ListCategories(page, size)

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

func SearchHashCategories(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	hsh, err := base64.URLEncoding.DecodeString(drx.FindParam(r, "hash"))

	if err != nil {
		log.Println("Hash Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	in := core.Category{}
	err = json.Unmarshal(hsh, &in)

	if err != nil {
		log.Println("Hash Bind Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	result, err := core.Context().SearchCategories(page, size, in)

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

func ViewCategory(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().GetCategory(key)

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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	obj := core.Category{}

	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Context().CreateCategory(obj)

	if err != nil {
		log.Println("Create Category Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Category{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateCategory(key, obj)

	if err != nil {
		log.Println("Update Category Error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
