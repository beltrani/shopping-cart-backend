package services

import (
    "context"
    "shopping-cart-backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// Global product collection reference
var productCollection *mongo.Collection

// Initialize product service with the MongoDB database
func InitProductService(db *mongo.Database) {
    productCollection = db.Collection("products")
}

// Check if a product exists in the catalog
func IsProductValid(productID string) (bool, models.Product, error) {
    var product models.Product
    err := productCollection.FindOne(context.TODO(), bson.M{"_id": productID}).Decode(&product)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return false, models.Product{}, nil
        }
        return false, models.Product{}, err
    }
    return true, product, nil
}

// Add a new product to the catalog
func AddProductToCatalog(product models.Product) error {
    _, err := productCollection.InsertOne(context.TODO(), product)
    return err
}
