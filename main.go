package main

import (
	"context"

	"github.com/KhachikAstoyan/go-api-template/core"
	"github.com/KhachikAstoyan/go-api-template/db"
	"github.com/KhachikAstoyan/go-api-template/db/repository"
	"github.com/KhachikAstoyan/go-api-template/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := core.LoadDotenvConfig()
	conn := db.CreateConnection(ctx, config.DB_URL)
	repo := repository.New(conn)

	app := &core.App{
		Config: config,
		DB:     repo,
	}

	http.StartServer(app)
}
