package cache

import "github.com/redis/go-redis/v9"

type Storage struct {
	Users users
}

type users interface {
	Get(string) (any, error)
	Set(any) error
	Delete(any)
}

func NewRedisStorage(db *redis.Client) Storage {
	return Storage{}
}
