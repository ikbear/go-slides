package tango

import (
	"net/http"
	"net/url"
)

type Getter interface {
	Get(url.Values, http.Header) (int, interface{}, http.Header)
}

type Poster interface {
	Post(url.Values, http.Header) (int, interface{}, http.Header)
}

type Putter interface {
	Put(url.Values, http.Header) (int, interface{}, http.Header)
}

type Deleter interface {
	Delete(url.Values, http.Header) (int, interface{}, http.Header)
}
