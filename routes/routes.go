package routes

import (
	"lumelTask1/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/refresh", controllers.RefreshData)

	router.GET("/api/revenue", controllers.GetTotalRevenue)
	router.GET("/api/revenue-by-product", controllers.GetRevenueByProduct)
	router.GET("/api/revenue-by-category", controllers.GetRevenueByCategory)
	router.GET("/api/revenue-by-region", controllers.GetRevenueByRegion)

	router.GET("/api/top-products", controllers.GetTopProductsOverall)
	router.GET("/api/top-products-by-category", controllers.GetTopProductsByCategory)
	router.GET("/api/top-products-by-region", controllers.GetTopProductsByRegion)

	router.GET("/api/total-customers", controllers.GetTotalCustomers)
	router.GET("/api/total-orders", controllers.GetTotalOrders)
	router.GET("/api/average-order-value", controllers.GetAverageOrderValue)

}
