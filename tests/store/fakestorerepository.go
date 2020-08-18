package main

import (
	"context"
	"time"

	store "github.com/Pierrelx/verre-tech/store"
)

//FakeStoreRepository - Faux repo pour requêter une liste de magasins en local au lieu de la bdd
type FakeStoreRepository struct {
	Stores []store.Store
}

//NewFakeService - Retourne une nouvelle instance initialisée avec une liste vide
func NewFakeService() (*FakeStoreRepository, error) {
	return &FakeStoreRepository{
		Stores: make([]store.Store, 0),
	}, nil
}

//LastInsertId - Retourne le plus grand ID (def: 0)
func (repo *FakeStoreRepository) LastInsertId() int {
	var maxID = 0
	for _, st := range repo.Stores {
		if st.ID > maxID {
			maxID = st.ID
		}
	}
	return maxID
}

func (repo *FakeStoreRepository) CreateStore(ctx context.Context, store store.Store) (int64, error) {
	store.CreatedOn = time.Now().Unix()
	store.UpdatedOn = time.Now().Unix()
	var nextID int = repo.LastInsertId() + 1
	store.ID = nextID
	repo.Stores = append(repo.Stores, store)
	var result int64 = int64(nextID)
	return result, nil
}

func (repo *FakeStoreRepository) GetStoreByID(ctx context.Context, id int) (store.Store, error) {
	var store store.Store
	for _, s := range repo.Stores {
		if s.ID == id {
			store = s
		}
	}
	return store, nil
}

func (repo *FakeStoreRepository) UpdateStore(ctx context.Context, store store.Store) (store.Store, error) {
	store.UpdatedOn = time.Now().Unix()
	for i, s := range repo.Stores {
		if s.ID == store.ID {
			repo.Stores[i] = store
		}
	}
	return store, nil
}

func (repo *FakeStoreRepository) DeleteStore(ctx context.Context, id int) error {
	for i, s := range repo.Stores {
		if s.ID == id {
			repo.Stores = append(repo.Stores[:i], repo.Stores[i+1:]...)
		}
	}
	return nil
}

func (repo *FakeStoreRepository) GetAll(ctx context.Context) ([]*store.Store, error) {
	var stores = make([]*store.Store, 0)
	for _, s := range repo.Stores {
		stores = append(stores, &s)
	}
	return stores, nil
}
