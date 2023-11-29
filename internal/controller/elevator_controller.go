package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/elevator-management-go/internal/service"
)

func StartServer() {
	router := gin.Default()

	router.GET("/elevators", GetElevators)
	router.GET("/hotels", GetHotels)
	// Add other routes for elevator operations and travel history

	router.Run(":8080")
}

func GetElevators(c *gin.Context) {
	elevators, err := service.GetElevators()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, elevators)
}

func GetHotels(c *gin.Context) {
	hotels, err := service.GetHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotels)
}
