package loader

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"lumelTask1/models" // Adjust the import path based on your project structure

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LoadData loads the CSV file into the MongoDB collections
func LoadData(csvFilePath string) error {

	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV headers: %v", err)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("sales_db")
	orderCollection := db.Collection("orders")
	productCollection := db.Collection("products")
	customerCollection := db.Collection("customers")

	// Maps to track inserted products and customers
	productMap := make(map[string]bool)
	customerMap := make(map[string]bool)

	// Process each row in the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV rows: %v", err)
	}

	for _, row := range rows {
		// Parse row data
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

		// Insert Product if not exists
		if !productMap[productID] {
			product := models.Product{
				ProductID:   productID,
				ProductName: productName,
				Category:    category,
				UnitPrice:   unitPrice,
			}
			_, err := productCollection.InsertOne(context.Background(), product)
			if err != nil {
				log.Printf("failed to insert product %s: %v", productID, err)
				continue
			}
			productMap[productID] = true
		}

		// Insert Customer if not exists
		if !customerMap[customerID] {
			customer := models.Customer{
				CustomerID:    customerID,
				CustomerName:  customerName,
				CustomerEmail: customerEmail,
				CustomerAddr:  customerAddr,
			}
			_, err := customerCollection.InsertOne(context.Background(), customer)
			if err != nil {
				log.Printf("failed to insert customer %s: %v", customerID, err)
				continue
			}
			customerMap[customerID] = true
		}

		// Insert Order
		order := models.Order{
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
		}
		_, err := orderCollection.InsertOne(context.Background(), order)
		if err != nil {
			log.Printf("failed to insert order %s: %v", orderID, err)
			continue
		}
	}

	log.Println("CSV data successfully loaded into MongoDB.")
	return nil
}

// parseInt converts a string to an integer
func parseInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("failed to parse integer: %v", err)
		return 0
	}
	return result
}

// parseFloat converts a string to a float
func parseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("failed to parse float: %v", err)
		return 0.0
	}
	return result
}
