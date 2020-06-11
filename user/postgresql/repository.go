package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-kit/kit/log"
	//pq driver pour le sql
	_ "github.com/lib/pq"

	u "github.com/PierreLx/verre-tech/user"
)

var (
	//ErrRepository est l'erreur de base de donnnées
	ErrRepository = errors.New("Unable to handle request")
)

type repository struct{
	db  	*sql.DB
	logger 	log.Logger
}

func New(db *sql.DB, logger log.Logger)(u.Repository, error){
	return &repository{
		db: db,
		logger: log.With(logger, "repository", "postgresql")
	}
}

func (repo *repository) CreateUser(ctx context.Context, user u.User)(int64, error){
	stmt, err := repo.db.Prepare("INSERT INTO users(first_name, last_name, address, city, postal_code, email, created_on_utc, updated_on_utc) VALUES($1, $2, $3, $4, $5, $6, 7, $8) RETURNING id as id")
	if err != nil{
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.FirstName, user.LastName)

}