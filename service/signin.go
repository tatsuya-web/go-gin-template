package service

import (
	"context"
	"fmt"

	"github.com/tatuya-web/go-gin-template/repository"
)

type Signin struct {
	DB             repository.Queryer
	Repo           UserGetter
	TokenGenerator TokenGenerator
}

func (s *Signin) Signin(ctx context.Context, email, pw string) (string, error) {
	u, err := s.Repo.GetUser(ctx, s.DB, email)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w", err)
	}

	if err := u.ComparePassword(pw); err != nil {
		return "", fmt.Errorf("warning password: %w", err)
	}

	jwt, err := s.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}
	return string(jwt), nil
}
