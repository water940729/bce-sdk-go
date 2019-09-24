package main

import (
	"github.com/baidubce/bce-sdk-go/services/appblb"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

const (
	AK       = ""
	SK       = ""
	EndPoint = "blb.bj.baidubce.com"
	LogDir   = "./logs/"
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)

	client, err := appblb.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createArgs := &appblb.CreateLoadBalancerArgs{
		ClientToken: getClientToken(),
		Name:        "sdk-blb",
		VpcId:       "",
		SubnetId:    "",
	}

	createResult, err := client.CreateLoadBalancer(createArgs)
	if err != nil {
		log.Errorf("create blb error: %+v\n", err)
		return
	}
	log.Info("create blb success: %+v\n", createResult)

	describeArgs := &appblb.DescribeLoadBalancersArgs{}
	describeResult, err := client.DescribeLoadBalancers(describeArgs)
	if err != nil {
		log.Errorf("describe blbs error: %+v\n", err)
		return
	}
	log.Info("describe blbs success: %+v\n", describeResult)

	detailResult, err := client.DescribeLoadBalancerDetail(createResult.BlbId)
	if err != nil {
		log.Errorf("detail blb error: %+v\n", err)
		return
	}
	log.Info("detail blb success: %+v\n", detailResult)

	updateArgs := &appblb.UpdateLoadBalancerArgs{
		Name:        "test-sdk",
		Description: "test desc",
	}
	err = client.UpdateLoadBalancer(createResult.BlbId, updateArgs)
	if err != nil {
		log.Errorf("update blb error: %+v\n", err)
		return
	}
	log.Info("update blb success\n")

	err = client.DeleteLoadBalancer(createResult.BlbId)
	if err != nil {
		log.Errorf("delete blb error: %+v\n", err)
		return
	}
	log.Info("delete blb success\n")
}

func getClientToken() string {
	return util.NewUUID()
}
