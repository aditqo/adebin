package store

import (
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address string
	Password string
	Database int
}

type Client struct {
	conn *redis.Client
}

func NewClient(c Config) Client {

	return Client{
		conn: redis.NewClient(&redis.Options{
			Addr: c.Address,
			Password: c.Password,
			DB: c.Database,
		}),
	}

}
