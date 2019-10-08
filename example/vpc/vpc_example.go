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

	vpcId, err := createVPC(client)
	if err != nil {
		log.Errorf("create vpc error: %v", err)
		return
	}

	if err := listVPC(client); err != nil {
		log.Errorf("list vpc error: %v", err)
		return
	}

	if err := getVPCDetail(client, vpcId); err != nil {
		log.Errorf("get vpc detail error: %v", err)
		return
	}

	if err := updateVPC(client, vpcId); err != nil {
		log.Errorf("update vpc error: %v", err)
		return
	}

	if err := deleteVPC(client, vpcId); err != nil {
		log.Errorf("delete vpc error: %v", err)
		return
	}
}

func createVPC(client *vpc.Client) (string, error) {
	log.Info("create vpc...")

	createVPCArgs := &vpc.CreateVPCArgs{
		Name:        "test-vpc",
		Description: "test-vpc-description",
		Cidr:        "102.168.0.0/24",
	}
	createResult, err := client.CreateVPC(createVPCArgs)
	if err != nil {
		return "", err
	}

	log.Infof("create vpc success: %v", createResult)
	return createResult.VPCID, nil
}

func listVPC(client *vpc.Client) error {
	log.Info("list vpc...")

	listArgs := &vpc.ListVPCArgs{
		MaxKeys: 2,
		Marker:  "vpc-f1mmb4ad1v5f",
	}
	listResult, err := client.ListVPC(listArgs)
	if err != nil {
		log.Errorf("list vpc error: %v", err)
		return err
	}

	log.Infoln("get vpc list: ", listResult)
	return nil
}

func getVPCDetail(client *vpc.Client, vpcId string) error {
	log.Info("get vpc detail...")

	getDetailResult, err := client.GetVPCDetail(vpcId)
	if err != nil {
		log.Errorf("get vpc %s detail error: %v", vpcId, err)
		return err
	}

	log.Infof("get vpc %s detail result: %v", vpcId, getDetailResult)
	return nil
}

func updateVPC(client *vpc.Client, vpcId string) error {
	log.Info("update vpc...")

	args := &vpc.UpdateVPCArgs{
		Name:        "test-vpc-02",
		Description: "test-vpc-description-update",
	}
	if err := client.UpdateVPC(vpcId, args); err != nil {
		log.Errorf("update vpc %s error: %v", vpcId, err)
		return err
	}

	log.Infof("update vpc %s success", vpcId)
	return nil
}

func deleteVPC(client *vpc.Client, vpcId string) error {
	log.Info("delete vpc...")

	if err := client.DeleteVPC(vpcId, ""); err != nil {
		log.Errorf("delete vpc %s error: %v", vpcId, err)
		return err
	}

	log.Infof("delete vpc %s success", vpcId)
	return nil
}
