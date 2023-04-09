package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AddPost struct {
	Service   AddPostService
	Validator *validator.Validate
}

func NewAddPosthandler(ap AddPostService, v *validator.Validate) *AddPost {
	return &AddPost{Service: ap, Validator: v}
}

func (ap *AddPost) ServeHTTP(ctx *gin.Context) {
	var p struct {
		Title   string `json:"title" validate:"required,max=255"`
		Content string `json:"content" validate:"required,max=255"`
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&p); err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	err := ap.Validator.Struct(p)
	if err != nil {
		ErrResponse(ctx, http.StatusBadRequest, "faild to post", err.Error())
		return
	}

	post, err := ap.Service.AddPost(ctx, p.Title, p.Content)
	if err != nil {
		ErrResponse(ctx, http.StatusInternalServerError, "faild to post", err.Error())
		return
	}

	APIResponse(ctx, http.StatusOK, "postを登録しました。", gin.H{"id": post.ID})
}
