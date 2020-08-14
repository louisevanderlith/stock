package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func GetProperties(w http.ResponseWriter, r *http.Request) {
	err := mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
