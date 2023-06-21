package db

import (
	"github.com/sullyvannunes/url-shortner/pkg/short_url"
)

type ShortURLRepositoryService struct {
	Client *DBClient
}

func (rp ShortURLRepositoryService) Find(code string) (short_url.URLEntity, error) {
	return rp.Client.Find(code)
}

func (rp ShortURLRepositoryService) Store(entity short_url.URLEntity) error {
	return rp.Client.Store(entity)
}

func (rp ShortURLRepositoryService) AddVisit(code string) error {
	return rp.Client.AddVisit(code)
}

func NewShortURLRepository() ShortURLRepositoryService {
	return ShortURLRepositoryService{
		Client: NewDBClient(),
	}
}
