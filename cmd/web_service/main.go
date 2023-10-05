package main

import (
	"log"

	"github.com/minhtri6179/tinyface-users-service/api"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	server := api.NewServer(rdb)
	err := server.StartServer(":8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
