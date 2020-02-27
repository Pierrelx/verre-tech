package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/PierreLx/verre-tech/services"
	store "github.com/PierreLx/verre-tech/stores"
	storeSt "github.com/PierreLx/verre-tech/structs/stores"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Démarrage du serveur")

	var storeItem storeSt.Store
	storeItem.Name = "Store"
	storeItem.Address = "Ici et là"
	storeItem.PostalCode = "76000"
	storeItem.County = "Seine-Maritime"
	storeItem.City = "City"
	storeItem.Type = "Classique"
	storeItem.Latitude = 46.2555
	storeItem.Longitude = 1.6497

	db.InsertStore(storeItem)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Page d'accueil"))
	})

	r.Route("/store", func(r chi.Router) {
		store.HandleRoutes(r)
	})

	r.Get("/testStore", func(w http.ResponseWriter, r *http.Request) {
		toReturn, _ := json.Marshal(storeItem)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(toReturn))
	})

	http.ListenAndServe(":3000", r)
}
