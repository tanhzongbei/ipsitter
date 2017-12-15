package modules

import (
	"time"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/redis.v5"
	"net/http"
	"ipsitter/config"
)

var (
	//DbConnection *gorm.DB
	CacheConnection *redis.Client
	httpClient *http.Client
)

func InitConnection() {
	//var err error
	//db, err := gorm.Open("mysql", config)
	//if err != nil {
	//	panic(err)
	//}
	//DbConnection = db

	redisOpt, err := redis.ParseURL(config.RedisCluster)
	if err != nil {
		panic(err)
	}
	CacheConnection = redis.NewClient(redisOpt)
	httpClient = &http.Client{Timeout: time.Second * time.Duration(config.HttpTimeout)}
}