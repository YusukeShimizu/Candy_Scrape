package redisdb

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	redis, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	err = redis.client.Set("test", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}
	val, err := redis.client.Get("test").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("test", val)
	err = redis.client.Del("test").Err()
	if err != nil {
		t.Fatal(err)
	}
	intCmd := redis.client.Exists("test")
	fmt.Println("test", intCmd.Val())
}

func TestGet(t *testing.T) {
	redis, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	val, bool, err := redis.Get("test")
	fmt.Println(val, bool, err)
}
