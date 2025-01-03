package controllers

import (
	"lumelTask1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTotalRevenue(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	revenue, err := services.CalculateTotalRevenue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_revenue": revenue})
}

func GetRevenueByProduct(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	revenueByProduct, err := services.CalculateRevenueByProduct(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, revenueByProduct)
}

func GetRevenueByCategory(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	revenueByCategory, err := services.CalculateRevenueByCategory(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, revenueByCategory)
}

func GetRevenueByRegion(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	revenueByRegion, err := services.CalculateRevenueByRegion(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, revenueByRegion)
}
