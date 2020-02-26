package main

import (
	"fmt"
	"net/http"

	db "github.com/PierreLx/verre-tech/services"
	store "github.com/PierreLx/verre-tech/stores"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Démarrage du serveur")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Page d'accueil"))
	})

	r.Route("/store", func(r chi.Router) {
		store.HandleRoutes(r)
	})

	db.Connection()

	http.ListenAndServe(":3000", r)
}
