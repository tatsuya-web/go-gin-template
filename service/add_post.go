package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

type AddPost struct {
	DB   infra.Execer
	Repo PostAdder
}

func (ap *AddPost) AddPost(ctx context.Context, title string, content string) (*model.Post, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("ユーザーが見つかりません。")
	}

	p := &model.Post{
		Title:   title,
		Content: content,
		UserID:  id,
	}
	err := ap.Repo.AddPost(ctx, ap.DB, p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}
