package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tatuya-web/go-gin-template/domain/model"
)

type UpdatePost struct {
	Service   UpdatePostService
	Validator *validator.Validate
}

func NewUpdatePosthandler(up UpdatePostService, v *validator.Validate) *UpdatePost {
	return &UpdatePost{Service: up, Validator: v}
}

func (up *UpdatePost) ServeHTTP(ctx *gin.Context) {
	var p struct {
		ID      int64  `json:"id" validate:"required"`
		Title   string `json:"title" validate:"required,max=255"`
		Content string `json:"content" validate:"required,max=255"`
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&p); err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	err := up.Validator.Struct(p)
	if err != nil {
		ErrResponse(ctx, http.StatusBadRequest, "faild to post", err.Error())
		return
	}

	post, err := up.Service.UpdatePost(ctx, model.PostID(p.ID), p.Title, p.Content)
	if err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	rsp := struct {
		ID      model.PostID `json:"id"`
		Title   string       `json:"title"`
		Content string       `json:"content"`
	}{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
	APIResponse(ctx, http.StatusOK, "postを更新しました。", rsp)
}
