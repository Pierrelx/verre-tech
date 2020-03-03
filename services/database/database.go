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
func InsertStore(item storeSt.Store) (bool, int64, error) {
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
		return false, 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return true, lastID, nil
	}
	return false, 0, err
}

//UpdateStore met à jour une boutique
func UpdateStore(item storeSt.Store) (bool, int64, error) {
	localDB, err := Connection()

	query, err := localDB.Prepare("UPDATE Stores SET Name = $1, Address = $2, PostalCode = $3, County = $4, City = $5, Type = $6, Latitude = $7, Longitude = $8 WHERE Id = $9")

	if err != nil {
		println("Erreur lors de la préparation de la requête UPDATE STORE")
		panic(err.Error())
	}

	defer query.Close()

	res, err := query.Exec(item.Name, item.Address, item.PostalCode, item.County, item.City, item.Type, item.Latitude, item.Longitude, item.ID)
	if err != nil {
		println("Erreur lors de l'execution de la requête UPDATE STORE")
		return false, 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return false, 0, err
	}
	return true, rows, nil
}

//DeleteStore permet de supprimer un magasin
func DeleteStore(storeID int64) (bool, error) {
	localDB, err := Connection()

	query, err := localDB.Prepare("DELETE FROM Stores WHERE Id = $1")

	if err != nil {
		println("Erreur lors de la préparation de la requête DELETE STORE")
		panic(err.Error())
	}

	res, err := query.Exec(storeID)
	if err != nil {
		println("Erreur lors de l'exécution de la requête DELETE STORE")
		return false, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil && rowsCount < 0 {
		return false, err
	}
	return true, nil
}
