package redismem_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-redis/redis"
	imemory "github.com/idirall22/crypto_app/account/adapters/memory"
	redismem "github.com/idirall22/crypto_app/account/adapters/memory/redis"
	"github.com/idirall22/crypto_app/account/config"
)

var (
	memTest imemory.IMemoryStore
	rdbTest *redis.Client
)

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../../../../")
	if err != nil {
		panic(err)
	}

	conn := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})
	defer conn.Close()

	rdbTest = conn
	err = rdbTest.Ping().Err()
	if err != nil {
		panic(err)
	}
	memTest = redismem.NewRedisMemory(conn)

	os.Exit(m.Run())
}
