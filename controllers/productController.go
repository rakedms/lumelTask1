package controllers

import (
	"lumelTask1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopProductsOverall(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")
	limit := c.DefaultQuery("limit", "10")

	topProducts, err := services.GetTopProductsOverall(startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"top_products_overall": topProducts})
}

func GetTopProductsByCategory(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")
	limit := c.DefaultQuery("limit", "10")
	category := c.DefaultQuery("category", "")

	topProductsByCategory, err := services.GetTopProductsByCategory(startDate, endDate, category, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"top_products_by_category": topProductsByCategory})
}

func GetTopProductsByRegion(c *gin.Context) {
	startDate := c.DefaultQuery("start_date", "2020-01-01")
	endDate := c.DefaultQuery("end_date", "2025-01-01")
	limit := c.DefaultQuery("limit", "10")
	region := c.DefaultQuery("region", "")

	topProductsByRegion, err := services.GetTopProductsByRegion(startDate, endDate, region, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"top_products_by_region": topProductsByRegion})
}
