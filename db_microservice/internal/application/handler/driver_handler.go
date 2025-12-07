package handler

import (
	"bitaksi_burakcanheyal/db_microservice/internal/domain/dto"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type DriverHandler struct {
	service *service.DriverService
}

func NewDriverHandler(s *service.DriverService) *DriverHandler {
	return &DriverHandler{service: s}
}

// ──────────────────────────── CREATE ────────────────────────────
func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var req dto.CreateDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request body",
			"detail": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id, err := h.service.CreateDriver(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateDriverResponse{ID: id})
}

// ──────────────────────────── UPDATE ────────────────────────────
func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	var req dto.UpdateDriverRequest
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing driver id"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request body",
			"detail": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.UpdateDriver(ctx, id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// ──────────────────────────── LIST ────────────────────────────
func (h *DriverHandler) ListDrivers(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("pageSize")

	page, err1 := strconv.Atoi(pageStr)
	pageSize, err2 := strconv.Atoi(sizeStr)

	if err1 != nil || page < 1 {
		page = 1
	}
	if err2 != nil || pageSize < 1 {
		pageSize = 20
	}

	drivers, err := h.service.ListDrivers(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drivers)
}

// ──────────────────────────── NEARBY ────────────────────────────
func (h *DriverHandler) GetNearbyDrivers(c *gin.Context) {
	var req dto.NearbyDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request body",
			"detail": err.Error(),
		})
		return
	}

	result, err := h.service.GetNearbyDrivers(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
