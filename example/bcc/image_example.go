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
	log.SetLogHandler(log.STDOUT)

	client, err := bcc.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createImageArgs := &api.CreateImageArgs{
		ImageName:  "sdk-test",
		InstanceId: "i-ZXhqfS39",
	}
	createResult, err := client.CreateImage(createImageArgs)
	if err != nil {
		log.Errorf("create image error: %+v\n", err)
		return
	}
	log.Info("create image success: %+v\n", createResult)

	listArgs := &api.ListImageArgs{}
	listResult, err := client.ListImage(listArgs)
	if err != nil {
		log.Errorf("list image error: %+v\n", err)
		return
	}
	log.Info("list image success: %+v\n", listResult)

	getDetailResult, err := client.GetImageDetail(createResult.ImageId)
	if err != nil {
		log.Errorf("get image detail error: %+v\n", err)
		return
	}
	log.Infof("get image detail success: %+v\n", getDetailResult)

	remoteCopyArgs := &api.RemoteCopyImageArgs{
		Name:       "sdk-test2",
		DestRegion: []string{"gz"},
	}
	err = client.RemoteCopyImage(createResult.ImageId, remoteCopyArgs)
	if err != nil {
		log.Errorf("remote copy image error: %+v\n", err)
		return
	}
	log.Infof("remote copy image success\n")

	err = client.CancelRemoteCopyImage(createResult.ImageId)
	if err != nil {
		log.Errorf("cancel remote copy image error: %+v\n", err)
		return
	}
	log.Infof("cancel remote copy image success\n")

	shareArgs := &api.SharedUser{
		AccountId: "id",
	}
	err = client.ShareImage(createResult.ImageId, shareArgs)
	if err != nil {
		log.Errorf("share image error: %+v\n", err)
		return
	}
	log.Infof("share image success\n")

	sharedResult, err := client.GetImageSharedUser(createResult.ImageId)
	if err != nil {
		log.Errorf("get image shared user error: %+v\n", err)
		return
	}
	log.Infof("get image shared user success: %+v\n", sharedResult)

	unsharedArgs := &api.SharedUser{
		AccountId: "id",
	}
	err = client.UnShareImage(createResult.ImageId, unsharedArgs)
	if err != nil {
		log.Errorf("unShared image error: %+v\n", err)
		return
	}
	log.Infof("unShared image success\n")

	osArgs := &api.GetImageOsArgs{
		InstanceIds: []string{"i-ZXhqfS39"},
	}
	osInfoResult, err := client.GetImageOS(osArgs)
	if err != nil {
		log.Errorf("get image os info error: %+v\n", err)
		return
	}
	log.Infof("get image os info success: %+v\n", osInfoResult)

	err = client.DeleteImage(createResult.ImageId)
	if err != nil {
		log.Errorf("delete image error: %+v\n", err)
		return
	}
	log.Infof("delete image success\n")
}
