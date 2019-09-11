package main

import (
	"github.com/baidubce/bce-sdk-go/services/appblb"
	"github.com/baidubce/bce-sdk-go/util/log"
	"github.com/gofrs/uuid"
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

	blbId := "lb-b960c43e"

	AppServerGroupExample(client, blbId)
	AppServerGroupPortExample(client, blbId)
	BlbRsExample(client, blbId)
}

func AppServerGroupExample(client *appblb.Client, blbId string) {
	createArgs := &appblb.CreateAppServerGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdk-test",
	}
	createResult, err := client.CreateAppServerGroup(blbId, createArgs)
	if err != nil {
		log.Errorf("create blb app server group error: %+v\n", err)
		return
	}
	log.Info("create blb app server group success: %+v\n", createResult)

	updateArgs := &appblb.UpdateAppServerGroupArgs{
		SgId:        createResult.Id,
		Name:        "test-sssdk",
		Description: "test desc",
		ClientToken: getClientToken(),
	}
	err = client.UpdateAppServerGroup(blbId, updateArgs)
	if err != nil {
		log.Errorf("update blb app server group error: %+v\n", err)
		return
	}
	log.Info("update blb app server group success\n")

	describeArgs := &appblb.DescribeAppServerGroupArgs{}
	describeResult, err := client.DescribeAppServerGroup(blbId, describeArgs)
	if err != nil {
		log.Errorf("describe blb app server group error: %+v\n", err)
		return
	}
	log.Info("describe blb app server group success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppServerGroupArgs{
		SgId:        createResult.Id,
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppServerGroup(blbId, deleteArgs)
	if err != nil {
		log.Errorf("delete blb app server group error: %+v\n", err)
		return
	}
	log.Info("delete blb app server group success\n")
}

func AppServerGroupPortExample(client *appblb.Client, blbId string) {
	createASGArgs := &appblb.CreateAppServerGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdk-test",
	}
	createASGResult, err := client.CreateAppServerGroup(blbId, createASGArgs)
	if err != nil {
		log.Errorf("create blb app server group error: %+v\n", err)
		return
	}
	log.Info("create blb app server group success: %+v\n", createASGResult)

	createArgs := &appblb.CreateAppServerGroupPortArgs{
		ClientToken: getClientToken(),
		SgId:        createASGResult.Id,
		Port:        80,
		Type:        "TCP",
	}
	createResult, err := client.CreateAppServerGroupPort(blbId, createArgs)
	if err != nil {
		log.Errorf("create blb app server group port error: %+v\n", err)
		return
	}
	log.Info("create blb app server group port success: %+v\n", createResult)

	updateArgs := &appblb.UpdateAppServerGroupPortArgs{
		ClientToken:                 getClientToken(),
		SgId:                        createASGResult.Id,
		PortId:                      createResult.Id,
		HealthCheck:                 "TCP",
		HealthCheckPort:             30,
		HealthCheckIntervalInSecond: 10,
		HealthCheckTimeoutInSecond:  10,
	}
	err = client.UpdateAppServerGroupPort(blbId, updateArgs)
	if err != nil {
		log.Errorf("update blb app server group port error: %+v\n", err)
		return
	}
	log.Info("update blb app server group port success\n")

	deleteArgs := &appblb.DeleteAppServerGroupPortArgs{
		SgId:        createASGResult.Id,
		PortIdList:  []string{createResult.Id},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppServerGroupPort(blbId, deleteArgs)
	if err != nil {
		log.Errorf("delete blb app server group port error: %+v\n", err)
		return
	}
	log.Info("delete blb app server group port success\n")

	deleteASGArgs := &appblb.DeleteAppServerGroupArgs{
		SgId:        createASGResult.Id,
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppServerGroup(blbId, deleteASGArgs)
	if err != nil {
		log.Errorf("delete blb app server group error: %+v\n", err)
		return
	}
	log.Info("delete blb app server group success\n")
}

func BlbRsExample(client *appblb.Client, blbId string) {
	createASGArgs := &appblb.CreateAppServerGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdk-test",
	}
	createASGResult, err := client.CreateAppServerGroup(blbId, createASGArgs)
	if err != nil {
		log.Errorf("create blb app server group error: %+v\n", err)
		return
	}
	log.Info("create blb app server group success: %+v\n", createASGResult)

	createArgs := &appblb.CreateBlbRsArgs{
		BlbRsWriteOpArgs: appblb.BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        createASGResult.Id,
			BackendServerList: []appblb.AppBackendServer{
				{InstanceId: "i-0W3dBG7G", Weight: 30},
			},
		},
	}
	err = client.CreateBlbRs(blbId, createArgs)
	if err != nil {
		log.Errorf("create blb rs error: %+v\n", err)
		return
	}
	log.Info("create blb rs success\n")

	updateArgs := &appblb.UpdateBlbRsArgs{
		BlbRsWriteOpArgs: appblb.BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        createASGResult.Id,
			BackendServerList: []appblb.AppBackendServer{
				{InstanceId: "i-0W3dBG7G", Weight: 50},
			},
		},
	}
	err = client.UpdateBlbRs(blbId, updateArgs)
	if err != nil {
		log.Errorf("update blb rs error: %+v\n", err)
		return
	}
	log.Info("update blb rs success\n")

	describeArgs := &appblb.DescribeBlbRsArgs{
		SgId: createASGResult.Id,
	}
	describeResult, err := client.DescribeBlbRs(blbId, describeArgs)
	if err != nil {
		log.Errorf("describe blb rs error: %+v\n", err)
		return
	}
	log.Info("describe blb rs success: %+v\n", describeResult)

	describeMountResult, err := client.DescribeRsMount(blbId, createASGResult.Id)
	if err != nil {
		log.Errorf("describe blb rs mount error: %+v\n", err)
		return
	}
	log.Info("describe blb rs mount success: %+v\n", describeMountResult)

	describeUnMountResult, err := client.DescribeRsUnMount(blbId, createASGResult.Id)
	if err != nil {
		log.Errorf("describe blb rs unmount error: %+v\n", err)
		return
	}
	log.Info("describe blb rs unmount success: %+v\n", describeUnMountResult)

	deleteArgs := &appblb.DeleteBlbRsArgs{
		SgId:                createASGResult.Id,
		BackendServerIdList: []string{"i-0W3dBG7G"},
		ClientToken:         getClientToken(),
	}
	err = client.DeleteBlbRs(blbId, deleteArgs)
	if err != nil {
		log.Errorf("delete blb rs error: %+v\n", err)
		return
	}
	log.Info("delete blb rs success\n")

	deleteASGArgs := &appblb.DeleteAppServerGroupArgs{
		SgId:        createASGResult.Id,
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppServerGroup(blbId, deleteASGArgs)
	if err != nil {
		log.Errorf("delete blb app server group error: %+v\n", err)
		return
	}
	log.Info("delete blb app server group success\n")
}

func getClientToken() string {
	u, _ := uuid.NewV4()
	return u.String()
}
