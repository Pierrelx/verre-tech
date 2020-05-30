// package faker

// import (
// 	"time"

// 	"database/sql"

// 	store "github.com/Pierrelx/verre-tech/store"
// 	"github.com/icrowley/fake"

// 	//pq driver pour le sql
// 	_ "github.com/lib/pq"
// )

// // main teste la connection à la base de données
// var db *sql.DB

// func Connexion(){
// 	var err error
// 	db, err = sql.Open("postgres", "port=5432 host=176.132.70.65 user=vtuser password=vt2020@oid78 dbname=verre_tech sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	faker.FakeStore()
// }


// func (db *sql.DB) CreateStore(store store.Store) (int64, error) {
// 	stmt, err := db.Prepare("INSERT INTO stores(name, address, postal_code, county, city, type, latitude, longitude, created_on_utc, updated_on_utc) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id AS id")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	store.CreatedOn = time.Now().Unix()
// 	store.UpdatedOn = time.Now().Unix()
// 	res, err := stmt.Exec(store.Name, store.Address, store.PostalCode, store.County, store.City, store.Type, store.Latitude, store.Longitude, store.CreatedOn, store.UpdatedOn)
// 	if err != nil {
// 		return 0, err
// 	}
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return id, err
// 	}
// 	return id, nil
// }

// //https://godoc.org/github.com/icrowley/fake
// func FakeStore() {
// 	err := fake.SetLang("fr")
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i := 0; i < 45; i++ {
// 		storeToCreate := new(store.Store)
// 		storeToCreate.Name = fake.Industry()
// 		storeToCreate.Address = fake.StreetAddress()
// 		storeToCreate.PostalCode = fake.Zip()
// 		storeToCreate.County = fake.Word()
// 		storeToCreate.City = fake.City()
// 		storeToCreate.Type = fake.Word()
// 		storeToCreate.Latitude = fake.Latitude()
// 		storeToCreate.Longitude = fake.Longitude()
// 		storeToCreate.CreatedOn = int64(time.Now().Unix())
// 		storeToCreate.UpdatedOn = int64(time.Now().Unix())

// 		db.CreateStore(storeToCreate)
// 	}
// }
