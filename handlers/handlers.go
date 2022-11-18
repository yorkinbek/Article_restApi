package handlers

import (
	"github/yorqinbek/CRUD/article/config"
	"github/yorqinbek/CRUD/article/storage"
)

// Handler ...
type handler struct {
	Stg storage.StorageI
	Cfg config.Config
}

// NewHandler ...
func NewHandler(stg storage.StorageI, cfg config.Config) handler {
	return handler{
		Stg: stg,
		Cfg: cfg,
	}
}
