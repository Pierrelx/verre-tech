package stores

//Store est un magasin
type Store struct {
	ID         int
	Name       string
	Address    string
	PostalCode string
	County     string
	City       string
	Type       string
	Latitude   float32
	Longitude  float32
}
