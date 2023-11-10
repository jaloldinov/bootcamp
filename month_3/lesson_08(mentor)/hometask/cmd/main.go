package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"market/api"
	"market/api/handler"
	"market/config"
	grpc_client "market/grpc"
	"market/pkg/logger"
	"market/storage/postgres"
	"market/storage/redis"

	goRedis "github.com/redis/go-redis/v9"
)

func main() {

	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := postgres.NewStorage(context.Background(), cfg)
	redisStrg, err := redis.NewCache(context.Background(), cfg)
	if err != nil {
		return
	}
	client, err := grpc_client.New(cfg)
	if err != nil {
		log.Fatal("error while init grpc clients", logger.Any("error:", err))
	}
	h := handler.NewHandler(strg, redisStrg, log, client)

	r := api.NewServer(h)
	r.Run()

	// redisFunc()
}

var client = goRedis.NewClient(&goRedis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

type User struct {
	Name string
	Age  int
}

func redisFunc() {

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println(pong)

	// err = client.Set(context.Background(), "name", "Shohruh", 0).Err()
	// if err != nil {
	// 	log.Println("Set key:", err)
	// }
	// fmt.Println("set")
	// val, err := client.Get(context.Background(), "name").Result()
	// if err != nil {
	// 	log.Println("Get key:", err)
	// }

	// fmt.Println(val)

	// user, err := json.Marshal(
	// 	User{
	// 		Name: "Alex",
	// 		Age:  22,
	// 	},
	// )

	// if err != nil {
	// 	log.Println("marshal user:", err)
	// }

	// err = client.Set(context.Background(), "user", user, 0).Err()
	// if err != nil {
	// 	log.Println("set user:", err)
	// }

	data, err := client.Get(context.Background(), "user").Result()
	if err != nil {
		log.Println("get user:", err)
	}

	var user1 User

	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		log.Println("unmarshal data:", err)
	}

	fmt.Println(user1)
}
