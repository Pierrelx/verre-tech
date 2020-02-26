package cartitems

//CartItem représente un article du panier
type CartItem struct {
	ID              int
	UserID          int
	ProductID       int
	StoreID         int
	OrderID         int
	ProductQuantity int
	UnitPrice       float32
}
