package store

import "context"

//Store représente un magasin
type Store struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	PostalCode string  `json:"postal_code"`
	County     string  `json:"county"`
	City       string  `json:"city"`
	Type       string  `json:"type"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
	CreatedOn  int64   `json:"created_on,omitempty"`
	UpdatedOn  int64   `json:"updated_on,omitempty"`
}

// Repository pour les magasins
type Repository interface {
	CreateStore(ctx context.Context, store Store) (int64, error)
	GetStoreByID(ctx context.Context, id int) (Store, error)
}
