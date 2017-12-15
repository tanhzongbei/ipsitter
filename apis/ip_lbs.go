package apis

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ipsitter/biz"
)

func QueryLBSByIP(c *gin.Context) {
	ip, have_ip := c.GetQuery("ip")
	if !have_ip {
		c.JSON(http.StatusOK, gin.H{
			"result": "bad params",
			"reason": "ip is required",
		})
	} else {
		lbsInfo, _ := biz.GetIPInfo(ip)
		c.JSON(http.StatusOK, gin.H{
			"result" : "ok",
			"ip_info" : lbsInfo})
	}
	return
}