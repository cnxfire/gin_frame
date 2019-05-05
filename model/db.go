package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_frame/config"
	"os"
	"time"
)

// DB 数据库连接
var DB *gorm.DB

// RedisPool Redis连接池
var RedisPool *redis.Pool

func initDB() {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		setting.Config.Database.User, setting.Config.Database.Password, setting.Config.Database.Host, setting.Config.Database.Port, setting.Config.Database.Database, setting.Config.Database.Charset) //拼接URL
	db, err := gorm.Open(setting.Config.Database.Dialect, url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if setting.Dev == setting.Config.RunMode {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(setting.Config.Database.MaxidleConns)
	db.DB().SetMaxOpenConns(setting.Config.Database.MaxopenConns)
	DB = db
	if setting.Config.RunMode =="dev" {

	}

}
func initRedis() {
	url := fmt.Sprintf("%s:%d", setting.Config.RedisConfig.Host, setting.Config.RedisConfig.Port) //拼接URL
	setting.Config.RedisConfig.Url = url
	RedisPool = &redis.Pool{
		MaxIdle:     setting.Config.RedisConfig.MaxIdle,
		MaxActive:   setting.Config.RedisConfig.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.Config.RedisConfig.Url, redis.DialPassword(setting.Config.RedisConfig.Password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	initDB()
	initRedis()
}
