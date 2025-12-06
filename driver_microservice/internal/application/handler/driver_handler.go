package handler

import (
	"bitaksi_burakcanheyal/driver_microservice/internal/domain/model"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DriverGatewayHandler struct {
	driverClient *model.Driver
}

func NewDriverGatewayHandler(dc *model.Driver) *DriverGatewayHandler {
	return &DriverGatewayHandler{driverClient: dc}
}

// ───────────────────────────────
// POST /drivers
// ───────────────────────────────
func (h *DriverGatewayHandler) CreateDriver(c *gin.Context) {

	bodyBytes, _ := io.ReadAll(c.Request.Body)

	resp, err := h.driverClient.ForwardPost(
		c.Request.Context(),
		"/drivers",
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
// PUT /drivers/:id
// ───────────────────────────────
func (h *DriverGatewayHandler) UpdateDriver(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	bodyBytes, _ := io.ReadAll(c.Request.Body)

	resp, err := h.driverClient.ForwardPut(
		c.Request.Context(),
		"/drivers/"+id,
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
// GET /drivers
// ───────────────────────────────
func (h *DriverGatewayHandler) ListDrivers(c *gin.Context) {

	// Query paramlarını aynen DB servisine aktar
	rawQuery := c.Request.URL.RawQuery
	path := "/drivers"
	if rawQuery != "" {
		path += "?" + rawQuery
	}

	resp, err := h.driverClient.ForwardGet(
		c.Request.Context(),
		path,
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
// GET /drivers/nearby
// ───────────────────────────────
func (h *DriverGatewayHandler) GetNearbyDrivers(c *gin.Context) {

	rawQuery := c.Request.URL.RawQuery
	path := "/drivers/nearby"
	if rawQuery != "" {
		path += "?" + rawQuery
	}

	resp, err := h.driverClient.ForwardGet(
		c.Request.Context(),
		path,
	)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}
