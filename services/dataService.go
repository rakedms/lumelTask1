package services

import (
	"context"
	"fmt"
	"lumelTask1/database"
	"lumelTask1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CalculateTotalRevenue calculates the total revenue for a date range
func CalculateTotalRevenue(startDate, endDate string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	filter := bson.M{
		"date_of_sale": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var totalRevenue float64
	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return 0, fmt.Errorf("error decoding order: %v", err)
		}
		totalRevenue += float64(order.QuantitySold) * order.UnitPrice
	}
	return totalRevenue, nil
}

func CalculateRevenueByProduct(startDate, endDate string) (map[string]float64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$product_name",
				"revenue": bson.M{
					"$sum": bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price"}},
				},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	revenueByProduct := make(map[string]float64)
	for cursor.Next(ctx) {
		var result struct {
			ProductName string  `bson:"_id"`
			Revenue     float64 `bson:"revenue"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		revenueByProduct[result.ProductName] = result.Revenue
	}
	return revenueByProduct, nil
}

func CalculateRevenueByCategory(startDate, endDate string) (map[string]float64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$category",
				"revenue": bson.M{
					"$sum": bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price"}},
				},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	revenueByCategory := make(map[string]float64)
	for cursor.Next(ctx) {
		var result struct {
			Category string  `bson:"_id"`
			Revenue  float64 `bson:"revenue"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		revenueByCategory[result.Category] = result.Revenue
	}
	return revenueByCategory, nil
}

func CalculateRevenueByRegion(startDate, endDate string) (map[string]float64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$region",
				"revenue": bson.M{
					"$sum": bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price"}},
				},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	revenueByRegion := make(map[string]float64)
	for cursor.Next(ctx) {
		var result struct {
			Region  string  `bson:"_id"`
			Revenue float64 `bson:"revenue"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		revenueByRegion[result.Region] = result.Revenue
	}
	return revenueByRegion, nil
}

func GetTopProductsOverall(startDate, endDate, limit string) ([]models.Order, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$product_name",
				"total_quantity": bson.M{
					"$sum": "$quantity_sold",
				},
			}},
		},
		{
			{Key: "$sort", Value: bson.M{
				"total_quantity": -1, // Sort by total quantity sold, in descending order
			}},
		},
		{
			{Key: "$limit", Value: 10}, // Limit to top 10 products
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error getting top products overall: %v", err)
	}
	defer cursor.Close(ctx)

	var topProducts []models.Order
	for cursor.Next(ctx) {
		var result struct {
			ProductName   string `bson:"_id"`
			TotalQuantity int    `bson:"total_quantity"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		topProducts = append(topProducts, models.Order{
			ProductName:  result.ProductName,
			QuantitySold: result.TotalQuantity,
		})
	}

	return topProducts, nil
}

func GetTopProductsByCategory(startDate, endDate, category, limit string) ([]models.Order, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
				"category": category,
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$product_name",
				"total_quantity": bson.M{
					"$sum": "$quantity_sold",
				},
			}},
		},
		{
			{Key: "$sort", Value: bson.M{
				"total_quantity": -1, // Sort by total quantity sold, in descending order
			}},
		},
		{
			{Key: "$limit", Value: 10}, // Limit to top 10 products
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error getting top products by category: %v", err)
	}
	defer cursor.Close(ctx)

	var topProductsByCategory []models.Order
	for cursor.Next(ctx) {
		var result struct {
			ProductName   string `bson:"_id"`
			TotalQuantity int    `bson:"total_quantity"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		topProductsByCategory = append(topProductsByCategory, models.Order{
			ProductName:  result.ProductName,
			QuantitySold: result.TotalQuantity,
		})
	}

	return topProductsByCategory, nil
}

// GetTopProductsByRegion retrieves the top N products by region based on quantity sold within a date range
func GetTopProductsByRegion(startDate, endDate, region, limit string) ([]models.Order, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
				"region": region,
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id": "$product_name",
				"total_quantity": bson.M{
					"$sum": "$quantity_sold",
				},
			}},
		},
		{
			{Key: "$sort", Value: bson.M{
				"total_quantity": -1, // Sort by total quantity sold, in descending order
			}},
		},
		{
			{Key: "$limit", Value: 10}, // Limit to top 10 products
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error getting top products by region: %v", err)
	}
	defer cursor.Close(ctx)

	var topProductsByRegion []models.Order
	for cursor.Next(ctx) {
		var result struct {
			ProductName   string `bson:"_id"`
			TotalQuantity int    `bson:"total_quantity"`
		}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error decoding result: %v", err)
		}
		topProductsByRegion = append(topProductsByRegion, models.Order{
			ProductName:  result.ProductName,
			QuantitySold: result.TotalQuantity,
		})
	}

	return topProductsByRegion, nil
}

func GetTotalCustomers(startDate, endDate string) (int64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Use distinct to count unique customers within the date range
	cursor, err := collection.Distinct(ctx, "customer_id", bson.M{
		"date_of_sale": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	})
	if err != nil {
		return 0, fmt.Errorf("error fetching distinct customer IDs: %v", err)
	}
	return int64(len(cursor)), nil
}

// GetTotalOrders fetches the total number of orders within the given date range
func GetTotalOrders(startDate, endDate string) (int64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.M{
		"date_of_sale": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	})
	if err != nil {
		return 0, fmt.Errorf("error counting orders: %v", err)
	}
	return count, nil
}

// GetAverageOrderValue calculates the average order value within the given date range
func GetAverageOrderValue(startDate, endDate string) (float64, error) {
	client := database.GetClient()
	collection := client.Database("sales_db").Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.M{
				"date_of_sale": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			}},
		},
		{
			{Key: "$group", Value: bson.M{
				"_id":          nil,
				"total_value":  bson.M{"$sum": bson.M{"$multiply": []interface{}{"$quantity_sold", "$unit_price"}}},
				"total_orders": bson.M{"$sum": 1},
			}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %v", err)
	}
	defer cursor.Close(ctx)

	var result struct {
		TotalValue  float64 `bson:"total_value"`
		TotalOrders int64   `bson:"total_orders"`
	}

	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			return 0, fmt.Errorf("error decoding result: %v", err)
		}
		if result.TotalOrders == 0 {
			return 0, nil // Avoid division by zero
		}
		averageOrderValue := result.TotalValue / float64(result.TotalOrders)
		return averageOrderValue, nil
	}

	return 0, fmt.Errorf("no data found")
}
