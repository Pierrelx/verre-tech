package database

import (
	"database/sql"
	"log"

	//Driver pour postgresql (à vérifier)
	_ "github.com/lib/pq"
)

func main() {

}

//Connection permet d'ouvrir la base de données
func Connection() {
	connStr := "user=postgres dbname=TestGolang password=JuVT75XX  host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("CREATE TABLE TestTable(ID int)")
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		print(rows)
	}
}
