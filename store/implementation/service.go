package implementation

import (
	"context"
	"database/sql"
	"time"

	storesvc "github.com/PierreLx/verre-tech-ms/store"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository storesvc.Repository
	logger     log.Logger
}

// NewService créé et retourne une nouvelle instannce du StoreService
func NewService(rep storesvc.Repository, logger log.Logger) storesvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, store storesvc.Store) (int64, error) {
	logger := log.With(s.logger, "method", "Create")
	store.CreatedOn = time.Now().Unix()
	var id int64
	if id, err := s.repository.CreateStore(ctx, store); err != nil {
		level.Error(logger).Log("err", err)
		return id, storesvc.ErrCmdRepository
	}
	return id, nil
}

func (s *service) GetByID(ctx context.Context, id int) (storesvc.Store, error) {
	logger := log.With(s.logger, "method", "GetByID")
	store, err := s.repository.GetStoreByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		if err == sql.ErrNoRows {
			return store, storesvc.ErrStoreNotFound
		}
		return store, storesvc.ErrQueryRepository
	}
	return store, nil
}

func (s *service) UpdateStore(ctx context.Context, store storesvc.Store) (storesvc.Store, error) {
	logger := log.With(s.logger, "method", "Update")
	store, err := s.repository.UpdateStore(ctx, store)
	if err != nil {
		level.Error(logger).Log("err", err)
		if err == sql.ErrNoRows {
			return store, storesvc.ErrStoreNotFound
		}
		return store, storesvc.ErrQueryRepository
	}
	return store, nil
}

func (s *service) DeleteStore(ctx context.Context, id int) error {
	logger := log.With(s.logger, "method", "Delete")
	err := s.repository.DeleteStore(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		if err == sql.ErrNoRows {
			return storesvc.ErrStoreNotFound
		}
		return storesvc.ErrQueryRepository
	}
	return nil
}
