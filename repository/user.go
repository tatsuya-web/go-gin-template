package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/infra"
)

func (r *Repository) RegisterUser(ctx context.Context, db infra.Execer, u *model.User) error {
	u.CreatedAt = r.Clocker.Now()
	u.UpdatedAt = r.Clocker.Now()

	sql := `INSERT INTO users (
	email, password, role, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(ctx, sql, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == ErrCodeMySQLDuplicateEntry {
			return fmt.Errorf("cannot create same email user: %w", ErrAlreadyEntry)
		}
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = model.UserID(id)
	return nil
}

func (r *Repository) GetUser(ctx context.Context, db infra.Queryer, email string) (*model.User, error) {
	u := &model.User{}
	sql := `SELECT
			id, email, password, role, created_at, updated_at 
			FROM users WHERE email = ?`
	if err := db.GetContext(ctx, u, sql, email); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) GetOwn(ctx context.Context, db infra.Queryer, id model.UserID) (*model.User, error) {
	u := &model.User{}

	sql := `SELECT
								id, email, role, created_at, updated_at
								FROM users
								WHERE id = ?`
	if err := db.GetContext(ctx, u, sql, id); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(ctx context.Context, db infra.Execer, u *model.User) error {
	u.UpdatedAt = r.Clocker.Now()

	sql := `UPDATE users
								SET email = ?,
								updated_at = ?
								WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, u.Email, u.UpdatedAt, u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, db infra.Execer, u *model.User) error {
	sql := `DELETE FROM users WHERE id = ?`
	_, err := db.ExecContext(
		ctx, sql, u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
