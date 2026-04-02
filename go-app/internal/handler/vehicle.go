package handler

import (
	"go-app/internal/dto"
	"go-app/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	repo *repository.Repo
	db   *gorm.DB
}

func New(r *repository.Repo, db *gorm.DB) *Handler {
	return &Handler{
		repo: r,
		db:   db,
	}
}

func (h *Handler) GetLast(c *gin.Context) {
	id := c.Param("id")

	data, err := h.repo.GetLast(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if data == nil || data.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "vehicle not found"})
		return
	}

	resp := dto.VehicleLocationResponse{
		VehicleID: data.VehicleID,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Timestamp: data.Timestamp.Unix(),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHistory(c *gin.Context) {
	id := c.Param("id")

	startStr := c.Query("start")
	endStr := c.Query("end")

	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start and end required"})
		return
	}

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start"})
		return
	}

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end"})
		return
	}

	if start > end {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start must be less than end"})
		return
	}

	data, err := h.repo.GetHistory(id, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []dto.VehicleLocationResponse

	for _, d := range data {
		result = append(result, dto.VehicleLocationResponse{
			VehicleID: d.VehicleID,
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
			Timestamp: d.Timestamp.Unix(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Health(c *gin.Context) {

	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db error"})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "db down"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "go-app",
	})
}
