package controllers

import (
    "net/http"
    "shopping-cart-backend/services"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/gin-gonic/gin"
)

// Endpoint to add a product to the cart
func AddProductToCart(c *gin.Context) {
    var request struct {
        ProductID string `json:"id"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
        return
    }

    userID := c.Query("userId")
    isValid, product, err := services.IsProductValid(request.ProductID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating product"})
        return
    }

    if !isValid {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    cart, err := services.AddProductToCart(userID, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding product to cart"})
        return
    }

    c.JSON(http.StatusOK, cart)
}

// Endpoint to remove a product from the cart
func RemoveProductFromCart(c *gin.Context) {
    userID := c.Query("userId")
    productID := c.Query("productId")

    cart, err := services.RemoveProductFromCart(userID, productID)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in cart"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing product from cart"})
        return
    }

    c.JSON(http.StatusOK, cart)
}

// Endpoint to update product quantity in the cart
func UpdateProductQuantity(c *gin.Context) {
    var request struct {
        ProductID   string `json:"productId"`
        NewQuantity int    `json:"newQuantity"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
        return
    }

    userID := c.Query("userId")
    cart, err := services.UpdateProductQuantity(userID, request.ProductID, request.NewQuantity)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in cart"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product quantity in cart"})
        return
    }

    c.JSON(http.StatusOK, cart)
}

// Endpoint to get the total number of items in the cart
func GetCartItemsCount(c *gin.Context) {
    userID := c.Query("userId")

    count, err := services.GetCartItemsCount(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting item count from cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"totalItems": count})
}

// Endpoint to get the total price of the cart
func GetCartTotalPrice(c *gin.Context) {
    userID := c.Query("userId")

    totalPrice, err := services.GetCartTotalPrice(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting total price of cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"totalPrice": totalPrice})
}
