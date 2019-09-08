package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Properties struct {
}

func (req *Properties) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusOK, nil
}
