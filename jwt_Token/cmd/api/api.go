package api

import (
	"net/http"

	"alilacream/jwt/handler"
	"alilacream/jwt/model"
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
	a.srv.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the server once again"))
	})
	a.srv.mux.HandleFunc("/register", handler.Register(users))
	a.srv.mux.HandleFunc("/login", handler.Login(users))
}

func (a *Api) Run() error {
	// would cause an error
	return http.ListenAndServe(a.srv.addr, a.srv.mux)
}
