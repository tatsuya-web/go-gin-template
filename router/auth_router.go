package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tatuya-web/go-gin-template/auth"
	"github.com/tatuya-web/go-gin-template/config"
	"github.com/tatuya-web/go-gin-template/handler"
	"github.com/tatuya-web/go-gin-template/middleware"
	"github.com/tatuya-web/go-gin-template/repository"
	"github.com/tatuya-web/go-gin-template/service"
	"github.com/tatuya-web/go-gin-template/util"
)

func SetAuthRouting(ctx context.Context, db *sqlx.DB, router *gin.Engine, cfg *config.Config) error {
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
	rg := router.Group("/api/v1").
		Use(middleware.AuthMiddleware(jwter))

	//POST登録
	addPostHandler := handler.NewAddPosthandler(
		&service.AddPost{DB: db, Repo: &rep},
		validate,
	)
	rg.POST("posts", addPostHandler.ServeHTTP)

	//POST更新
	updatePostHandler := handler.NewUpdatePosthandler(
		&service.UpdatePost{DBExec: db, DBQuery: db, Repo: &rep},
		validate,
	)
	rg.PATCH("posts", updatePostHandler.ServeHTTP)

	//POST削除
	deletePostHandler := handler.NewDeletePostHandler(
		&service.DeletePost{DBExec: db, DBQuery: db, Repo: &rep},
		validate,
	)
	rg.DELETE("posts", deletePostHandler.ServeHTTP)

	//POST一覧 (roleがユーザーの場合は自身にひもずくPOSTだけがし取得される)
	listPostHandler := handler.NewListPostHandler(
		&service.ListPost{DB: db, Repo: &rep},
	)
	rg.GET("posts", listPostHandler.ServeHTTP)

	//プロフィール
	ownUserHandler := handler.NewOwnUserHandler(
		&service.OwnUser{DB: db, Repo: &rep},
	)
	rg.GET("user", ownUserHandler.ServeHTTP)

	//USER更新
	updateUserHandler := handler.NewUpdateUserHandler(
		&service.UpdateUser{DB: db, Repo: &rep},
		validate,
	)
	rg.PATCH("user", updateUserHandler.ServeHTTP)

	//USER削除
	deleteUserHandler := handler.NewDeleteUserHandler(
		&service.DeleteUser{DB: db, Repo: &rep, TokenDeleter: jwter},
		validate,
	)
	rg.DELETE("user", deleteUserHandler.ServeHTTP)

	//サインアウト
	signoutHandler := handler.NewSignoutHandler(
		&service.Signout{TokenDeleter: jwter},
	)
	rg.POST("signout", signoutHandler.ServeHTTP)

	return nil
}
