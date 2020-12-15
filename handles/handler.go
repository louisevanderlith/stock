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
	r.HandleFunc("/categories/{pagesize:[A-Z][0-9]+}", SearchClientCategories).Methods(http.MethodGet)

	r.HandleFunc("/info", GetCategories).Methods(http.MethodGet)
	r.HandleFunc("/info/{pagesize:[A-Z][0-9]+}", SearchCategories).Methods(http.MethodGet)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", ViewCategory).Methods(http.MethodGet)
	r.HandleFunc("/info", CreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/info/{key:[0-9]+\\x60[0-9]+}", UpdateCategory).Methods(http.MethodPut)

	//stock
	r.HandleFunc("/{category:[a-zA-Z]+}/{key:[0-9]+\\x60[0-9]+}", ViewStock).Methods(http.MethodGet)
	r.HandleFunc("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", SearchStock).Methods(http.MethodGet)
	r.HandleFunc("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchStock).Methods(http.MethodGet)
	r.HandleFunc("/{category:[a-zA-Z]+}", CreateStock).Methods(http.MethodPost)
	r.HandleFunc("/{category:[a-zA-Z]+}", UpdateStock).Methods(http.MethodPut)

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
