package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

type DeleteUser struct {
	DB           infra.Execer
	Repo         UserDeleter
	TokenDeleter TokenDeleter
}

func (du *DeleteUser) DeleteUser(ctx context.Context, r *http.Request, id model.UserID) error {
	if !auth.CheckOwn(ctx, id) {
		return fmt.Errorf("権限のないユーザーです。")
	}

	if err := du.TokenDeleter.DeleteToken(ctx, r, id); err != nil {
		return fmt.Errorf("failed to token: %w", err)
	}

	u := &model.User{
		ID: id,
	}
	err := du.Repo.DeleteUser(ctx, du.DB, u)
	if err != nil {
		return fmt.Errorf("failed to post: %w", err)
	}

	return nil
}
