package main

import (
	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

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

	createCdsArgs := &api.CreateCDSVolumeArgs{
		PurchaseCount: 1,
		CdsSizeInGB:   40,
		StorageType:   api.StorageTypeSSD,
		Billing: &api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
	}
	createResult, err := client.CreateCDSVolume(createCdsArgs)
	if err != nil {
		log.Errorf("create cds volume error: %+v\n", err)
		return
	}
	log.Info("create cds success: %+v\n", createResult)

	listCdsArgs := &api.ListCDSVolumeArgs{}
	listCdsResult, err := client.ListCDSVolume(listCdsArgs)
	if err != nil {
		log.Errorf("list cds volume error: %+v\n", err)
		return
	}
	log.Info("list cds volume success: %+v\n", listCdsResult)

	volumeDetail, err := client.GetCDSVolumeDetail((*createResult.VolumeIds)[0])
	if err != nil {
		log.Errorf("get cds volume detail error: %+v\n", err)
		return
	}
	log.Infof("get cds volume detail success: %+v\n", volumeDetail)

	attachCdsArgs := &api.AttachVolumeArgs{
		InstanceId: "i-AhFKN8Gb",
	}
	attachCdsResult, err := client.AttachCDSVolume((*createResult.VolumeIds)[0], attachCdsArgs)
	if err != nil {
		log.Errorf("attach cds to bcc error: %+v\n", err)
		return
	}
	log.Infof("attach cds to bcc success: %+v\n", attachCdsResult)

	detachCdsArgs := &api.DetachVolumeArgs{
		InstanceId: "i-AhFKN8Gb",
	}
	err = client.DetachCDSVolume((*createResult.VolumeIds)[0], detachCdsArgs)
	if err != nil {
		log.Errorf("detach cds volume error: %+v\n", err)
		return
	}
	log.Infof("detach cds volume success\n")

	resizeCdsArgs := &api.ResizeCSDVolumeArgs{
		NewCdsSizeInGB: 100,
	}
	err = client.ResizeCDSVolume((*createResult.VolumeIds)[0], resizeCdsArgs)
	if err != nil {
		log.Errorf("resize cds volume error: %+v\n", err)
		return
	}
	log.Infof("resize cds volume success\n")

	renameArgs := &api.RenameCSDVolumeArgs{
		Name: "tf-testVolume",
	}
	err = client.RenameCDSVolume((*createResult.VolumeIds)[0], renameArgs)
	if err != nil {
		log.Errorf("rename cds volume error: %+v\n", err)
		return
	}
	log.Infof("rename cds volume success\n")

	modifyArgs := &api.ModifyCSDVolumeArgs{
		CdsName: "tf-aaa",
		Desc:    "tf-desc",
	}
	err = client.ModifyCDSVolume((*createResult.VolumeIds)[0], modifyArgs)
	if err != nil {
		log.Errorf("modify cds volume error: %+v\n", err)
		return
	}
	log.Infof("modify cds volume success\n")

	err = client.DeleteCDSVolume((*createResult.VolumeIds)[0])
	if err != nil {
		log.Errorf("delete cds volume error: %+v\n", err)
		return
	}
	log.Infof("delete cds success\n")
}