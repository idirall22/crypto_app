package redismem

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func (r *RedisMemory) SetLoginAttemps(key string, value int, exp time.Duration) error {
	return r.conn.Set(key, value, exp).Err()
}

func (r *RedisMemory) GetLoginAttemps(key string) (int, error) {
	result, err := r.conn.Get(key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if result == "" {
		return 0, nil
	}
	return strconv.Atoi(result)
}
