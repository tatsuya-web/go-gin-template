package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tatuya-web/go-gin-template/auth"
)

type Signout struct {
	TokenDeleter TokenDeleter
}

func (s *Signout) Signout(ctx context.Context, r *http.Request) error {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("ユーザーが見つかりません。")
	}

	if err := s.TokenDeleter.DeleteToken(ctx, r, id); err != nil {
		return fmt.Errorf("failed to token: %w", err)
	}

	return nil
}
