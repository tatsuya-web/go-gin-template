package repository

import (
	clock "github.com/tatuya-web/go-gin-template/utils"
)

type Repository struct {
	Clocker clock.Clocker
}
