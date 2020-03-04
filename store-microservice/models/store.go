package stores

//Store est un magasin
type Store struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	PostalCode string  `json:"postalCode"`
	County     string  `json:"county"`
	City       string  `json:"city"`
	Type       string  `json:"type"`
	Latitude   float32 `json:"latitude"`
	Longitude  float32 `json:"longitude"`
}
