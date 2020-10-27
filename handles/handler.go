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
	//cars
	r.Handle("/cars", mw.Handler(http.HandlerFunc(GetCars))).Methods(http.MethodGet)
	r.Handle("/cars/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewCar))).Methods(http.MethodGet)
	r.Handle("/cars/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchCars))).Methods(http.MethodGet)
	r.Handle("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchCars))).Methods(http.MethodGet)
	r.Handle("/cars", mw.Handler(http.HandlerFunc(CreateCar))).Methods(http.MethodPost)
	r.Handle("/cars/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateCar))).Methods(http.MethodPut)

	//parts
	r.Handle("/parts", mw.Handler(http.HandlerFunc(GetParts))).Methods(http.MethodGet)
	r.Handle("/parts/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewPart))).Methods(http.MethodGet)
	r.Handle("/parts/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchParts))).Methods(http.MethodGet)
	r.Handle("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchParts))).Methods(http.MethodGet)
	r.Handle("/parts", mw.Handler(http.HandlerFunc(CreatePart))).Methods(http.MethodPost)
	r.Handle("/parts/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdatePart))).Methods(http.MethodPut)

	//properties
	r.Handle("/properties", mw.Handler(http.HandlerFunc(GetProperties))).Methods(http.MethodGet)
	r.Handle("/properties/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewProperty))).Methods(http.MethodGet)
	r.Handle("/properties/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchProperties))).Methods(http.MethodGet)
	r.Handle("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchProperties))).Methods(http.MethodGet)
	r.Handle("/properties", mw.Handler(http.HandlerFunc(CreateProperty))).Methods(http.MethodPost)
	r.Handle("/properties/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateProperty))).Methods(http.MethodPut)

	//services
	r.Handle("/services", mw.Handler(http.HandlerFunc(GetServices))).Methods(http.MethodGet)
	r.Handle("/services/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewService))).Methods(http.MethodGet)
	r.Handle("/services/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchServices))).Methods(http.MethodGet)
	r.Handle("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchServices))).Methods(http.MethodGet)
	r.Handle("/services", mw.Handler(http.HandlerFunc(CreateService))).Methods(http.MethodPost)
	r.Handle("/services/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateService))).Methods(http.MethodPut)

	//clothing
	r.Handle("/clothes", mw.Handler(http.HandlerFunc(GetClothing))).Methods(http.MethodGet)
	r.Handle("/clothes/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewClothing))).Methods(http.MethodGet)
	r.Handle("/clothes/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchClothing))).Methods(http.MethodGet)
	r.Handle("/clothes/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchClothing))).Methods(http.MethodGet)
	r.Handle("/clothes", mw.Handler(http.HandlerFunc(CreateClothing))).Methods(http.MethodPost)
	r.Handle("/clothes/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateClothing))).Methods(http.MethodPut)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "stock.cars.search", scrt)

	//if err != nil {
	//	panic(err)
	//}

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
