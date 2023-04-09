package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-gin-template/domain/model"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func NewRegisterUserHandler(ru RegisterUserService, v *validator.Validate) *RegisterUser {
	return &RegisterUser{Service: ru, Validator: v}
}

func (ru *RegisterUser) ServeHTTP(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" validate:"required,email,max=255"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	if err := ru.Validator.Struct(input); err != nil {
		ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	u, err := ru.Service.RegisterUser(ctx, input.Email, input.Password, input.Role)

	if err != nil {
		ErrResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())
		return
	}
	rsp := struct {
		ID model.UserID `json:"id"`
	}{ID: u.ID}

	APIResponse(ctx, http.StatusCreated, "本登録が完了しました。", rsp)
}
