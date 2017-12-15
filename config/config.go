package config

import (
"ipsitter/slutils"
)


var (
	RedisCluster     string		//redis
	SentryDsn        string     //sentry
	HttpTimeout		 int	//http timeout
)

const IP_REDIS_FIELD = "IP_FIELD"

func InitConfig() {
	RedisCluster = slutils.GetSingle("redis")
	SentryDsn = slutils.GetSingle("sentry_dsn")
	HttpTimeout = 1
}
