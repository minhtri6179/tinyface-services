package api

import (
	"github.com/minhtri6179/tinyface-users-service/store"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"honnef.co/go/tools/config"
)

type Server struct {
	redisDatabase *redis.Client
	store         store.Store
	config        *config.Config
}

func NewServer(cfg *config.Config, rdb *redis.Client, db *gorm.DB) *Server {
	server := &Server{redisDatabase: rdb}
	server.store = store.NewStore(db)
	server.config = cfg
	return server
}
