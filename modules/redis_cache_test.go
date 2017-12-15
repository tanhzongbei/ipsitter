package modules

import (
	"gopkg.in/redis.v5"
	"testing"
	"fmt"
)

func TestSetIPCache(t *testing.T)  {
	redisOpt, err := redis.ParseURL("redis://10.10.28.2:6379/0")
	if err != nil {
		panic(err)
	}
	CacheConnection = redis.NewClient(redisOpt)

	var ip_test string
	ip_test = "14.152.49.250"
	SetIPCache(ip_test, "\u4e2d\u56fd|\u5e7f\u4e1c|\u4f5b\u5c71")

	lbsInfo, _ := GetIPCache(ip_test)
	fmt.Printf("%s", lbsInfo)
}
