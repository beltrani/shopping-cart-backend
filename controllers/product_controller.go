package controllers

import (
    "net/http"
    "shopping-cart-backend/models"
    "shopping-cart-backend/services"
    "github.com/gin-gonic/gin"
)

// Endpoint to add a product to the catalog
func AddProductToCatalog(c *gin.Context) {
    var product models.Product
    if err := c.BindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
        return
    }

    err := services.AddProductToCatalog(product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding product to catalog"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product added to catalog successfully"})
}
