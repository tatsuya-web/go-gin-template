package auth

import (
	"context"

	"github.com/tatuya-web/go-gin-template/domain/model"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID model.UserID) error
	Load(ctx context.Context, key string) (model.UserID, error)
	Delete(ctx context.Context, key string) error
}
