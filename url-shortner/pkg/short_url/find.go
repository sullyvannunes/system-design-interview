package short_url

import (
	"errors"
)

func FindLongURLByShortURL(shortUrl string, repo Repository) (string, error) {
	if shortUrl == "" {
		return "", errors.New("invalid shortUrl")
	}

	urlEntity, err := repo.Find(shortUrl)
	if err != nil {
		return "", err
	}

	return urlEntity.LongURL, nil
}
