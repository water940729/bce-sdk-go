package main

import (
	"os"

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

var client *vpc.Client

func init() {
	log.SetLogLevel(log.INFO)
	log.SetLogHandler(log.STDOUT)
	log.SetLogDir(LogDir)

	var err error
	if client, err = vpc.NewClient(AK, SK, EndPoint); err != nil {
		log.Errorln("create client failed: %+v", err)
		os.Exit(1)
	}
}

func main() {
	if err := createNatGateway(); err != nil {
		return
	}

	natId, err := listNatGateway()
	if err != nil {
		return
	}

	if err := getNatGatewayDetail(natId); err != nil {
		return
	}

	if err := updateNatGateway(natId); err != nil {
		return
	}

	if err := bindEips(natId); err != nil {
		return
	}

	if err := unBindEips(natId); err != nil {
		return
	}

	if err := RenewNatGateway(natId); err != nil {
		return
	}

	if err := deleteNatGateway(natId); err != nil {
		return
	}
}

func createNatGateway() error {
	vpcId := "vpc-4njbqurm0uag"
	args := &vpc.CreateNatGatewayArgs{
		Name:  "nat_gateway",
		VpcId: vpcId,
		Spec:  vpc.NAT_GATEWAY_SPEC_SMALL,
		Billing: &vpc.Billing{
			PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
		},
	}
	result, err := client.CreateNatGateway(args)
	if err != nil {
		log.Errorf("create nat gateway error: %v", err)
		return err
	}

	log.Infof("create nat gateway success, result: %v", result)
	return nil
}

func listNatGateway() (string, error) {
	vpcId := "vpc-4njbqurm0uag"
	args := &vpc.ListNatGatewayArgs{
		VpcId: vpcId,
	}
	result, err := client.ListNatGateway(args)
	if err != nil {
		log.Errorf("list nat gateway error: %v", err)
		return "", err
	}

	log.Infof("list nat gateway success, result: %v", result)
	return result.Nats[0].Id, nil
}

func getNatGatewayDetail(natId string) error {
	result, err := client.GetNatGatewayDetail(natId)
	if err != nil {
		log.Errorf("get nat gateway error: %v", err)
		return err
	}

	log.Infof("get nat gateway success, result: %v", result)
	return nil
}

func updateNatGateway(natId string) error {
	args := &vpc.UpdateNatGatewayArgs{
		Name: "bbbb",
	}

	if err := client.UpdateNatGateway(natId, args); err != nil {
		log.Errorf("update nat gateway error: %v", err)
		return err
	}

	log.Infof("update nat gateway %s success.", natId)
	return nil
}

func bindEips(natId string) error {
	args := &vpc.BindEipsArgs{
		Eips: []string{"180.76.164.111"},
	}
	if err := client.BindEips(natId, args); err != nil {
		log.Errorf("bind eips error: %v", err)
		return err
	}

	log.Infof("bind eips success.")
	return nil
}

func unBindEips(natId string) error {
	args := &vpc.UnBindEipsArgs{
		Eips: []string{"180.76.164.111"},
	}
	if err := client.UnBindEips(natId, args); err != nil {
		log.Errorf("unbind eips error: %v", err)
		return err
	}

	log.Infof("unbind eips success.")
	return nil
}

func deleteNatGateway(natId string) error {
	if err := client.DeleteNatGateway(natId, ""); err != nil {
		log.Errorf("delete nat gateway error: %v", err)
		return err
	}

	log.Infof("delete nat gateway success.")
	return nil
}

func RenewNatGateway(natId string) error {
	args := &vpc.RenewNatGatewayArgs{
		Billing: &vpc.Billing{
			Reservation: &vpc.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}
	if err := client.RenewNatGateway(natId, args); err != nil {
		log.Errorf("renew nat gateway error: %v", err)
		return err
	}

	log.Infof("renew nat gateway %s success.", natId)
	return nil
}
