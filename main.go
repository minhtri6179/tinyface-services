package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/minhtri6179/tinyface-users-service/configs"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	path = flag.String("conf", "config.yml", "config path for this service")
)

func main() {

	config, err := configs.LoadConfig(*path)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	dsn := "host=localhost user=root password=secret dbname=auth port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}
	s := gapi.NewServer(config, rdb, db)
	pb.RegisterIdentityServer(server, s)
	fmt.Println("Server is running on port 9000")
	if err := server.Serve(conn); err != nil {
		panic(err)
	}
}
