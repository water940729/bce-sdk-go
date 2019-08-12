package main

import (
	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/util/log"
)

// bcc example config
const (
	AK       = "ak"
	SK       = "sk"
	EndPoint = "bcc.bj.baidubce.com"
	LogDir   = "./logs/"
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)

	client, err := bcc.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	listInstanceTypeResult, err := client.ListInstanceType()
	if err != nil {
		log.Errorf("list instance type error: %+v\n", err)
		return
	}
	log.Info("list instance type success: %+v\n", listInstanceTypeResult)

	listZoneResult, err := client.ListZone()
	if err != nil {
		log.Errorf("list zone error: %+v\n", err)
		return
	}
	log.Info("list zone success: %+v\n", listZoneResult)
}
