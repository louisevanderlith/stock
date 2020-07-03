package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func GetProperties(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	ctx.Serve(http.StatusOK, mix.JSON(nil))
}
