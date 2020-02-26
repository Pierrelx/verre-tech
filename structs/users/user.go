package users

//User représente un utilisateur
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Address     string
	PostalCode  string
	PhoneNumber int
	Email       string
	Password    string
	StoreID     int
}
