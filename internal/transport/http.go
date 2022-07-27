package transport

import (
	"fmt"
	"github.com/coreyvan/go-address/internal/service"
	"go.uber.org/zap"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

type HTTP struct {
	host    string
	port    int
	log     *zap.SugaredLogger
	mux     *http.ServeMux
	service service.Service
}

type HTTPConfig struct {
	Host string
	Port int
}

func NewHTTPTransport(cfg HTTPConfig, log *zap.SugaredLogger, srv service.Service) *HTTP {
	return &HTTP{
		host:    cfg.Host,
		port:    cfg.Port,
		mux:     http.NewServeMux(),
		log:     log,
		service: srv,
	}
}

func (h *HTTP) Handle(pattern string, handler Handler) {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			h.log.Errorw("error from handler", "error", err)
		}
	}

	h.mux.HandleFunc(pattern, wrapped)
}

func (h *HTTP) Listen() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", h.host, h.port), h.mux)
}
