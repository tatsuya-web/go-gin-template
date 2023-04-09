package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
	"github.com/tatuya-web/go-gin-template/repository"
)

type UpdatePost struct {
	DBExec  infra.Execer
	DBQuery infra.Queryer
	Repo    PostUpdater
}

func (up *UpdatePost) UpdatePost(ctx context.Context, id model.PostID, title string, content string) (*model.Post, error) {
	if !up.Repo.IsOwnPost(ctx, up.DBQuery, id) {
		return nil, fmt.Errorf("failed to post: %w", repository.ErrUnauthorizedUser)
	}

	p := &model.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}
	err := up.Repo.UpdatePost(ctx, up.DBExec, p)
	if err != nil {
		return nil, fmt.Errorf("failed to post: %w", err)
	}

	return p, nil
}
