package app

import (
	"github.com/certified-juniors/AtomHackEarthBackend/internal/config"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/http/handler"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/http/repository"
)

type Application struct {
	cfg     *config.App
	handler *handler.Handler
}

func New() (*Application, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	repo, err := repository.New(cfg)
	if err != nil {
		return nil, err
	}

	h := handler.New(repo, cfg)

	app := &Application{
		cfg:     cfg,
		handler: h,
	}

	return app, nil
}
