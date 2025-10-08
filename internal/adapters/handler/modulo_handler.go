package handler

import (
	"net/http"

	"fsm-modulo-three/internal/core/domain"
	"fsm-modulo-three/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type ModuloHandler struct {
	Service ports.ModuloService
}

func NewModuloHandler(service ports.ModuloService) *ModuloHandler {
	return &ModuloHandler{Service: service}
}

// Check godoc
// @Summary Compute modulo of a binary string
// @Description Returns remainder and divisibility of a binary string modulo a number
// @Tags modulo
// @Accept json
// @Produce json
// @Param request body domain.BinaryRequest true "Binary request"
// @Success 200 {object} domain.BinaryResponse
// @Failure 400 {object} map[string]string
// @Router /api/check [post]
func (h *ModuloHandler) Check(c *gin.Context) {
	var req domain.BinaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	rem, err := h.Service.Compute(req.Binary, req.Mod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.BinaryResponse{
		Binary:      req.Binary,
		Remainder:   rem,
		IsDivisible: rem == 0,
	})
}
