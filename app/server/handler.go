package server

import (
	mod "demoLoginServer/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func register(c *gin.Context) {
	var user = mod.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println("API:", user)

	db := c.MustGet("db").(*gorm.DB)

	err := user.Create(db)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(409, err)
	}

}

func login(c *gin.Context) {
	var user = mod.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	success, err := user.Login(db)
	if success {
		c.JSON(200, user)
	} else {
		c.JSON(401, err)
	}
}
