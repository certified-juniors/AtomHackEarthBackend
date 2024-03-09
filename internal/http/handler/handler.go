package handler

import (
	"github.com/certified-juniors/AtomHackEarthBackend/internal/config"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/http/repository"
)

type Handler struct {
	r *repository.Repository
}

func New(repo *repository.Repository, config *config.App) *Handler {
	return &Handler{
		r: repo,
	}
}
