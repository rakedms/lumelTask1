package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductID   string             `bson:"product_id"`   // Unique product ID
	ProductName string             `bson:"product_name"` // Product name
	Category    string             `bson:"category"`     // Product category
	UnitPrice   float64            `bson:"unit_price"`   // Price per unit
}
