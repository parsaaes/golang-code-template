package redis

import (
	"github.com/go-redis/redis"
	"github.com/rashadansari/golang-code-template/config"
)

func Create(c config.RedisConfig) (client redis.Cmdable, closeFunc func() error) {
	result := redis.NewClient(
		&redis.Options{
			Addr:            c.Address,
			PoolSize:        c.PoolSize,
			DialTimeout:     c.DialTimeout,
			ReadTimeout:     c.ReadTimeout,
			WriteTimeout:    c.WriteTimeout,
			PoolTimeout:     c.PoolTimeout,
			IdleTimeout:     c.IdleTimeout,
			MinIdleConns:    c.MinIdleConns,
			MaxRetries:      c.MaxRetries,
			MinRetryBackoff: c.MinRetryBackoff,
			MaxRetryBackoff: c.MaxRetryBackoff,
		},
	)

	return result, result.Close
}
