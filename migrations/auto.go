package main

import (
	"log"
	"os"
	"path/filepath"
	"projects/GoLinkStat/internal/link"
	"projects/GoLinkStat/internal/user"
	"runtime"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Не удается определить путь к текущему файлу")
	}
	projectRoot := filepath.Dir(filepath.Dir(filename))
	envPath := filepath.Join(projectRoot, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		panic(err.Error())
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&link.Link{}, &user.User{})
}
