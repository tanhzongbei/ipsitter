package modules

import (
	"ipsitter/config"
)

func SetIPCache(ip string, lbsInfo string) {
	CacheConnection.HSet(config.IP_REDIS_FIELD, ip, lbsInfo)
}

func GetIPCache(ip string) (lbsInfo string, err error)  {
	values := CacheConnection.HGet(config.IP_REDIS_FIELD, ip)
	err = values.Err()
	lbsInfo = values.Val()
	return
}