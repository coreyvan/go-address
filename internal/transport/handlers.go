package transport

import (
	"encoding/json"
	"github.com/coreyvan/go-address/internal/service"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func (h *HTTP) CreateAddress() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return service.NewInternalServerError("reading response body").Wrap(err)
		}
		//	create handler for creating an address
		var in service.CreateAddress
		if err := json.Unmarshal(bytes, &in); err != nil {
			return service.NewInternalServerError("marshaling into create request").Wrap(err)
		}

		addr, err := h.service.CreateAddress(in)
		if err != nil {
			return service.NewInternalServerError("creating address").Wrap(err)
		}

		out, err := json.Marshal(addr)
		if err != nil {
			return service.NewInternalServerError("marshaling return object")
		}

		h.handleRequest(w, http.StatusOK, string(out))
		return nil
	}
}

func (h *HTTP) GetAddress() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)

		id, ok := vars["id"]
		if !ok {
			return service.NewBadRequestError("no ID in request")
		}

		addr, err := h.service.GetAddressByID(id)
		if err != nil {
			return service.NewInternalServerError("getting address").Wrap(err)
		}

		out, err := json.Marshal(addr)
		if err != nil {
			return service.NewInternalServerError("marshaling response").Wrap(err)
		}

		h.handleRequest(w, http.StatusOK, string(out))
		return nil
	}
}

func (h *HTTP) SearchAddress() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		search := r.FormValue("search")
		if search == "" {
			return nil
		}

		addr, err := h.service.GetAddressBySearch(search)
		if err != nil {
			return service.NewInternalServerError("searching for address").Wrap(err)
		}

		out, err := json.Marshal(addr)
		if err != nil {
			return service.NewInternalServerError("marshaling response").Wrap(err)
		}

		h.handleRequest(w, http.StatusOK, string(out))
		return nil
	}
}

func (h *HTTP) HandleRoot() Handler {
	return func(writer http.ResponseWriter, r *http.Request) error {
		panic("implement me")
	}
}
