package routes

import (
	"go-app/internal/handler"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handler.Handler) {

	r.GET("/health", h.Health)
	r.GET("/vehicles/:id/location", h.GetLast)
	r.GET("/vehicles/:id/history", h.GetHistory)
}
