package transport

import (
	"github.com/Pierrelx/verre-tech/store"
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

//UpdateRequest est le corps d'une requête d'update
type UpdateRequest struct {
	Store store.Store `json:"store"`
}

//UpdateResponse est le corps d'une réponse d'update
type UpdateResponse struct {
	Store store.Store `json:"store"`
	Err   error       `json:"error"`
}

//DeleteRequest est le corps d'une requête de suppression
type DeleteRequest struct {
	ID int `json:"id"`
}

//DeleteResponse est le corps de la réponse de suppression
type DeleteResponse struct {
	Err error `json:"error"`
}

//GetAllRequest est le corps de la requête de liste
type GetAllRequest struct {
}

//ListStoreResponse est le corps de la réponse de liste des magasins
type ListStoreResponse struct {
	Stores []*store.Store `json:"stores"`
	Err    error          `json:"error"`
}
