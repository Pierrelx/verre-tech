package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PierreLx/verre-tech/store-microservice/common"
)

//PostgreSQL est un objet représentant une base de données
type PostgreSQL struct {
	PgDb         *sql.DB
	DatabaseName string
}

//Connection permet d'ouvrir la base de données
func Connection(db *PostgreSQL) error {
	db.DatabaseName = common.Config.DbUserName

	connString := fmt.Sprintf("user=%u dbname=%n password=%p  host=%h port=%o sslmode=disable", common.Config.DbUserName, common.Config.DbPassword, common.Config.DbAddress, common.Config.DbHost, common.Config.DbPort)
	var err error
	db.PgDb, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
		return err
	}

	println("Base de données connectée")
	return nil
}

//Close ferme la session sql
func (db *PostgreSQL) Close() {
	if db.PgDb != nil {
		db.PgDb.Close()
	}
}
