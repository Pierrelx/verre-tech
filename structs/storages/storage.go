package storages

//Storage représente un entrepôt
type Storage struct {
	ID         int
	City       string
	Address    string
	PostalCode string
	County     string
	Name       string
	Latitude   float32
	Longitude  float32
}
