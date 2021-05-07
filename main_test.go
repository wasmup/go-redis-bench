package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/mediocregopher/radix/v3"
)

func BenchmarkRadix3(b *testing.B) {
	rdb, err := radix.Dial("tcp", "localhost:6379")
	if err != nil {
		b.Fatal(err)
	}
	defer rdb.Close()

	err = radix.Cmd(nil, "DEL", "set1").Run(rdb)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		radix.FlatCmd(nil, "SADD", "set1", fmt.Sprint(i)).Run(rdb)
		if err != nil {
			log.Fatal(err)
		}
	}

	var n int64
	mn := radix.MaybeNil{Rcv: &n}
	err = radix.Cmd(&mn, "SCARD", "set1").Run(rdb)
	if err != nil {
		log.Fatal(err)
	} else if mn.Nil {
		fmt.Println("rcv is nil")
	} else {
		if int(n) != b.N {
			b.Error("n")
		}
	}
}

func BenchmarkRedigo(b *testing.B) {
	rdb, err := redigo.Dial("tcp", "localhost:6379")
	if err != nil {
		b.Fatal(err)
	}
	defer rdb.Close()

	err = rdb.Send("DEL", "set1")
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rdb.Send("SADD", "set1", i)
		if err != nil {
			log.Fatal(err)
		}
	}

	n, err := redigo.Int64(rdb.Do("SCARD", "set1"))
	if err != nil {
		log.Fatal(err)
	}
	if int(n) != b.N {
		b.Error("n")
	}
}

func BenchmarkGoRedis(b *testing.B) {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	var cmd = rdb.Del(ctx, "set1")
	cmd.Result()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rdb.SAdd(ctx, "set1", i).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	cmd = rdb.SCard(ctx, "set1")
	n, err := cmd.Result()
	if err != nil {
		log.Fatal(err)
	}
	if int(n) != b.N {
		b.Error("n")
	}
}
