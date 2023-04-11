package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/config"
	"github.com/tatuya-web/go-gin-template/handler"
	"github.com/tatuya-web/go-gin-template/repository"
	"github.com/tatuya-web/go-gin-template/service"
	"github.com/tatuya-web/go-gin-template/util"
)

func SetRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {

	clocker := util.RealClocker{}
	rep := repository.Repository{Clocker: clocker}
	validate := validator.New()
	rcli, err := repository.NewKVS(ctx, cfg)
	if err != nil {
		return err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return err
	}

	//ルートグループ作成
	rg := router.Group("/api/v1")

	//ヘルスチェック
	healthHandler := handler.NewHealthhandler()
	router.GET("/health", healthHandler.ServeHTTP)

	//ユーザー登録
	registerUserHandler := handler.NewRegisterUserHandler(
		&service.RegisterUser{DB: db, Repo: &rep},
		validate,
	)
	rg.POST("register", registerUserHandler.ServeHTTP)

	//サインイン
	signinUserHandler := handler.NewSigninHandler(
		&service.Signin{DB: db, Repo: &rep, TokenGenerator: jwter},
		validate,
	)
	rg.POST("signin", signinUserHandler.ServeHTTP)

	return nil
}
