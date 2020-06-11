package user 

import "context"

type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	City string `json:"city"`
	PostalCode string `json:"postal_code"`
	Email string `json:"email"`
	CreatedOn  int64   `json:"created_on_utc,omitempty"`
	UpdatedOn  int64   `json:"updated_on_utc,omitempty"`
}

//Gestion des mots de passe
//https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
//https://gowebexamples.com/password-hashing/
type UserPassword struct{
	ID int `json:"int"`
	UserId int `json:"user_id"`
	PasswordHash string `json:"password_hash"`
	CreatedOn  int64   `json:"created_on_utc,omitempty"`
	UpdatedOn  int64   `json:"updated_on_utc,omitempty"`
}

type Repository interface{
	CreateUser(ctx context.Context, user User)(int64, error)
	UpdateUser(ctx context.Context, user User)(error)
	GetUserById(ctx context.Context, ID int)(User, error)
	GetUserByMail(ctx context.Context, mail string)(User, error)
	CreateUserPassword(ctx context.Context, password string, user User)(error)
	UpdateUserPassword(ctx context.Context, oldPassword string, newPassword string, user User)(error)
	// ComparePassword(ctx context.Context, password string)(error)//interne
	// SetUserCookie(ctx context.Context, user User)(error) //interne
	// RemoveUserCookie(ctx context.Context, user User)(error) //interne
	LoginUser(ctx context.Context, email string, password string)(error)
	LogoutUser(ctx context.Context, user User)(error)
}