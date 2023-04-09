package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Signout struct {
	Service SignoutService
}

func NewSignoutHandler(s SignoutService) *Signout {
	return &Signout{Service: s}
}

func (s *Signout) ServeHTTP(ctx *gin.Context) {
	if err := s.Service.Signout(ctx, ctx.Request); err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "サインアウトエラー", err.Error())
		return
	}

	APIResponse(ctx, http.StatusOK, "サインアウトしました。", nil)
}
