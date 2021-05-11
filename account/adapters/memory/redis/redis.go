package redismem

import (
	"github.com/go-redis/redis"
)

type RedisMemory struct {
	conn *redis.Client
}

func NewRedisMemory(conn *redis.Client) *RedisMemory {
	return &RedisMemory{
		conn: conn,
	}
}
