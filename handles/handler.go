package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := kong.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	//cars
	getC := ins.Middleware("stock.cars.search", scrt, GetCars)
	r.HandleFunc("/cars", getC).Methods(http.MethodGet)

	viewC := ins.Middleware("stock.cars.view", scrt, ViewCar)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", viewC).Methods(http.MethodGet)

	srchC := ins.Middleware("stock.cars.search", scrt, SearchCars)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}", srchC).Methods(http.MethodGet)
	r.HandleFunc("/cars/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchC).Methods(http.MethodGet)

	createC := ins.Middleware("stock.cars.create", scrt, CreateCar)
	r.HandleFunc("/cars", createC).Methods(http.MethodPost)

	updateC := ins.Middleware("stock.cars.update", scrt, UpdateCar)
	r.HandleFunc("/cars/{key:[0-9]+\\x60[0-9]+}", updateC).Methods(http.MethodPut)

	//parts
	getPa := ins.Middleware("stock.parts.search", scrt, GetParts)
	r.HandleFunc("/parts", getPa).Methods(http.MethodGet)

	viewPa := ins.Middleware("stock.parts.view", scrt, ViewPart)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", viewPa).Methods(http.MethodGet)

	srchPa := ins.Middleware("stock.parts.search", scrt, SearchParts)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}", srchPa).Methods(http.MethodGet)
	r.HandleFunc("/parts/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchPa).Methods(http.MethodGet)

	createPa := ins.Middleware("stock.parts.create", scrt, CreatePart)
	r.HandleFunc("/parts", createPa).Methods(http.MethodPost)

	updatePa := ins.Middleware("stock.parts.update", scrt, UpdatePart)
	r.HandleFunc("/parts/{key:[0-9]+\\x60[0-9]+}", updatePa).Methods(http.MethodPut)

	//properties
	getP := ins.Middleware("stock.properties.search", scrt, GetProperties)
	r.HandleFunc("/properties", getP).Methods(http.MethodGet)

	viewP := ins.Middleware("stock.properties.view", scrt, ViewPart)
	r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", viewP).Methods(http.MethodGet)

	srchP := ins.Middleware("stock.properties.search", scrt, SearchParts)
	r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}", srchP).Methods(http.MethodGet)
	r.HandleFunc("/properties/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchP).Methods(http.MethodGet)

	createP := ins.Middleware("stock.properties.create", scrt, CreatePart)
	r.HandleFunc("/properties", createP).Methods(http.MethodPost)

	updateP := ins.Middleware("stock.properties.update", scrt, UpdateProperty)
	r.HandleFunc("/properties/{key:[0-9]+\\x60[0-9]+}", updateP).Methods(http.MethodPut)

	//services
	getS := ins.Middleware("stock.services.search", scrt, GetServices)
	r.HandleFunc("/services", getS).Methods(http.MethodGet)

	viewS := ins.Middleware("stock.services.view", scrt, ViewService)
	r.HandleFunc("/services/{key:[0-9]+\\x60[0-9]+}", viewS).Methods(http.MethodGet)

	srchS := ins.Middleware("stock.services.search", scrt, SearchServices)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}", srchS).Methods(http.MethodGet)
	r.HandleFunc("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srchS).Methods(http.MethodGet)

	createS := ins.Middleware("stock.services.create", scrt, CreateService)
	r.HandleFunc("/services", createS).Methods(http.MethodPost)

	updateS := ins.Middleware("stock.services.update", scrt, UpdateService)
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
