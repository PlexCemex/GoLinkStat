package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projects/GoLinkStat/configs"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return &Db{
		DB: db,
	}
}
