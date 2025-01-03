package controllers

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"lumelTask1/database"
	"net/http"
	"os"
	"strconv"

	"lumelTask1/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RefreshData handles the data refresh API request
func RefreshData(c *gin.Context) {
	csvFilePath := c.Query("file")
	if csvFilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CSV file path is required"})
		return
	}

	// Open the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open CSV file: %v", err)})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the headers
	_, err = reader.Read()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read CSV headers: %v", err)})
		return
	}

	client := database.GetClient()

	db := client.Database("sales_db")
	orderCollection := db.Collection("orders")
	productCollection := db.Collection("products")
	customerCollection := db.Collection("customers")

	rows, err := reader.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read CSV rows: %v", err)})
		return
	}

	// Process the rows and refresh the data
	for _, row := range rows {
		orderID := row[0]
		productID := row[1]
		customerID := row[2]
		productName := row[3]
		category := row[4]
		region := row[5]
		dateOfSale := row[6]
		quantitySold := parseInt(row[7])
		unitPrice := parseFloat(row[8])
		discount := parseFloat(row[9])
		shippingCost := parseFloat(row[10])
		paymentMethod := row[11]
		customerName := row[12]
		customerEmail := row[13]
		customerAddr := row[14]

		// Upsert Product
		_, err = productCollection.UpdateOne(
			context.Background(),
			bson.M{"product_id": productID},
			bson.M{"$set": models.Product{
				ProductID:   productID,
				ProductName: productName,
				Category:    category,
				UnitPrice:   unitPrice,
			}},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Printf("Failed to upsert product %s: %v", productID, err)
		}

		// Upsert Customer
		_, err = customerCollection.UpdateOne(
			context.Background(),
			bson.M{"customer_id": customerID},
			bson.M{"$set": models.Customer{
				CustomerID:    customerID,
				CustomerName:  customerName,
				CustomerEmail: customerEmail,
				CustomerAddr:  customerAddr,
			}},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Printf("Failed to upsert customer %s: %v", customerID, err)
		}

		// Upsert Order
		_, err = orderCollection.UpdateOne(
			context.Background(),
			bson.M{"order_id": orderID},
			bson.M{"$set": models.Order{
				OrderID:       orderID,
				ProductID:     productID,
				CustomerID:    customerID,
				ProductName:   productName,
				Category:      category,
				Region:        region,
				DateOfSale:    dateOfSale,
				QuantitySold:  quantitySold,
				UnitPrice:     unitPrice,
				Discount:      discount,
				ShippingCost:  shippingCost,
				PaymentMethod: paymentMethod,
				CustomerName:  customerName,
				CustomerEmail: customerEmail,
				CustomerAddr:  customerAddr,
			}},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Printf("Failed to upsert order %s: %v", orderID, err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data refreshed successfully"})
}

func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}

func parseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return result
}
