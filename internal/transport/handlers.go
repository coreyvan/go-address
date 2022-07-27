package transport

import (
	"fmt"
	"net/http"
)

func (h *HTTP) HandleAddress() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		switch r.Method {
		case http.MethodPost:
			//	create handler for creating an address
			panic("implement me")
		case http.MethodGet:
			// create handler for getting an address
			panic("implement me")
		default:
			h.handleBadRequest(w, fmt.Sprintf("Invalid method: %s", r.Method))
			return nil
		}
	}
}

func (h *HTTP) HandleRoot() Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		panic("implement me")
	}
}

func (h *HTTP) handleBadRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	if _, err := fmt.Fprint(w, message); err != nil {
		h.log.Errorw("writing to client", "error", err)
	}
}
