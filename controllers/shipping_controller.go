package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Mock data for shipping options
type ShippingOption struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
    Time  string  `json:"time"`  // Estimated delivery time
}

// Endpoint to return shipping options for a product
func GetShippingOptions(c *gin.Context) {
    // This is a mock function; you can later integrate this with an actual shipping service.
    options := []ShippingOption{
        {
            Name:  "Standard Shipping",
            Price: 5.00,
            Time:  "5-7 Business Days",
        },
        {
            Name:  "Express Shipping",
            Price: 15.00,
            Time:  "1-2 Business Days",
        },
        {
            Name:  "Overnight Shipping",
            Price: 25.00,
            Time:  "1 Business Day",
        },
    }

    // Return the shipping options as JSON
    c.JSON(http.StatusOK, options)
}
