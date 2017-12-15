package modules

import (
	"testing"
	"net/http"
	"time"
)

func TestQueryLBSByIpFromBaidu(t *testing.T) {
	httpClient = &http.Client{Timeout: time.Second * time.Duration(1)}

	res, err := QueryLBSByIpFromBaidu("14.152.49.250")
	if err != nil {
		println(res)
		t.Fail()
	}
}