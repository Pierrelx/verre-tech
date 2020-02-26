package orders

import "time"

//Order représente une commande
type Order struct {
	ID              int
	ProductQuantity int
	TotalPrice      float32
	WithdrawalDate  time.Time
	Status          string
	UserID          int
}
