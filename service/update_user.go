package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

type UpdateUser struct {
	DB   infra.Execer
	Repo UserUpdater
}

func (uu *UpdateUser) UpdateUser(ctx context.Context, id model.UserID, email string) (*model.User, error) {
	if !auth.CheckOwn(ctx, id) {
		return nil, fmt.Errorf("権限のないユーザーです。")
	}

	u := &model.User{
		ID:    id,
		Email: email,
	}
	err := uu.Repo.UpdateUser(ctx, uu.DB, u)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return u, nil
}
