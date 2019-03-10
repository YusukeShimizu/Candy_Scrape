package redisdb

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func NewClient() (*Redis, error) {
	r := Redis{}
	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := r.client.Ping().Result()
	if err != nil {
		return &r, err
	}
	fmt.Println(pong, err)
	return &r, nil
}

func (r *Redis) GetSet(key, value string) error {
	stringCmd := r.client.GetSet(key, value)
	if stringCmd.Err() != nil {
		return stringCmd.Err()
	}
	return nil
}

func (r *Redis) Set(key, Value string) error {
	err := r.client.Set(key, Value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) (string, bool, error) {
	val, err := r.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
		return val, false, nil
	} else if err != nil {
		return val, false, err
	}
	return val, true, nil
}
