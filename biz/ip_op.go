package biz

import "ipsitter/modules"

func GetIPInfo(ip string) (lbsinfo string, err error) {
	lbsinfo, err = modules.GetIPCache(ip)
	if err == nil {
		return lbsinfo, err
	} else {
		lbsinfo, err = modules.QueryLBSByIpFromBaidu(ip)
		if err == nil{
			modules.SetIPCache(ip, lbsinfo)
		}
	}
	return
}