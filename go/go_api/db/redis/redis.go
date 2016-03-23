package redis

import (
	"github.com/garyburd/redigo/redis"
)

type RedisDB struct {
	pool *redis.Pool
}

func NewRedisConnection() *RedisDB {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("SELECT", 0)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}

	redis := &RedisDB{pool: pool}

	return redis
}

func (r *RedisDB) Ping() string {
	_, err := r.pool.Get().Do("PING")
	if err == nil {
		return "WORKING"
	} else {
		return "NOT WORKING"
	}
}
