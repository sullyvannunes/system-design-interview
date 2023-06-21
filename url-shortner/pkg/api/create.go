package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sullyvannunes/url-shortner/pkg/short_url"
)

type CreateRequest struct {
	URL string `json:"url"`
}

func (srv Server) HandleGenerateURLEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request CreateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(422)
		response := UnprocessableEntityResponseBody{
			ErrorMessage: "invalid payload",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	shortUrl, err := short_url.Create(request.URL, srv.ShortURLRepository)
	if err != nil {
		errMessage := fmt.Errorf("[%v] error creating new short url %w\n", time.Now(), err)
		fmt.Println(errMessage)
		w.WriteHeader(422)
		response := UnprocessableEntityResponseBody{
			ErrorMessage: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortUrl,
	})

}
