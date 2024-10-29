package cache

import (
	"github.com/go-redis/redis/v8"
)

var (
	RedisCli *redis.Client
)

const (
	KeyNamePrefixHash = "hash:"
)

func init() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:7429",
		Password: "G^&*H(UIONUBXRDFGndu",
		DB:       0,
	})
}
