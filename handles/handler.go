package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	//cars
	getC := kong.ResourceMiddleware("stock.cars.search", scrt, secureUrl, GetCars)
	r.HandleFunc("/cars", getC).Methods(http.MethodGet)

	viewC := kong.ResourceMiddleware("stock.cars.view", scrt, secureUrl, ViewCars)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", viewC).Methods(http.MethodGet)

	srchC := kong.ResourceMiddleware("stock.cars.search", scrt, secureUrl, SearchCars)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", srchC).Methods(http.MethodGet)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchC).Methods(http.MethodGet)

	createC := kong.ResourceMiddleware("stock.cars.create", scrt, secureUrl, CreateCars)
	r.HandleFunc("/cars", createC).Methods(http.MethodPost)

	updateC := kong.ResourceMiddleware("stock.cars.update", scrt, secureUrl, UpdateCars)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", updateC).Methods(http.MethodPut)

	//parts
	getPa := kong.ResourceMiddleware("stock.parts.search", scrt, secureUrl, GetParts)
	r.HandleFunc("/parts", getPa).Methods(http.MethodGet)

	viewPa := kong.ResourceMiddleware("stock.parts.view", scrt, secureUrl, ViewParts)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", viewPa).Methods(http.MethodGet)

	srchPa := kong.ResourceMiddleware("stock.parts.search", scrt, secureUrl, SearchParts)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", srchPa).Methods(http.MethodGet)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchPa).Methods(http.MethodGet)

	createPa := kong.ResourceMiddleware("stock.parts.create", scrt, secureUrl, CreateParts)
	r.HandleFunc("/parts", createPa).Methods(http.MethodPost)

	updatePa := kong.ResourceMiddleware("stock.parts.update", scrt, secureUrl, UpdateParts)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", updatePa).Methods(http.MethodPut)

	//properties
	getP := kong.ResourceMiddleware("stock.properties.search", scrt, secureUrl, GetProperties)
	r.HandleFunc("/properties", getP).Methods(http.MethodGet)
	/*
		viewP := kong.ResourceMiddleware("stock.properties.view", scrt, secureUrl, ViewParts)
		r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", viewP).Methods(http.MethodGet)

		srchP := kong.ResourceMiddleware("stock.properties.search", scrt, secureUrl, SearchParts)
		r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", srchP).Methods(http.MethodGet)
		r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchP).Methods(http.MethodGet)

		createP := kong.ResourceMiddleware("stock.properties.create", scrt, secureUrl, CreateParts)
		r.HandleFunc("/properties", createP).Methods(http.MethodPost)

		updateP := kong.ResourceMiddleware("stock.properties.update", scrt, secureUrl, UpdateProperties)
		r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", updateP).Methods(http.MethodPut)
	*/
	//services
	getS := kong.ResourceMiddleware("stock.services.search", scrt, secureUrl, GetServices)
	r.HandleFunc("/services", getS).Methods(http.MethodGet)

	viewS := kong.ResourceMiddleware("stock.services.view", scrt, secureUrl, ViewServices)
	r.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", viewS).Methods(http.MethodGet)

	srchS := kong.ResourceMiddleware("stock.services.search", scrt, secureUrl, SearchServices)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", srchS).Methods(http.MethodGet)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchS).Methods(http.MethodGet)

	createS := kong.ResourceMiddleware("stock.services.create", scrt, secureUrl, CreateServices)
	r.HandleFunc("/services", createS).Methods(http.MethodPost)

	updateS := kong.ResourceMiddleware("stock.services.update", scrt, secureUrl, UpdateServices)
	r.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", updateS).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "stock.cars.search", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
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
