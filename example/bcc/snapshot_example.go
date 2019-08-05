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
	log.SetLogHandler(log.STDOUT)

	client, err := bcc.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createArgs := &api.CreateSnapshotArgs{
		VolumeId:     "",
		SnapshotName: "sdk",
		Description:  "create by sdk",
	}
	createResult, err := client.CreateSnapshot(createArgs)
	if err != nil {
		log.Errorf("create snapshot error: %+v\n", err)
		return
	}
	log.Info("create snapshot success: %+v\n", createResult)

	listArgs := &api.ListSnapshotArgs{}
	listResult, err := client.ListSnapshot(listArgs)
	if err != nil {
		log.Errorf("list snapshot error: %+v\n", err)
		return
	}
	log.Info("list snapshot success: %+v\n", listResult)

	getSnapshotResult, err := client.GetSnapshotDetail(createResult.SnapshotId)
	if err != nil {
		log.Errorf("get snapshot detail error: %+v\n", err)
		return
	}
	log.Info("get snapshot detail success: %+v\n", getSnapshotResult)

	err = client.DeleteSnapshot(createResult.SnapshotId)
	if err != nil {
		log.Errorf("delete snapshot error: %+v\n", err)
		return
	}
	log.Info("delete snapshot success\n")
}
