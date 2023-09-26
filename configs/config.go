package configs

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgres *PostgresConfig `yaml:"postgres"`
	Authen   *AuthenConfig   `yaml:"authen"`
	Port     int             `yaml:"port"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open AuthenticateAndPost config (path=%s) error: %s", path, err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read AuthenticateAndPost config (path=%s) error: %s", path, err)
	}

	allInOneConf := &Config{}
	if err := yaml.Unmarshal(bs, allInOneConf); err != nil {
		return nil, fmt.Errorf("unmarshal AuthenticateAndPost config (path=%s) error: %s", path, err)
	}
	return allInOneConf, nil

}
