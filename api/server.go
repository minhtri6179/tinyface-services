package api

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	redisDatabase *redis.Client
	router        *gin.Engine
}

func NewServer(rdb *redis.Client) *Server {
	server := &Server{redisDatabase: rdb}
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", server.Pong)
	}
	server.router = router
	return server
}

func (s *Server) StartServer(address string) error {
	return s.router.Run(address)
}
