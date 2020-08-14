package main

import (
	"flag"
	"github.com/louisevanderlith/stock/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/stock/core"
)

func main() {
	security := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	manager := flag.String("manager", "http://localhost:8097", "Manager Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8101",
		Handler:      handles.SetupRoutes(*srcSecrt, *security, *manager),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
