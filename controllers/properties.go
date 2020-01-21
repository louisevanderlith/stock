package controllers

import (
	"net/http"
)

type Properties struct {
}

func (req *Properties) Get(c *gin.Context) {
	return http.StatusOK, nil
}
