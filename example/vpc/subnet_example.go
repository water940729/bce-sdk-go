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

	subnetId, err := createSubnet(client)
	if err != nil {
		return
	}

	if err = listSubnet(client); err != nil {
		return
	}

	if err = getSubnetDetail(client, subnetId); err != nil {
		return
	}

	if err = updateSubnet(client, subnetId); err != nil {
		return
	}

	if err = deleteSubnet(client, subnetId); err != nil {
		return
	}
}

func createSubnet(client *vpc.Client) (string, error) {
	log.Infoln("create subnet...")

	args := &vpc.CreateSubnetArgs{
		Name:     "my-subnet",
		ZoneName: "cn-bj-a",
		Cidr:     "192.168.1.0/24",
		VpcId:    "vpc-4njbqurm0uag",
	}
	result, err := client.CreateSubnet(args)
	if err != nil {
		log.Errorf("create subnet error: %v", err)
		return "", err
	}

	log.Infof("create subnet success: %v", result)
	return result.SubnetId, nil
}

func listSubnet(client *vpc.Client) error {
	log.Info("list subnet...")

	args := &vpc.ListSubnetArgs{}
	result, err := client.ListSubnets(args)
	if err != nil {
		log.Errorf("list subnets error: %v", err)
		return err
	}
	log.Infof("list subnets: %v", result)
	return nil
}

func getSubnetDetail(client *vpc.Client, subnetId string) error {
	log.Info("get subnet detail...")

	result, err := client.GetSubnetDetail(subnetId)
	if err != nil {
		log.Errorf("get subnet detail error: %v", err)
		return err
	}
	log.Infof("get subnet detail: %v", result)

	return nil
}

func updateSubnet(client *vpc.Client, subnetId string) error {
	log.Info("update subnet...")

	args := &vpc.UpdateSubnetArgs{
		Name: "update-name",
	}
	if err := client.UpdateSubnet(subnetId, args); err != nil {
		log.Errorf("update subnet error: %v", err)
		return err
	}
	log.Infof("update subnet success.")

	return nil
}

func deleteSubnet(client *vpc.Client, subnetId string) error {
	log.Info("delete subnet...")

	if err := client.DeleteSubnet(subnetId, ""); err != nil {
		log.Errorf("delete subnet error: %v", err)
		return err
	}
	log.Infof("delete subnet success")

	return nil
}
