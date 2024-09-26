package main

import (
    "log"
    "shopping-cart-backend/controllers"
    "shopping-cart-backend/services"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "time"
)

func main() {
    // Connect to MongoDB
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    db := client.Database("shopping_cart")

    // Initialize services
    services.InitProductService(db)
    services.InitCartService(db)

    // Initialize Gin server
    r := gin.Default()

    // Set trusted proxies (for example, localhost)
    err = r.SetTrustedProxies([]string{"127.0.0.1"})
    if err != nil {
        log.Fatal(err)
    }

    // Catalog routes
    r.POST("/catalog/add", controllers.AddProductToCatalog)

    // Cart routes
    r.POST("/cart/add", controllers.AddProductToCart)
    r.DELETE("/cart/remove", controllers.RemoveProductFromCart)
    r.PUT("/cart/update", controllers.UpdateProductQuantity)
    r.GET("/cart/items-count", controllers.GetCartItemsCount)
    r.GET("/cart/total-price", controllers.GetCartTotalPrice)

    // Run server
    r.Run(":8080")
}
