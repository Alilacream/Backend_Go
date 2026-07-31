package api

import (
	"alilacream/jwt/handler"
	"alilacream/jwt/model"
	"net/http"
)

var users *map[string]model.User

type Server struct {
	addr string
	mux  *http.ServeMux
}
type Api struct {
	srv Server
}

func NewServerApi(addr string, mux *http.ServeMux) *Api {
	return &Api{
		Server{
			addr: addr,
			mux:  mux,
		},
	}
}

func (a *Api) Routes() {
	router := a.srv.mux
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the server once again"))
	})
	router.HandleFunc("POST /register", handler.Register(users))
	router.HandleFunc("POST /login", handler.Login(users))
}

func (a *Api) Run() error {
	// would cause an error
	return http.ListenAndServe(a.srv.addr, a.srv.mux)
}
