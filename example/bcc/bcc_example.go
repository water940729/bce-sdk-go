package main

import (
	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
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
	log.SetLogHandler(log.FILE)
	log.SetLogDir(LogDir)

	client, err := bcc.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createInstanceArgs := &api.CreateInstanceArgs{
		ImageId: "m-DpgNg8lO",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		InstanceType:        api.InstanceTypeN1,
		CpuCount:            1,
		MemoryCapacityInGB:  1,
		RootDiskSizeInGb:    40,
		RootDiskStorageType: api.StorageTypeCloudHP1,
		CreateCdsList: []api.CreateCdsModel{
			{
				StorageType: api.StorageTypeSSD,
				CdsSizeInGB: 0,
			},
		},
		AdminPass: "123qaz!@#",
		Name:      "terraform_sdkTest",
	}
	createResult, err := client.CreateInstance(createInstanceArgs)
	if err != nil {
		log.Errorf("create instance failed: %+v\n", err)
		return
	}
	log.Info("create instance success: %+v\n", createResult)

	listArgs := &api.ListInstanceArgs{}
	listResult, err := client.ListInstances(listArgs)
	if err != nil {
		log.Errorf("list instance failed: %+v\n", err)
		return
	}
	log.Info("list instance success: %+v\n", listResult)

	instanceDetail, err := client.GetInstanceDetail(createResult.InstanceIds[0])
	if err != nil {
		log.Errorf("get instance detail failed: %+v\n", err)
		return
	}
	log.Info("get instance detail success: %+v\n", instanceDetail)

	resizeArgs := &api.ResizeInstanceArgs{
		CpuCount:           2,
		MemoryCapacityInGB: 4,
	}
	err = client.ResizeInstance(createResult.InstanceIds[0], resizeArgs)
	if err != nil {
		log.Errorf("resize instance error: %+v\n", err)
		return
	}
	log.Info("resize instance success.\n")

	err = client.StopInstance(createResult.InstanceIds[0], true)
	if err != nil {
		log.Errorf("stop instance error: %+v\n", err)
		return
	}
	log.Info("stop instance success.\n")

	err = client.StartInstance(createResult.InstanceIds[0])
	if err != nil {
		log.Errorf("start instance error: %+v\n", err)
		return
	}
	log.Info("start instance success.\n")

	err = client.RebootInstance(createResult.InstanceIds[0], true)
	if err != nil {
		log.Errorf("reboot instance error: %+v\n", err)
		return
	}
	log.Info("reboot instance success.\n")

	rebuildArgs := &api.RebuildInstanceArgs{
		ImageId:   "m-DpgNg8lO",
		AdminPass: "123qaz!@#",
	}
	err = client.RebuildInstance(createResult.InstanceIds[0], rebuildArgs)
	if err != nil {
		log.Errorf("rebuild instance error: %+v\n", err)
		return
	}
	log.Info("rebuild instance success.\n")

	changeArgs := &api.ChangeInstancePassArgs{
		AdminPass: "321zaq#@!",
	}
	err = client.ChangeInstancePass(createResult.InstanceIds[0], changeArgs)
	if err != nil {
		log.Errorf("change instance pass error: %+v\n", err)
		return
	}
	log.Info("change instance pass success.\n")

	modifyArgs := &api.ModifyInstanceAttributeArgs{
		Name: "terraform-modify",
	}
	err = client.ModifyInstanceAttribute(createResult.InstanceIds[0], modifyArgs)
	if err != nil {
		log.Errorf("modify instance error: %+v\n", err)
		return
	}
	log.Info("modify instance success.\n")

	vncResult, err := client.GetInstanceVNC(createResult.InstanceIds[0])
	if err != nil {
		log.Errorf("get instance vnc error: %+v\n", err)
		return
	}
	log.Info("get instance vnc success: %+v\n", vncResult)

	err = client.BindSecurityGroup(createResult.InstanceIds[0], "g-x7jis4ytps4e")
	if err != nil {
		log.Errorf("bind instance sg vnc error: %+v\n", err)
		return
	}
	log.Info("bind instance sg success\n")

	err = client.UnBindSecurityGroup(createResult.InstanceIds[0], "g-x7jis4ytps4e")
	if err != nil {
		log.Errorf("bind instance sg vnc error: %+v\n", err)
		return
	}
	log.Info("bind instance sg success\n")

	err = client.DeleteInstance(createResult.InstanceIds[0])
	if err != nil {
		log.Errorf("delete instance failed: %+v\n", err)
		return
	}
	log.Info("delete instance success.\n")
}
