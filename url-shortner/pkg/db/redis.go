package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sullyvannunes/url-shortner/pkg/shortened_url"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
		},
	)

	fmt.Println("Redis client setup is done")

	return &RedisClient{
		Client: client,
	}
}

func (rdb RedisClient) Find(code string) (shortened_url.URLEntity, error) {
	ctx := context.Background()
	urlEntityKey := "entity:" + code
	var entity shortened_url.URLEntity
	err := rdb.Client.HGetAll(ctx, urlEntityKey).Scan(&entity)
	if err != nil {
		return shortened_url.URLEntity{}, err
	}

	return entity, nil
}

func (rdb RedisClient) Store(urlEntity shortened_url.URLEntity) error {
	ctx := context.Background()
	urlEntityKey := "entity:" + urlEntity.Code
	return rdb.Client.HSet(ctx, urlEntityKey, urlEntity).Err()
}

func (rdb RedisClient) AddVisit(code string) error {
	ctx := context.Background()
	urlVisityKey := "visit:" + code
	return rdb.Client.Incr(ctx, urlVisityKey).Err()
}
