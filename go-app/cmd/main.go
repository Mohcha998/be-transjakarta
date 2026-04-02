package main

import (
	"go-app/internal/config"
	"go-app/internal/database"
	"go-app/internal/handler"
	"go-app/internal/repository"
	"go-app/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	db := database.Init(cfg.DB)

	repo := repository.New(db)
	h := handler.New(repo, db)

	r := gin.Default()

	routes.Register(r, h)

	r.Run(":8080")
}
