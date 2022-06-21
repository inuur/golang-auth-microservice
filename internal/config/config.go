package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `env:"IS_DEBUG" env-default:"true"`
	Listen  struct {
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8080"`
	}
	MongoDB struct {
		Username string `env:"MONGO_USERNAME"`
		Password string `env:"MONGO_PASSWORD"`
		Host     string `env:"MONGO_HOST" env-required:"true"`
		Port     string `env:"MONGO_PORT" env-required:"true"`
		Database string `env:"MONGO_DATABASE" env-required:"true"`
	}
	JWT struct {
		SecretKey           string `env:"JWT_SECRET"`
		AccessTokenExpTime  int    `env:"ACCESS_TOKEN_EXP_MINUTES" env-default:"15"`
		RefreshTokenExpTime int    `env:"REFRESH_TOKEN_EXP_HOURS" env-default:"48"`
	}
	Bcrypt struct {
		Cost string `env:"BCRYPT_COST" env-default:"10"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	log.Println(instance)
	return instance
}
