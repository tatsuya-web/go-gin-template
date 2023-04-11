package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   repository.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(
	ctx context.Context,
	email,
	password,
	role string) (*model.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	u := &model.User{
		Email:    email,
		Password: string(pw),
		Role:     role,
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, u); err != nil {
		return nil, fmt.Errorf("faild to register: %w", err)
	}
	return u, nil
}
