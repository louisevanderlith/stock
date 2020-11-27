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

	r.Handle("/categories", mw.Handler(http.HandlerFunc(GetClientCategories))).Methods(http.MethodGet)
	r.Handle("/categories/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchClientCategories))).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(GetCategories))).Methods(http.MethodGet)
	r.Handle("/info/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchCategories))).Methods(http.MethodGet)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewCategory))).Methods(http.MethodGet)
	r.Handle("/info", mw.Handler(http.HandlerFunc(CreateCategory))).Methods(http.MethodPost)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateCategory))).Methods(http.MethodPut)

	//stock
	r.Handle("/{category:[a-zA-Z]+}/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewStock))).Methods(http.MethodGet)
	r.Handle("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchStock))).Methods(http.MethodGet)
	r.Handle("/{category:[a-zA-Z]+}/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchStock))).Methods(http.MethodGet)
	r.Handle("/{category:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(CreateStock))).Methods(http.MethodPost)
	r.Handle("/{category:[a-zA-Z]+}", mw.Handler(http.HandlerFunc(UpdateStock))).Methods(http.MethodPut)

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
