package svc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"zero-tiktok/service/user/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	DBList      *DBList
	AsynqClient *asynq.Client
}

type DBList struct {
	Mysql *gorm.DB
	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		DBList:      initDB(c),
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{Addr: c.DBList.Redis.Address, Password: c.DBList.Redis.Password}),
	}
}

func initDB(c config.Config) *DBList {
	dbList := new(DBList)
	dbList.Mysql = initMysql(c)
	//dbList.Redis = initRedis(c)

	return dbList
}

func initMysql(c config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBList.Mysql.Username,
		c.DBList.Mysql.Password,
		c.DBList.Mysql.Address,
		c.DBList.Mysql.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.DBList.Mysql.TablePrefix, // 表名前缀
			SingularTable: true,                       // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// 自动建表
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err)
	//}

	return db
}

func initRedis(c config.Config) *redis.Client {
	fmt.Println("connect Redis ...")
	db := redis.NewClient(&redis.Options{
		Addr:     c.DBList.Redis.Address,
		Password: c.DBList.Redis.Password,
		//DB:       c.DBList.Redis.DB,
		//超时
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  3 * time.Second,
	})
	_, err := db.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect Redis failed")
		panic(err)
	}
	fmt.Println("connect Redis success")
	return db
}
