package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OwnUser struct {
	Service OwnUserService
}

func NewOwnUserHandler(uo OwnUserService) *OwnUser {
	return &OwnUser{Service: uo}
}

func (uo *OwnUser) ServeHTTP(ctx *gin.Context) {
	user, err := uo.Service.OwnUser(ctx)
	if err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to user", err.Error())
		return
	}

	APIResponse(ctx, http.StatusOK, "ユーザー情報", user)
}
