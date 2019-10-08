package main

import (
	"github.com/baidubce/bce-sdk-go/services/eip"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

const (
	AK       = ""
	SK       = ""
	EndPoint = "eip.bj.baidubce.com"
	LogDir   = "./logs/"
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)

	client, err := eip.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createArgs := &eip.CreateEipArgs{
		Name:            "sdk-eip",
		BandWidthInMbps: 1000,
		Billing: &eip.Billing{
			PaymentTiming: "Postpaid",
			BillingMethod: "ByTraffic",
		},
		ClientToken: getClientToken(),
	}
	createResult, err := client.CreateEip(createArgs)
	if err != nil {
		log.Errorf("create eip error: %+v\n", err)
		return
	}
	log.Info("create eip success: %+v\n", createResult)

	resizeArgs := &eip.ResizeEipArgs{
		NewBandWidthInMbps: 200,
		ClientToken:        getClientToken(),
	}
	err = client.ResizeEip(createResult.Eip, resizeArgs)
	if err != nil {
		log.Errorf("resize eip error: %+v\n", err)
		return
	}
	log.Info("resize eip success\n")

	bindArgs := &eip.BindEipArgs{
		InstanceType: "BCC",
		InstanceId:   "",
		ClientToken:  getClientToken(),
	}
	err = client.BindEip(createResult.Eip, bindArgs)
	if err != nil {
		log.Errorf("eip bind bcc error: %+v\n", err)
		return
	}
	log.Info("eip bind bcc success\n")

	listArgs := &eip.ListEipArgs{
		Eip: createResult.Eip,
	}
	listResult, err := client.ListEip(listArgs)
	if err != nil {
		log.Errorf("list eip error: %+v\n", err)
		return
	}
	log.Info("list eip success: %+v\n", listResult)

	err = client.UnBindEip(createResult.Eip, getClientToken())
	if err != nil {
		log.Errorf("eip unbind bcc error: %+v\n", err)
		return
	}
	log.Info("eip unbind bcc success\n")

	err = client.DeleteEip(createResult.Eip, getClientToken())
	if err != nil {
		log.Errorf("delete eip error: %+v\n", err)
		return
	}
	log.Info("delete eip success\n")
}

func getClientToken() string {
	return util.NewUUID()
}
