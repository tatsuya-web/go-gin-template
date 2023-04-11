package service

import (
	"context"
	"net/http"

	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/repository"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . PostAdder PostUpdater PostLister UserRegister UserGetter TokenGenerator TokenDeleter OwnGetter UserDeleter
type PostAdder interface {
	AddPost(ctx context.Context, db repository.Execer, p *model.Post) error
}

type PostUpdater interface {
	IsOwnPost(ctx context.Context, db repository.Queryer, id model.PostID) bool
	UpdatePost(ctx context.Context, db repository.Execer, p *model.Post) error
}

type PostDeleter interface {
	IsOwnPost(ctx context.Context, db repository.Queryer, id model.PostID) bool
	DeletePost(ctx context.Context, db repository.Execer, p *model.Post) error
}

type PostLister interface {
	ListPosts(ctx context.Context, db repository.Queryer, id model.UserID) (model.Posts, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db repository.Execer, u *model.User) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db repository.Queryer, email string) (*model.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u model.User) ([]byte, error)
}

type TokenDeleter interface {
	DeleteToken(ctx context.Context, r *http.Request, id model.UserID) error
}

type OwnGetter interface {
	GetOwn(ctx context.Context, db repository.Queryer, id model.UserID) (*model.User, error)
}

type UserUpdater interface {
	UpdateUser(ctx context.Context, db repository.Execer, u *model.User) error
}

type UserDeleter interface {
	DeleteUser(ctx context.Context, db repository.Execer, p *model.User) error
}
