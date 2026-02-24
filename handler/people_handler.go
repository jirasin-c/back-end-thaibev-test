package handler

import (
	"net/http"
	"thaibev-test/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePerson(c *gin.Context) {
	var req service.CreatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := h.PeopleService.Create(req)
	if err != nil {
		switch err {
		case service.ErrBadRequest:
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, res)
}
