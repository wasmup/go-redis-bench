package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println(randStr())

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var cmd = rdb.SCard(ctx, "set1")
	fmt.Print("set1 len= ")
	fmt.Println(cmd.Result())

	cmd = rdb.Del(ctx, "set1")
	fmt.Print("del: ")
	fmt.Println(cmd.Result())

	err := rdb.SAdd(ctx, "set1", 1, 2, 3, 4).Err()
	if err != nil {
		log.Fatal(err)
	}
	cmd = rdb.SCard(ctx, "set1")
	fmt.Println(cmd.Result())

	ms := rdb.SMembers(ctx, "set1")
	a, err := ms.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func randStr() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
