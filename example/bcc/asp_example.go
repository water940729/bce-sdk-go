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

	createArgs := &api.CreateASPArgs{
		Name:"sdk-create",
		TimePoints:[]string{"20"},
		RepeatWeekdays:[]string{"1", "5"},
		RetentionDays:"7",
	}
	createResult, err := client.CreateAutoSnapshotPolicy(createArgs)
	if err != nil {
		log.Errorf("create ASP failed: %+v\n", err)
		return
	}
	log.Info("create ASP success: %+v\n", createResult)

	listArgs := &api.ListASPArgs{}
	listResult, err := client.ListAutoSnapshotPolicy(listArgs)
	if err != nil {
		log.Errorf("list ASP failed: %+v\n", err)
		return
	}
	log.Info("list ASP success: %+v\n", listResult)

	getDetailResult, err := client.GetAutoSnapshotPolicy(createResult.AspId)
	if err != nil {
		log.Errorf("get ASP detail failed: %+v\n", err)
		return
	}
	log.Info("get ASP detail success: %+v\n", getDetailResult)

	attachArgs := &api.AttachASPArgs{
		VolumeIds:[]string{"v-Trb3rQXa"},
	}
	err = client.AttachAutoSnapshotPolicy(createResult.AspId, attachArgs)
	if err != nil {
		log.Errorf("attach ASP failed: %+v\n", err)
		return
	}
	log.Info("attach ASP success\n")

	detachArgs := &api.DetachASPArgs{
		VolumeIds:[]string{"v-Trb3rQXa"},
	}
	err = client.DetachAutoSnapshotPolicy(createResult.AspId, detachArgs)
	if err != nil {
		log.Errorf("Detach ASP failed: %+v\n", err)
		return
	}
	log.Info("Detach ASP success\n")

	err = client.DeleteAutoSnapshotPolicy(createResult.AspId)
	if err != nil {
		log.Errorf("Delete ASP failed: %+v\n", err)
		return
	}
	log.Info("Delete ASP success\n")
}