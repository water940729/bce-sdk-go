package main

import (
	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util/log"
)

// vpc example config
const (
	AK       = "XXXX"
	SK       = "XXXX"
	EndPoint = "bcc.bj.baidubce.com"
	LogDir   = "./logs/"
)

func main() {
	log.SetLogLevel(log.INFO)
	log.SetLogHandler(log.STDOUT)

	client, err := vpc.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorln("create client failed: %+v", err)
		return
	}

	getRouteTableResult, err := getRouteTableDetail(client)
	if err != nil {
		return
	}

	createRouteRuleResult, err := createRouteRule(client, getRouteTableResult)
	if err != nil {
		return
	}

	if err = deleteRouteRule(client, createRouteRuleResult); err != nil {
		return
	}
}

func getRouteTableDetail(client *vpc.Client) (*vpc.GetRouteTableResult, error) {
	vpcId := "vpc-4njbqurm0uag"
	result, err := client.GetRouteTableDetail("", vpcId)
	if err != nil {
		log.Errorf("get route table error: %v", err)
		return nil, err
	}
	log.Infof("get route table success: %v", result)

	return result, err
}

func createRouteRule(client *vpc.Client, getRouteTableResult *vpc.GetRouteTableResult) (*vpc.CreateRouteRuleResult, error) {
	createRouteRuleArgs := &vpc.CreateRouteRuleArgs{
		RouteTableId:       getRouteTableResult.RouteTableId,
		SourceAddress:      "192.168.0.0/24",
		DestinationAddress: "192.168.1.0/24",
		NexthopId:          "vpn-bx831m72tqqd",
		NexthopType:        "vpn",
		Description:        "test description",
	}
	createRouteRuleResult, err := client.CreateRouteRule(createRouteRuleArgs)
	if err != nil {
		log.Errorf("create route rule error: %v", err)
		return nil, err
	}
	log.Infof("create route rule success, result: %v", createRouteRuleResult)
	return createRouteRuleResult, err
}

func deleteRouteRule(client *vpc.Client, createRouteRuleResult *vpc.CreateRouteRuleResult) error {
	if err := client.DeleteRouteRule(createRouteRuleResult.RouteRuleId, ""); err != nil {
		log.Errorf("delete route rule error: %v", err)
		return err
	}
	log.Infof("delete route rule %s success.", createRouteRuleResult.RouteRuleId)
	return nil
}
