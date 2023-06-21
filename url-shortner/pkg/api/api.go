package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sullyvannunes/url-shortner/pkg/db"
	"github.com/sullyvannunes/url-shortner/pkg/short_url"
)

type UnprocessableEntityResponseBody struct {
	ErrorMessage string `json:"error_message"`
}

type Server struct {
	Router             http.Handler
	Timeout            time.Duration
	ShortURLRepository short_url.Repository
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

	srv.ShortURLRepository = db.NewShortURLRepository()
	srv.SetupRoutes()

	return srv
}
