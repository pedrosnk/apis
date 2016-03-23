package redis

import (
	"github.com/garyburd/redigo/redis"
)

func New() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", d.host)
			if err != nil {
				return nil, err
			}
			if d.password != "" {
				_, err = conn.Do("AUTH", d.password)
				if err != nil {
					return nil, err
				}
			}
			_, err = conn.Do("SELECT", d.db)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}
}
