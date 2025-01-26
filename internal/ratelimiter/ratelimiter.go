package ratelimiter

import "time"

type Limiter interface {
	Allow(ip string) (bool, time.Duration)
}
