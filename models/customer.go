package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    string             `bson:"customer_id"`    // Unique customer ID
	CustomerName  string             `bson:"customer_name"`  // Name of the customer
	CustomerEmail string             `bson:"customer_email"` // Email of the customer
	CustomerAddr  string             `bson:"customer_address"`
}
