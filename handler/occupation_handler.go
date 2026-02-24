package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOccupations(c *gin.Context) {
	list, err := h.OccupationService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}
