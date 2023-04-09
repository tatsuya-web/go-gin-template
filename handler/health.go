package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct{}

func NewHealthhandler() *Health {
	return &Health{}
}

func (ap *Health) ServeHTTP(ctx *gin.Context) {
	APIResponse(ctx, http.StatusOK, "ok", nil)
}
