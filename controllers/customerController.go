package controllers

import (
	"lumelTask1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTotalCustomers(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	totalCustomers, err := services.GetTotalCustomers(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_customers": totalCustomers})
}

func GetTotalOrders(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	totalOrders, err := services.GetTotalOrders(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_orders": totalOrders})
}

func GetAverageOrderValue(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")

	averageOrderValue, err := services.GetAverageOrderValue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"average_order_value": averageOrderValue})
}
