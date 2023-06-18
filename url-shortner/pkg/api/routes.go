package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *Server) SetupRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/codes", srv.HandleGetCode).Methods(http.MethodGet)
	router.HandleFunc("/", srv.HandleGenerateURLEntity).Methods(http.MethodPost)

	srv.Router = router
}
