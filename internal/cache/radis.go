package cache

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Host    string
	Port    string
	PW      string // Password
	DB      int
	Enabled bool
}

func NewRedisClient(addr, pw string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})
}
