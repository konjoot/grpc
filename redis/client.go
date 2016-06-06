package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func init() {
	if pool != nil {
		return
	}

	pool = newPool("localhost:6379")
}

func New() redis.Conn {
	return pool.Get()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
