package modules

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Sina_LBS struct {
	Result  uint32	`json:"ret"`
	Provice string	`json:"province"`
	Country string	`json:"country"`
	City	string	`json:"city"`
}


func QueryLBSByIpFromBaidu(ip string) (lbsInfo string, err error){
	url := fmt.Sprintf("http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=json&ip=%s", ip)
	res, err := httpClient.Get(url)
	if err != nil {
		return "", err
	} else {
		var data_bytes []byte
		data_bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		} else {
			var sina_res = &Sina_LBS{}
			err = json.Unmarshal(data_bytes, sina_res)
			lbsInfo = fmt.Sprintf("%s|%s|%s", sina_res.Country, sina_res.Provice, sina_res.City)
			return lbsInfo, err
		}
	}
}