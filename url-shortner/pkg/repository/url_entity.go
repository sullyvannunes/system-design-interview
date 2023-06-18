package repository

import (
	"github.com/sullyvannunes/url-shortner/pkg/db"
	"github.com/sullyvannunes/url-shortner/pkg/shortened_url"
)

type Repository interface {
	Find(code string) (shortened_url.URLEntity, error)
	Store(shortened_url.URLEntity) error
	AddVisit(code string) error
}

type RepositoryService struct {
	Client *db.RedisClient
}

func (rp RepositoryService) Find(code string) (shortened_url.URLEntity, error) {
	return rp.Client.Find(code)
}

func (rp RepositoryService) Store(entity shortened_url.URLEntity) error {
	return rp.Client.Store(entity)
}

func (rp RepositoryService) AddVisit(code string) error {
	return rp.Client.AddVisit(code)
}

func NewRepository() RepositoryService {
	return RepositoryService{
		Client: db.NewRedisClient(),
	}
}
