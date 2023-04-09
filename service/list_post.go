package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

type ListPost struct {
	DB   infra.Queryer
	Repo PostLister
}

func (lp *ListPost) ListPosts(ctx context.Context) (model.Posts, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	posts, err := lp.Repo.ListPosts(ctx, lp.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return posts, nil
}
