package handler

import (
	"context"
	"net/http"

	"github.com/tatuya-web/go-gin-template/domain/model"
)

type AddPostService interface {
	AddPost(ctx context.Context, title string, content string) (*model.Post, error)
}

type UpdatePostService interface {
	UpdatePost(ctx context.Context, id model.PostID, title string, content string) (*model.Post, error)
}

type DeletePostService interface {
	DeletePost(ctx context.Context, id model.PostID) error
}

type ListPostService interface {
	ListPosts(ctx context.Context) (model.Posts, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, email, password, role string) (*model.User, error)
}

type SigninService interface {
	Signin(ctx context.Context, email, pw string) (string, error)
}

type SignoutService interface {
	Signout(ctx context.Context, r *http.Request) error
}

type OwnUserService interface {
	OwnUser(ctx context.Context) (*model.User, error)
}

type UpdateUserService interface {
	UpdateUser(ctx context.Context, id model.UserID, email string) (*model.User, error)
}

type DeleteUserService interface {
	DeleteUser(ctx context.Context, r *http.Request, id model.UserID) error
}
