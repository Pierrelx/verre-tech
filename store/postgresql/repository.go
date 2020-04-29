package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	//pq driver pour le sql
	_ "github.com/lib/pq"

	"github.com/PierreLx/verre-tech/store"
)

var (
	//ErrRepository est l'erreur de base de données
	ErrRepository = errors.New("Unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

//New retourne un repo avec postgresql
func New(db *sql.DB, logger log.Logger) (store.Repository, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "postgresql"),
	}, nil
}

func (repo *repository) CreateStore(ctx context.Context, store store.Store) (int64, error) {
	stmt, err := repo.db.Prepare("INSERT INTO store(name, address, postal_code, county, city, type, latitude, longitude, created_on, updated_on) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id AS id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(store.Name, store.Address, store.PostalCode, store.County, store.City, store.Type, store.Latitude, store.Longitude, store.CreatedOn, store.UpdatedOn)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

func (repo *repository) GetStoreByID(ctx context.Context, id int) (store.Store, error) {
	var store store.Store
	stmt, err := repo.db.Prepare("SELECT id, name, address, postal_code, county, city, type, latitude, longitude, created_on, updated_on FROM store WHERE id = $1")
	if err != nil {
		return store, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&store.ID, &store.Name, &store.Address, &store.PostalCode, &store.County, &store.City, &store.Type, &store.Latitude, &store.Longitude, &store.CreatedOn, &store.UpdatedOn)
	if err != nil {
		return store, err
	}
	return store, err
}

func (repo *repository) UpdateStore(ctx context.Context, store store.Store) (store.Store, error) {
	stmt, err := repo.db.Prepare("UPDATE store SET name = $1, address = $2, postal_code = $3, county = $4, city = $5, type = $6, latitude = $7, longitude = $8, created_on = $9, updated_on = $10 WHERE id = $11")
	if err != nil {
		println(err.Error())
		return store, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(store.Name, store.Address, store.PostalCode, store.County, store.City, store.Type, store.Latitude, store.Longitude, store.CreatedOn, store.UpdatedOn, store.ID).Scan(&store.ID, &store.Name, &store.Address, &store.PostalCode, &store.County, &store.City, &store.Type, &store.Latitude, &store.Longitude, &store.CreatedOn, &store.UpdatedOn)
	if err != nil {
		println(err.Error())
		return store, nil
	}
	return store, err
}

func (repo *repository) DeleteStore(ctx context.Context, id int) error {
	stmt, err := repo.db.Prepare("DELETE FROM store WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetAll(ctx context.Context) ([]*store.Store, error) {
	stores := make([]*store.Store, 0)
	rows, err := repo.db.Query("SELECT * FROM store")
	if err != nil {
		return stores, ErrRepository
	}
	defer rows.Close()
	for rows.Next() {
		store := new(store.Store)
		err := rows.Scan(&store.ID, &store.City, &store.Address, &store.PostalCode, &store.Latitude, &store.Longitude, &store.Type, &store.Name, &store.County, &store.CreatedOn, &store.UpdatedOn)
		if err != nil {
			return stores, err
		}
		stores = append(stores, store)
	}
	return stores, err
}
