package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-gin-template/domain/model"
)

type UpdateUser struct {
	Service   UpdateUserService
	Validator *validator.Validate
}

func NewUpdateUserHandler(uu UpdateUserService, v *validator.Validate) *UpdateUser {
	return &UpdateUser{Service: uu, Validator: v}
}

func (uu *UpdateUser) ServeHTTP(ctx *gin.Context) {
	var u struct {
		ID    int64  `json:"id" validate:"required"`
		Email string `json:"email" validate:"required,email,max=255"`
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&u); err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	err := uu.Validator.Struct(u)
	if err != nil {
		ErrResponse(ctx, http.StatusBadRequest, "faild to user", err.Error())
		return
	}

	user, err := uu.Service.UpdateUser(ctx, model.UserID(u.ID), u.Email)
	if err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	rsp := struct {
		ID    model.UserID `json:"id"`
		Email string       `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	}
	APIResponse(ctx, http.StatusOK, "postを登録しました。", rsp)
}
