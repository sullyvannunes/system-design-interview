package short_url

import (
	"errors"

	"github.com/teris-io/shortid"
)

func Create(url string, repo Repository) (string, error) {
	if url == "" {
		return "", errors.New("url must not be empty")
	}

	entity := URLEntity{
		LongURL:  url,
		ShortURL: shortid.MustGenerate(),
	}

	err := repo.Store(entity)
	if err != nil {
		return "", errors.New("something went wrong. try contact administrator")
	}

	return entity.ShortURL, nil
}
