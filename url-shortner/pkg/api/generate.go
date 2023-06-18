package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sullyvannunes/url-shortner/pkg/shortened_url"
	"github.com/teris-io/shortid"
)

func (srv Server) HandleGenerateURLEntity(w http.ResponseWriter, r *http.Request) {
	type errorMessage struct {
		ErrorMessage string `json:"error_message"`
	}

	body := struct {
		URL string `json:"url"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.URL == "" {
		invalidResponse := errorMessage{
			ErrorMessage: "Invalid payload",
		}

		content, err := json.Marshal(invalidResponse)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("internal server error"))
			return
		}

		w.WriteHeader(422)
		w.Write(content)
		w.Header().Add("Content-Type", "application/json")

		return
	}

	entity := shortened_url.URLEntity{
		Code:      shortid.MustGenerate(),
		URL:       body.URL,
		CreatedAt: time.Now(),
	}

	err = srv.Repository.Store(entity)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	b, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(201)
	w.Write(b)
}
