package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tatuya-web/go-gin-template/config"
	"github.com/tatuya-web/go-gin-template/repository"
	"github.com/tatuya-web/go-gin-template/router"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if cfg.Env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	// ミドルウェアの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	//DB初期化
	db, cleanup, err := repository.New(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	//ルーティング初期化
	if err := router.SetRouting(ctx, db, r, cfg); err != nil {
		return nil
	}
	if err := router.SetAuthRouting(ctx, db, r, cfg); err != nil {
		return nil
	}

	//サーバー起動
	log.Printf("Listening and serving HTTP on :%v", cfg.Port)
	server := NewServer(r, fmt.Sprintf(":%d", cfg.Port))
	return server.Run(ctx)
}
