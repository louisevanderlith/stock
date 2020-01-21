package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Properties struct {
}

func (req *Properties) Get(c *gin.Context) {
	return http.StatusOK, nil
}
