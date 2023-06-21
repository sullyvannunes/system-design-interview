package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sullyvannunes/url-shortner/pkg/short_url"
)

type DBClient struct {
	Client *redis.Client
}

func NewDBClient() *DBClient {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
		},
	)

	fmt.Println("Redis client setup is done")

	return &DBClient{
		Client: client,
	}
}

func (rdb DBClient) Find(code string) (short_url.URLEntity, error) {
	ctx := context.Background()
	urlEntityKey := "entity:" + code
	var entity short_url.URLEntity
	err := rdb.Client.HGetAll(ctx, urlEntityKey).Scan(&entity)
	if err != nil {
		return short_url.URLEntity{}, err
	}

	return entity, nil
}

func (rdb DBClient) Store(urlEntity short_url.URLEntity) error {
	ctx := context.Background()
	urlEntityKey := "entity:" + urlEntity.ShortURL
	return rdb.Client.HSet(ctx, urlEntityKey, urlEntity).Err()
}

func (rdb DBClient) AddVisit(code string) error {
	ctx := context.Background()
	urlVisityKey := "visit:" + code
	return rdb.Client.Incr(ctx, urlVisityKey).Err()
}
