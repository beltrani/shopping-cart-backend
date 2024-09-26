package services

import (
    "context"
    "shopping-cart-backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Global cart collection reference
var cartCollection *mongo.Collection

// Initialize cart service with the MongoDB database
func InitCartService(db *mongo.Database) {
    cartCollection = db.Collection("carts")
}

// Add a product to the cart
func AddProductToCart(userID string, product models.Product) (models.Cart, error) {
    var cart models.Cart
    err := cartCollection.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&cart)

    if err != nil {
        cart = models.Cart{UserID: userID, Items: []models.CartItem{}}
    }

    found := false
    for i, item := range cart.Items {
        if item.Product.ID == product.ID {
            cart.Items[i].Quantity++
            cart.Items[i].Subtotal = cart.Items[i].Quantity * product.Price // Update subtotal
            found = true
            break
        }
    }

    if !found {
        cart.Items = append(cart.Items, models.CartItem{
            Product:  product,
            Quantity: 1,
            Subtotal: product.Price, // Subtotal when adding the product for the first time
        })
    }

    cart.TotalItems++
    cart.TotalPrice += product.Price

    _, err = cartCollection.UpdateOne(context.TODO(), bson.M{"userId": userID}, bson.M{
        "$set": cart,
    }, options.Update().SetUpsert(true))

    return cart, err
}

// Remove a product from the cart
func RemoveProductFromCart(userID, productID string) (models.Cart, error) {
    var cart models.Cart
    err := cartCollection.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&cart)

    if err != nil {
        return cart, err
    }

    found := false
    var priceToRemove float64
    for i, item := range cart.Items {
        if item.Product.ID == productID {
            priceToRemove = item.Product.Price * float64(item.Quantity)
            cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
            found = true
            break
        }
    }

    if !found {
        return cart, mongo.ErrNoDocuments
    }

    cart.TotalItems -= 1
    cart.TotalPrice -= priceToRemove

    _, err = cartCollection.UpdateOne(context.TODO(), bson.M{"userId": userID}, bson.M{
        "$set": cart,
    })

    return cart, err
}

// Update the quantity of a product in the cart
func UpdateProductQuantity(userID, productID string, newQuantity int) (models.Cart, error) {
    var cart models.Cart
    err := cartCollection.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&cart)
    if err != nil {
        return cart, err
    }

    found := false
    var priceDifference float64
    for i, item := range cart.Items {
        if item.Product.ID == productID {
            priceDifference = float64(newQuantity-item.Quantity) * item.Product.Price
            cart.Items[i].Quantity = newQuantity
            cart.Items[i].Subtotal = float64(newQuantity) * item.Product.Price // Update subtotal
            found = true
            break
        }
    }

    if !found {
        return cart, mongo.ErrNoDocuments
    }

    cart.TotalPrice += priceDifference

    _, err = cartCollection.UpdateOne(context.TODO(), bson.M{"userId": userID}, bson.M{
        "$set": cart,
    })

    return cart, err
}

// Get the total number of items in the cart
func GetCartItemsCount(userID string) (int, error) {
    var cart models.Cart
    err := cartCollection.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&cart)
    if err != nil {
        return 0, err
    }

    return cart.TotalItems, nil
}

// Get the total price of the cart
func GetCartTotalPrice(userID string) (float64, error) {
    var cart models.Cart
    err := cartCollection.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&cart)
    if err != nil {
        return 0, err
    }

    return cart.TotalPrice, nil
}
