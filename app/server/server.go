package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	hands := handlers{db}
	router.POST("/user/register", hands.register)
	router.POST("/user/login", hands.login)
	return router
}
