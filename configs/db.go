package configs

type PostgresConfig struct {
	Adress string `yaml:"address"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	DBName string `yaml:"dbname"`
	Port   string `yaml:"port"`
}
