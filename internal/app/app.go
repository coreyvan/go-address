package app

import (
	"github.com/coreyvan/go-address/internal/service"
	"github.com/coreyvan/go-address/internal/storage"
	"github.com/coreyvan/go-address/internal/transport"
	"go.uber.org/zap"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func Run(cfg Config, log *zap.SugaredLogger) error {
	log.Infow("app running...")

	tCfg := transport.HTTPConfig{
		Host: cfg.Host,
		Port: cfg.Port,
	}

	st, err := storage.CreatePostgresStorage(storage.PostgresConfig{Conn: cfg.Conn})
	if err != nil {
		return err
	}

	srv := service.NewService(st)

	t := transport.NewHTTPTransport(tCfg, log, srv)
	routes(t)

	log.Infof("listening on port %d...", cfg.Port)
	return t.Listen()
}

func routes(s *transport.HTTP) {
	s.Handle(http.MethodGet, "/", s.HandleRoot())
	s.Handle(http.MethodPost, "/address", s.CreateAddress())
	s.Handle(http.MethodGet, "/address/{id}", s.GetAddress())
	s.Handle(http.MethodGet, "/address", s.SearchAddress(), "search", "{search}")
}
