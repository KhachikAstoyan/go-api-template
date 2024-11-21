package core

import (
	"github.com/KhachikAstoyan/go-api-template/db/repository"
)

type App struct {
	Config Config
	DB     *repository.Queries
}
