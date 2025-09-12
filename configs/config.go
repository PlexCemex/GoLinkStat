package configs

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	_, filename, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("Не удается определить путь к текущему файлу")
    }
    projectRoot := filepath.Dir(filepath.Dir(filename))
    envPath := filepath.Join(projectRoot, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file, using default config: ", err)
	}
	return &Config{
		Db : DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
