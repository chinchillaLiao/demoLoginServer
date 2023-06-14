package server

import (
	mod "demoLoginServer/models"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlers struct {
	db *gorm.DB
}

func (h *handlers) register(c *gin.Context) {
	var user = mod.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("API:", user)
	user.Create(h.db)
	c.JSON(200, user)
}

func (h *handlers) login(c *gin.Context) {}
