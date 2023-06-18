package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sullyvannunes/url-shortner/pkg/repository"
)

type Server struct {
	Router     http.Handler
	Timeout    time.Duration
	Repository repository.Repository
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func StartServer(srv *Server) error {
	fmt.Println("Listening to port 3000")
	err := http.ListenAndServe(":3000", srv)
	if err != nil {
		panic(err)
	}
	return nil
}

func NewServer() *Server {
	srv := &Server{
		Timeout: 5 * time.Second,
	}

	srv.Repository = repository.NewRepository()
	srv.SetupRoutes()

	return srv
}
