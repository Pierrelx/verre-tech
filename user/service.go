package user

import (
	"context"
	"errors"
)

var (
	//ErrStoreNotFound => store non trouvé
	ErrStoreNotFound = errors.New("store not found")
	//ErrCmdRepository => erreur système
	ErrCmdRepository = errors.New("unable to command repository")
	//ErrQueryRepository => erreur db
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface{
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