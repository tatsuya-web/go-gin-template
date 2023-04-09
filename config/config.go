package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env        string `env:"APP_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"APP_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"APP_DB_PORT" envDefault:"3306"`
	DBUser     string `env:"APP_DB_USER" envDefault:"db_user"`
	DBPassword string `env:"APP_DB_PASSWORD" envDefault:"db_password"`
	DBName     string `env:"APP_DB_NAME" envDefault:"go_gin_template"`
	RedisHost  string `env:"APP_REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort  int    `env:"APP_REDIS_PORT" envDefault:"36379"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
