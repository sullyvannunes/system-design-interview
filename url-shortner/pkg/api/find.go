package api

import (
	"net/http"
	"strings"

	"github.com/sullyvannunes/url-shortner/pkg/short_url"
)

func (srv Server) HandleFindLongURLByShortURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	code := strings.Split(r.URL.EscapedPath(), "/")[1]

	longURL, err := short_url.FindLongURLByShortURL(code, srv.ShortURLRepository)

	if err != nil {
		w.WriteHeader(404)

		w.Write([]byte("{\"error_message\": \"code not found\" }"))
		return
	}

	http.Redirect(w, r, longURL, 302)
}
