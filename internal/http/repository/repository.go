package repository

import (
	"github.com/certified-juniors/AtomHackEarthBackend/internal/config"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/db"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/minio_client"
)

type Repository struct {
	db *db.Database
	mc *minio_client.Minio
}

func New(cfg *config.App) (*Repository, error) {
	var r Repository

	// Инициализация базы данных
	dbInstance := &db.Database{}
	err := dbInstance.New(&cfg.Database)
	if err != nil {
		return nil, err
	}
	r.db = dbInstance

	// Инициализация клиента MinIO
	mcInstance := &minio_client.Minio{}
	err = mcInstance.New(&cfg.Minio)
	if err != nil {
		return nil, err
	}
	r.mc = mcInstance

	return &r, nil
}
