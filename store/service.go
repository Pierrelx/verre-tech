package store

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

//Service correspond au service de magasin
type Service interface {
	Create(ctx context.Context, store Store) (int64, error)
	GetByID(ctx context.Context, id int) (Store, error)
	UpdateStore(ctx context.Context, store Store) (Store, error)
	DeleteStore(ctx context.Context, id int) error
}
