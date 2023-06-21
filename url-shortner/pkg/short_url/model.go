package short_url

import (
	"time"
)

type URLEntity struct {
	ShortURL  string    `redis:"code"`
	LongURL   string    `redis:"url"`
	CreatedAt time.Time `redis:"created_at"`
}

type Repository interface {
	Find(code string) (URLEntity, error)
	Store(URLEntity) error
	AddVisit(code string) error
}
