package tango

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type API struct {
	mux *http.ServeMux
}

func NewAPI() *API {
	return &API{}
}

func (api *API) newMux() *http.ServeMux {
	if api.mux == nil {
		api.mux = http.NewServeMux()
	}
	return api.mux
}

func (api *API) requestHandler(resource interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		if request.ParseForm() != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var handler func(url.Values, http.Header) (int, interface{}, http.Header)

		switch request.Method {
		case "GET":
			if resource, ok := resource.(Getter); ok {
				handler = resource.Get
			}
		case "POST":
			if resource, ok := resource.(Poster); ok {
				handler = resource.Post
			}
		case "PUT":
			if resource, ok := resource.(Putter); ok {
				handler = resource.Put
			}
		case "DELETE":
			if resource, ok := resource.(Deleter); ok {
				handler = resource.Delete
			}
		}

		if handler == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		code, data, header := handler(request.Form, request.Header)

		content, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		for name, values := range header {
			for _, value := range values {
				rw.Header().Add(name, value)
			}
		}
		rw.WriteHeader(code)
		rw.Write(content)
	}
}

func (api *API) Register(resource interface{}, paths ...string) {
	for _, path := range paths {
		api.newMux().HandleFunc(path, api.requestHandler(resource))
	}
}

func (api *API) Start(port int) error {
	if api.mux == nil {
		return errors.New("请先注册资源！")
	}
	portString := fmt.Sprintf(":%d", port)
	return http.ListenAndServe(portString, api.mux)
}
