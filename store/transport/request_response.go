package transport

import (
	"github.com/PierreLx/verre-tech-ms/store"
)

//CreateRequest avec les paramètres pour la création d'un magasin
type CreateRequest struct {
	Store store.Store
}

//CreateResponse paramètres de réponse à la création d'un magasin
type CreateResponse struct {
	ID  int64 `json:"id"`
	Err error `json:"error,omitempty"`
}

//GetByIDRequest contient les paramètres pour obtenir un magasin grâce à son id
type GetByIDRequest struct {
	ID int
}

//GetByIDResponse contient la réponse pour un magasin
type GetByIDResponse struct {
	Store store.Store `json:"store"`
	Err   error       `json:"error,omitempty"`
}
