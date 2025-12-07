package handler

import (
	"bitaksi_burakcanheyal/db_microservice/internal/application"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/dto"
	"bitaksi_burakcanheyal/db_microservice/internal/domain/service"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type DriverHandler struct {
	service *service.DriverService
}

func NewDriverHandler(s *service.DriverService) *DriverHandler {
	return &DriverHandler{service: s}
}

// CREATE
func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var req dto.CreateDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.Wrap("ERR_BAD_REQUEST"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id, err := h.service.CreateDriver(ctx, req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, dto.CreateDriverResponse{ID: id})
}

// UPDATE
func (h *DriverHandler) UpdateDriver(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.Error(application.ErrMissingID)
		return
	}

	var req dto.UpdateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.Wrap("ERR_BAD_REQUEST"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.service.UpdateDriver(ctx, id, req); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"status": "updated"})
}

// LIST
func (h *DriverHandler) ListDrivers(c *gin.Context) {

	pageStr := c.Query("page")
	sizeStr := c.Query("pageSize")

	page, err1 := strconv.Atoi(pageStr)
	size, err2 := strconv.Atoi(sizeStr)

	if err1 != nil || page < 1 {
		page = 1
	}
	if err2 != nil || size < 1 {
		size = 20
	}

	drivers, err := h.service.ListDrivers(c.Request.Context(), page, size)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, drivers)
}

// NEARBY
func (h *DriverHandler) GetNearbyDrivers(c *gin.Context) {

	var req dto.NearbyDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(application.Wrap("ERR_BAD_REQUEST"))
		return
	}

	result, err := h.service.GetNearbyDrivers(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
