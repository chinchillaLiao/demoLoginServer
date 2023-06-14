package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// hands := handlers{db}
	router.Use(DBMiddleware(db))
	router.POST("/user/register", register)
	router.POST("/user/login", login)

	return router
}

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
