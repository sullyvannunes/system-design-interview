package shortened_url

import "time"

type URLEntity struct {
	Code      string    `redis:"code"`
	URL       string    `redis:"url"`
	CreatedAt time.Time `redis:"created_at"`
}

type URLVisit int
