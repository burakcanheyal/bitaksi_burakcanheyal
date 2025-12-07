package handler

import (
	"bitaksi_burakcanheyal/driver_microservice/internal/domain/model"
	"bitaksi_burakcanheyal/driver_microservice/platform/validation"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverGatewayHandler struct {
	driverClient *model.Driver
}

func NewDriverGatewayHandler(dc *model.Driver) *DriverGatewayHandler {
	return &DriverGatewayHandler{driverClient: dc}
}

// ───────────────────────────────
// 1) POST /drivers → VALIDASYON OK
// ───────────────────────────────
func (h *DriverGatewayHandler) CreateDriver(c *gin.Context) {

	var req validation.AddDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Validation çalıştır
	if err := validation.ValidateAddDriver(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bodyBytes, _ := json.Marshal(req)

	resp, err := h.driverClient.ForwardPost(c.Request.Context(), "/drivers", bodyBytes)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}

// ───────────────────────────────
// 2) PUT /drivers  → VALIDASYON OK
//
// ───────────────────────────────
func (h *DriverGatewayHandler) UpdateDriver(c *gin.Context) {

	var req validation.UpdateDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Validation
	if err := validation.ValidateUpdateDriver(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bodyBytes, _ := json.Marshal(req)

	resp, err := h.driverClient.ForwardPut(
		c.Request.Context(),
		"/drivers/"+req.ID,
		bodyBytes,
	)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}

// ───────────────────────────────
// 3) GET /drivers?page=&pageSize= → VALIDASYON OK
// ───────────────────────────────
func (h *DriverGatewayHandler) ListDrivers(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	if err := validation.ValidateListParams(page, size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rawQuery := c.Request.URL.RawQuery
	path := "/drivers"
	if rawQuery != "" {
		path += "?" + rawQuery
	}

	resp, err := h.driverClient.ForwardGet(c.Request.Context(), path)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}

// ───────────────────────────────
// 4) POST /drivers/nearby → VALIDASYON OK
// ───────────────────────────────
func (h *DriverGatewayHandler) GetNearbyDrivers(c *gin.Context) {

	var req validation.NearbyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Validation
	if err := validation.ValidateNearby(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bodyBytes, _ := json.Marshal(req)

	resp, err := h.driverClient.ForwardPost(
		c.Request.Context(),
		"/drivers/nearby",
		bodyBytes,
	)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}
