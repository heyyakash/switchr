package cache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb *redis.Client
}

var Redisdb RedisClient

func (r *RedisClient) ConnectRedis() {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	testString := "Hello"
	ctx := context.Background()
	err := r.rdb.Set(ctx, testString, testString, 0).Err()
	if err != nil {
		log.Fatal("Redis Error : ", err)
	}
	val, err := r.rdb.Get(ctx, testString).Result()
	if err != nil {
		log.Fatal("Redis Error while Get: ", err)
	}
	if val == testString {
		log.Print("Redis Connected")
	} else {
		log.Fatal("Redis Error : val and teststring don't match", testString, testString)
	}
}

func (r *RedisClient) Set(key string, value interface{}) error {
	ctx := context.Background()
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = r.rdb.Set(ctx, key, val, 0).Err()
	return err
}

func (r *RedisClient) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.rdb.Get(ctx, key).Result()
	return val, err
}

func (r *RedisClient) Del(key string) {
	ctx := context.Background()
	err := r.rdb.Del(ctx, key).Err()
	log.Print(err)
}
