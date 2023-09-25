package configs

import (
	"gorm.io/driver/postgres"
)

type AuthenConfig struct {
	Port       int             `yaml:"port"`
	PostgreSQL postgres.Config `yaml:"postgresql"`
}
