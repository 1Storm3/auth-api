package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Host         string `env:"APP_HOST" envDefault:"localhost"`
	Port         string `env:"APP_PORT" envDefault:"8080"`
	JwtSecretKey string `env:"JWT_SECRET_KEY" env-required:"true"`
	JwtExpiresIn string `env:"JWT_EXPIRES_IN" env-default:"24h"`
}

type DBConfig struct {
	URL string `env:"DATABASE_URL" env-required:"true"`
}

func MustLoad() *Config {
	var cfg Config
	var err error

	configPath := fetchConfigPath()
	if configPath != "" {
		err = godotenv.Load(configPath)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Println("Не удалось загрузить конфигурацию")
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("Конфигурация некорректна:" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Путь к файлу конфигурации")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
