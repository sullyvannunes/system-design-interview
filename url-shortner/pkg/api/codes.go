package api

import (
	"fmt"
	"net/http"
)

func (srv Server) HandleGetCode(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	code := queries.Get("code")
	w.Header().Add("Content-Type", "application/json")

	entity, err := srv.Repository.Find(code)
	if err != nil {
		w.WriteHeader(404)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"error_message\": \"code not found\" }"))
		return
	}

	err = srv.Repository.AddVisit(code)
	if err != nil {
		w.WriteHeader(422)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"error_message\": \"invalid counter\" }"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("{\"code\": \"%s\", \"created_at\": %s}", entity.URL, entity.CreatedAt)))
}
