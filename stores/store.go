package store

import (
	"net/http"

	st "github.com/PierreLx/verre-tech/structs/stores"

	"github.com/go-chi/chi"
)

var shop st.Store

func main() {
	shop.ID = 45
	shop.Name = "Test"

	r := chi.NewRouter()
	r.Get("/shop", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(shop.Name))
	})
	http.ListenAndServe(":3000", r)
}

//HandleRoutes redirige les routes
func HandleRoutes(r chi.Router) {
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/oui", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oui"))
	})
}
