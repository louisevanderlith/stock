package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)

	r.HandleFunc("/categories", GetClientCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/name/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", SearchCategoryNameProducts).Methods(http.MethodGet)
	r.HandleFunc("/categories/name/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchCategoryNameProducts).Methods(http.MethodGet)
	r.HandleFunc("/categories/{pagesize:[A-Z][0-9]+}", SearchClientCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchClientCategories).Methods(http.MethodGet)

	r.HandleFunc("/info", GetCategories).Methods(http.MethodGet)
	r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}", SearchCategories).Methods(http.MethodGet)
	r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchCategories).Methods(http.MethodGet)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", ViewCategory).Methods(http.MethodGet)
	r.HandleFunc("/info", CreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", UpdateCategory).Methods(http.MethodPut)

	/*
		FindCategoriesByClient(page, size int, clientId string) (records.Page, error)
		FindCategoriesByOwner(page, size int, owner hsk.Key) (records.Page, error)
		FindProductsByItem(page, size int, itemKey hsk.Key) (records.Page, error)
		FindProductsByCategory(page, size int, categoryKey hsk.Key) (records.Page, error)
		FindProductsByCategoryName(page, size int, category string) (records.Page, error)
	*/

	//products
	r.HandleFunc("/products/{key:[0-9]+\\x60[0-9]+}", ViewProduct).Methods(http.MethodGet)
	r.HandleFunc("/products", CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{key:[0-9]+\\x60[0-9]+}", UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/products/item/{itemKey:[0-9]+\\x60[0-9]+}/{pagesize:[A-Z][0-9]+}", SearchItemProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/category/{categoryKey:[0-9]+\\x60[0-9]+}/{pagesize:[A-Z][0-9]+}", SearchCategoryProducts).Methods(http.MethodGet)
	//search
	r.HandleFunc("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", SearchCategoryNameProducts).Methods(http.MethodGet)
	r.HandleFunc("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchCategoryNameProducts).Methods(http.MethodGet)

	r.Use(mw.Handler)
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
