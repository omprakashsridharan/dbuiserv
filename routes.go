package main

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/connect", Connect)
		api.GET("/:id/databases", GetDatabases)
		api.GET("/:id/databases/:database/tables", GetTables)
	}
	return r
}
