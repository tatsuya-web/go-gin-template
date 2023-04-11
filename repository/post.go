package repository

import (
	"context"

	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/domain/model"
)

func (r *Repository) AddPost(ctx context.Context, db Execer, p *model.Post) error {
	p.CreatedAt = r.Clocker.Now()
	p.UpdatedAt = r.Clocker.Now()

	sql := `INSERT INTO posts (
		title, content, user_id, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, p.Title, p.Content, p.UserID, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = model.PostID(id)
	return nil
}

func (r *Repository) UpdatePost(ctx context.Context, db Execer, p *model.Post) error {
	p.UpdatedAt = r.Clocker.Now()

	sql := `UPDATE posts
								SET title = ?,
								content = ?,
								updated_at = ?
								WHERE id = ?`

	_, err := db.ExecContext(
		ctx, sql, p.Title, p.Content, p.UpdatedAt, p.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeletePost(ctx context.Context, db Execer, p *model.Post) error {
	sql := `DELETE FROM posts WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, p.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListPosts(ctx context.Context, db Queryer, id model.UserID) (model.Posts, error) {
	posts := model.Posts{}

	sql := `SELECT
								id, title, content, user_id, created_at, updated_at
								FROM posts
								WHERE user_id = ?`
	if err := db.SelectContext(ctx, &posts, sql, id); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *Repository) IsOwnPost(ctx context.Context, db Queryer, id model.PostID) bool {
	ownID, ok := auth.GetUserID(ctx)
	if !ok {
		return false
	}

	var userID int64

	sql := `SELECT
								user_id
								FROM posts
								WHERE id = ?`
	db.QueryRowxContext(ctx, sql, id).Scan(&userID)

	return ownID == model.UserID(userID)
}
