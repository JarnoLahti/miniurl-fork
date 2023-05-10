package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slog"
	"net/http"
)

type AddUrlReq struct {
	Url string `json:"url"`
}

type AddUrlRes struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type Handler interface {
	AddUrl(url string) (hash string, err error)
}

type API struct {
	handler Handler
}

func Bind(r *httprouter.Router, h Handler) {
	a := &API{handler: h}
	r.POST("/api/v1/url", a.AddUrl)
}

func (a *API) AddUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var b AddUrlReq

	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		//TODO
		return
	}

	hash, err := a.handler.AddUrl(b.Url)

	if err != nil {
		//TODO
		return
	}

	err = json.NewEncoder(w).Encode(AddUrlRes{
		Url:  b.Url,
		Hash: hash,
	})

	if err != nil {
		slog.Error(err.Error())
	}
}
