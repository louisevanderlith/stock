package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	//cars
	getC := kong.ResourceMiddleware(http.DefaultClient, "stock.cars.search", scrt, securityUrl, managerUrl, GetCars)
	r.HandleFunc("/cars", getC).Methods(http.MethodGet)

	viewC := kong.ResourceMiddleware(http.DefaultClient, "stock.cars.view", scrt, securityUrl, managerUrl, ViewCar)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", viewC).Methods(http.MethodGet)

	srchC := kong.ResourceMiddleware(http.DefaultClient, "stock.cars.search", scrt, securityUrl, managerUrl, SearchCars)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", srchC).Methods(http.MethodGet)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchC).Methods(http.MethodGet)

	createC := kong.ResourceMiddleware(http.DefaultClient, "stock.cars.create", scrt, securityUrl, managerUrl, CreateCar)
	r.HandleFunc("/cars", createC).Methods(http.MethodPost)

	updateC := kong.ResourceMiddleware(http.DefaultClient, "stock.cars.update", scrt, securityUrl, managerUrl, UpdateCar)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", updateC).Methods(http.MethodPut)

	//parts
	getPa := kong.ResourceMiddleware(http.DefaultClient, "stock.parts.search", scrt, securityUrl, managerUrl, GetParts)
	r.HandleFunc("/parts", getPa).Methods(http.MethodGet)

	viewPa := kong.ResourceMiddleware(http.DefaultClient, "stock.parts.view", scrt, securityUrl, managerUrl, ViewPart)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", viewPa).Methods(http.MethodGet)

	srchPa := kong.ResourceMiddleware(http.DefaultClient, "stock.parts.search", scrt, securityUrl, managerUrl, SearchParts)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", srchPa).Methods(http.MethodGet)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchPa).Methods(http.MethodGet)

	createPa := kong.ResourceMiddleware(http.DefaultClient, "stock.parts.create", scrt, securityUrl, managerUrl, CreatePart)
	r.HandleFunc("/parts", createPa).Methods(http.MethodPost)

	updatePa := kong.ResourceMiddleware(http.DefaultClient, "stock.parts.update", scrt, securityUrl, managerUrl, UpdatePart)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", updatePa).Methods(http.MethodPut)

	//properties
	getP := kong.ResourceMiddleware(http.DefaultClient, "stock.properties.search", scrt, securityUrl, managerUrl, GetProperties)
	r.HandleFunc("/properties", getP).Methods(http.MethodGet)

	viewP := kong.ResourceMiddleware(http.DefaultClient, "stock.properties.view", scrt, securityUrl, managerUrl, ViewPart)
	r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", viewP).Methods(http.MethodGet)

	srchP := kong.ResourceMiddleware(http.DefaultClient, "stock.properties.search", scrt, securityUrl, managerUrl, SearchParts)
	r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", srchP).Methods(http.MethodGet)
	r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchP).Methods(http.MethodGet)

	createP := kong.ResourceMiddleware(http.DefaultClient, "stock.properties.create", scrt, securityUrl, managerUrl, CreatePart)
	r.HandleFunc("/properties", createP).Methods(http.MethodPost)

	updateP := kong.ResourceMiddleware(http.DefaultClient, "stock.properties.update", scrt, securityUrl, managerUrl, UpdateProperty)
	r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", updateP).Methods(http.MethodPut)

	//services
	getS := kong.ResourceMiddleware(http.DefaultClient, "stock.services.search", scrt, securityUrl, managerUrl, GetServices)
	r.HandleFunc("/services", getS).Methods(http.MethodGet)

	viewS := kong.ResourceMiddleware(http.DefaultClient, "stock.services.view", scrt, securityUrl, managerUrl, ViewService)
	r.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", viewS).Methods(http.MethodGet)

	srchS := kong.ResourceMiddleware(http.DefaultClient, "stock.services.search", scrt, securityUrl, managerUrl, SearchServices)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", srchS).Methods(http.MethodGet)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchS).Methods(http.MethodGet)

	createS := kong.ResourceMiddleware(http.DefaultClient, "stock.services.create", scrt, securityUrl, managerUrl, CreateService)
	r.HandleFunc("/services", createS).Methods(http.MethodPost)

	updateS := kong.ResourceMiddleware(http.DefaultClient, "stock.services.update", scrt, securityUrl, managerUrl, UpdateService)
	r.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", updateS).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "stock.cars.search", scrt)

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
