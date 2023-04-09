package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

type OwnUser struct {
	DB   infra.Queryer
	Repo OwnGetter
}

func (uo *OwnUser) OwnUser(ctx context.Context) (*model.User, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	user, err := uo.Repo.GetOwn(ctx, uo.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return user, nil
}
