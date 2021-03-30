package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Connect(c *gin.Context) {
	var connectionOptions ConnectionOptions
	if err := c.ShouldBindJSON(&connectionOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbConnection := CreateDbConnection(connectionOptions)
	err := (*dbConnection).init()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		id := uuid.New().String()
		AddConnection(id, dbConnection)
		c.JSON(http.StatusOK, gin.H{"connectionId": id})
	}
}

func GetDatabases(c *gin.Context) {
	dbConnectionId := c.Param("id")
	dbConnection := GetConnection(dbConnectionId)
	if dbConnection != nil {
		result := (*dbConnection).GetDatabases()
		c.JSON(http.StatusOK, gin.H{"data": result})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid connection id"})
	}
}

func GetTables(c *gin.Context) {
	dbConnectionId := c.Param("id")
	database := c.Param("database")
	dbConnection := GetConnection(dbConnectionId)
	if dbConnection != nil {
		result := (*dbConnection).GetTables(database)
		c.JSON(http.StatusOK, gin.H{"data": result})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid connection id"})
	}
}
