package models

// Product structure represents a product in the catalog
type Product struct {
    ID    string  `json:"id" bson:"_id"`
    Name  string  `json:"name" bson:"name"`
    Price float64 `json:"price" bson:"price"`
}

// CartItem structure represents a product item in the cart, with a calculated subtotal
type CartItem struct {
    Product   Product `json:"product" bson:"product"`
    Quantity  int     `json:"quantity" bson:"quantity"`
    Subtotal  float64 `json:"subtotal" bson:"subtotal"` // Subtotal for this item (price * quantity)
}

// Cart structure represents a shopping cart
type Cart struct {
    UserID     string     `json:"userId" bson:"userId"`
    Items      []CartItem `json:"items" bson:"items"`
    TotalPrice float64    `json:"totalPrice" bson:"totalPrice"`
    TotalItems int        `json:"totalItems" bson:"totalItems"`
}