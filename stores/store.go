package store

import (
	"net/http"
	"strconv"

	db "github.com/PierreLx/verre-tech/services/database"
	st "github.com/PierreLx/verre-tech/structs/stores"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func main() {

	r := chi.NewRouter()
	r.Get("/shop", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Coucou"))
	})
	http.ListenAndServe(":3000", r)
}

//HandleRoutes redirige les routes
func HandleRoutes(r chi.Router) {
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", GetStore)
		r.Delete("/", DeleteStore)
		r.Put("/", UpdateStore)
	})

	r.Get("/all", GetAllStores)
}

//GetStore retourne un magasin selon son id
func GetStore(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if id < 0 || err != nil {
		render.Respond(w, r, []byte("Non trouvé"))
	} else {
		render.Respond(w, r, []byte("Retourner le store en JSON"))
	}
}

//GetAllStores permet de retourner tous les magasins existants
func GetAllStores(w http.ResponseWriter, r *http.Request) {

}

//CreateStore permet de créer un magasin
func CreateStore(store st.Store) (bool, error) {
	return db.InsertStore(store)
}

//UpdateStore permet de faire la mise à jour d'un magasin
func UpdateStore(w http.ResponseWriter, r *http.Request) {

}

//DeleteStore permet  de supprimer un store
func DeleteStore(w http.ResponseWriter, r *http.Request) {

}
