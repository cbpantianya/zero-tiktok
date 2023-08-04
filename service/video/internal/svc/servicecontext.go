package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zero-tiktok/service/video/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewDBConn(c config.Config) *gorm.DB {
	fmt.Println(c.DSN)
	db, err := gorm.Open(mysql.Open(c.DSN))
	if err != nil {
		panic(err)
	}

	return db
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     NewDBConn(c),
	}
}
