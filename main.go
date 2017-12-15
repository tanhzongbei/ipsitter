package main

import (
	"os"
	"flag"
	"runtime"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/getsentry/raven-go"
	log "github.com/alecthomas/log4go"

	"ipsitter/apis"
	"ipsitter/slutils"
	"ipsitter/config"
	"ipsitter/modules"
)

const log_config  = "logging.xml"

var (
	workerNum int
	releaseMode bool
	address string
	consulAddress string
	consulPrefix string
)

func init() {
	cpuNum := runtime.NumCPU()
	flag.IntVar(&workerNum, "worker", cpuNum, "runtime MAXPROCS value")
	flag.BoolVar(&releaseMode, "release", false, "gin mode")
	flag.StringVar(&address, "address", "127.0.0.1:8090", "server address")
	flag.StringVar(&consulAddress, "ca", "", "consul address")
	flag.StringVar(&consulPrefix, "cp", "", "consul prefix")

	log.LoadConfiguration(log_config)
}

func main() {
	flag.Parse()
	log.Info("address: %s MAXPROCS:%d release:%t", address, workerNum, releaseMode)
	runtime.GOMAXPROCS(workerNum)

	// 连接consul
	slutils.InitConsul(consulAddress, consulPrefix)
	// 获取consul配置
	config.InitConfig()
	// 连接数据库和缓存
	modules.InitConnection()

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(slutils.Prometheus())
	router.Use(slutils.Logger(slutils.DefaultMetricPath))

	sentryClient, err := raven.New(config.SentryDsn)
	if err != nil {
		panic(err)
	}
	router.Use(slutils.Recovery(sentryClient))

	// 监控接口
	router.GET(slutils.DefaultMetricPath, slutils.LatestMetrics)

	// interfaces
	router.GET("/ipsitter", apis.QueryLBSByIP)

	endless.ListenAndServe(address, router)
	os.Exit(0)
}
