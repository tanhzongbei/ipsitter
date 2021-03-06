package slutils

import (
	"fmt"
	"sync"
	"math/rand"

	consulApi "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/watch"
)

var (
	Address string
	Prefix string
	KV *consulApi.KV
	Catalog *consulApi.Catalog
)


// 监听consul配置
type WatchedParam struct {
	value string
	lock sync.RWMutex
}

func (v *WatchedParam) Get() string {
	v.lock.RLock()
	defer v.lock.RUnlock()
	return v.value
}

func (v *WatchedParam) Set(value string) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = value
}


func GetSingle(key string) string {
	pair, _, err := KV.Get(fmt.Sprintf("%s/%s", Prefix, key), nil)
	if err != nil {
		panic(err)
	}
	return string(pair.Value)
}

func WatchSingle(key string, param *WatchedParam) {
	params := map[string]interface{}{
		"type": "key",
		"key": fmt.Sprintf("%s/%s", Prefix, key),
	}
	plan, _ := consulWatch.Parse(params)
	plan.Handler = func(idx uint64, raw interface{}) {
		if raw == nil {
			return
		}

		v, ok := raw.(*consulApi.KVPair)
		if ok && v != nil {
			newValue := string(v.Value)
			param.Set(newValue)
		}
	}

	go plan.Run(Address)
}

func GetServiceAddress(serviceName string) string {
	pair, _, err := Catalog.Service(serviceName, "", nil)
	if err != nil {
		panic(err)
	}

	pairLength := len(pair)
	if pairLength == 0 {
		panic(fmt.Errorf("%s not exist", serviceName))
	}

	index := rand.Intn(pairLength)
	topService := pair[index]
	return fmt.Sprintf("%s:%d", topService.ServiceAddress, topService.ServicePort)
}

func InitConsul(address string, prefix string) {
	Address = address
	Prefix = prefix

	conConfig := consulApi.Config{Address: address}
	consulClient, err := consulApi.NewClient(&conConfig)
	if err != nil {
		panic(err)
	}

	KV = consulClient.KV()
	Catalog = consulClient.Catalog()
}
