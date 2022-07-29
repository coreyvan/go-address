package transport

import (
	"errors"
	"fmt"
	"github.com/coreyvan/go-address/internal/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

type HTTP struct {
	host    string
	port    int
	log     *zap.SugaredLogger
	mux     *mux.Router
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
		mux:     mux.NewRouter(),
		log:     log,
		service: srv,
	}
}

func (h *HTTP) Handle(method string, pattern string, handler Handler, queries ...string) {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			h.log.Errorw("error from handler", "error", err)
			h.handleError(w, err)
		}
	}

	h.mux.HandleFunc(pattern, wrapped).Methods(method).Queries(queries...)
}

func (h *HTTP) Listen() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", h.host, h.port), h.mux)
}

func (h *HTTP) handleRequest(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if _, err := fmt.Fprint(w, data); err != nil {
		h.log.Errorw("writing to client", "error", err)
	}
}

func (h *HTTP) handleError(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	message := "Internal server error"

	var n *service.NotFoundError
	if errors.As(err, &n); err != nil {
		code = http.StatusNotFound
		message = "Resource not found"
	}

	w.WriteHeader(code)
	if _, pErr := fmt.Fprint(w, message); pErr != nil {
		h.log.Errorw("error writing to client", "error", pErr)
	}
}
