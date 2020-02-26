package database

import (
	"database/sql"
	"log"

	storeSt "github.com/PierreLx/verre-tech/structs/stores"

	//Driver pour postgresql (à vérifier)
	_ "github.com/lib/pq"
)

const connStr = "user=postgres dbname=TestGolang password=JuVT75XX  host=localhost port=5432 sslmode=disable"

func main() {
	db, err := Connection()

	err = db.Ping()
	if err != nil {
		println("Erreur de connection à la base de données => ", err.Error)
	}

	defer db.Close()
}

//Connection permet d'ouvrir la base de données
func Connection() (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	println("Base de données connectée")
	return db, nil
}

//InsertStore insère une entité
func InsertStore(item storeSt.Store) {
	localDB, err := Connection()

	query, err := localDB.Prepare("INSERT INTO Stores(Name, Address, PostalCode, County, City, Type, Latitude, Longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		println("Erreur lors de la préparation de la requête")
		panic(err.Error())
	}

	defer query.Close()

	res, err := query.Exec(item.Name, item.Address, item.PostalCode, item.County, item.City, item.Type, item.Latitude, item.Longitude)
	if err != nil {
		println("Erreur lors de l'exécution de la requête")
		panic(err.Error())
	}
	print(res)
}
