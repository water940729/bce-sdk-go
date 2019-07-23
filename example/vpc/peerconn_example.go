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
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)
	log.SetLogDir(LogDir)

	var err error
	if client, err = vpc.NewClient(AK, SK, EndPoint); err != nil {
		log.Errorln("create client failed: %+v", err)
		os.Exit(1)
	}
}

func main() {
	if err := createPeerConn(); err != nil {
		return
	}

	peerConnId, err := listPeerConn()
	if err != nil {
		return
	}

	localIfId, err := getPeerConnDetail(peerConnId)
	if err != nil {
		return
	}

	if err := updatePeerConn(peerConnId, localIfId); err != nil {
		return
	}

	if err := resizePeerConn(peerConnId); err != nil {
		return
	}

	if err := renewPeerConn(peerConnId); err != nil {
		return
	}

	if err := openPeerConnSyncDNS(peerConnId); err != nil {
		return
	}

	if err := closePeerConnSyncDNS(peerConnId); err != nil {
		return
	}
}

func createPeerConn() error {
	localVpcId := "vpc-4njbqurm0uag"
	peerVpcId := "vpc-9dirbh05pkq4"
	args := &vpc.CreatePeerConnArgs{
		BandwidthInMbps: 10,
		LocalVpcId:      localVpcId,
		PeerVpcId:       peerVpcId,
		PeerRegion:      "bj",
		Billing: &vpc.Billing{
			PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
		},
	}
	result, err := client.CreatePeerConn(args)
	if err != nil {
		log.Errorf("create peerconn error: %v", err)
		return err
	}

	log.Infof("create peerconn success, result: %v", result)
	return nil
}

func listPeerConn() (string, error) {
	args := &vpc.ListPeerConnsArgs{}
	result, err := client.ListPeerConn(args)
	if err != nil {
		log.Errorf("list peer conn error: %v", err)
		return "", err
	}

	log.Infof("list peer conn success, result: %v", result)
	return result.PeerConns[0].PeerConnId, nil
}

func getPeerConnDetail(peerConnId string) (string, error) {
	result, err := client.GetPeerConnDetail(peerConnId, vpc.PEERCONN_ROLE_INITIATOR)
	if err != nil {
		log.Errorf("get peer conn detail error: %v", err)
		return "", err
	}

	log.Infof("get peer conn %s detail success, result: %v", peerConnId, result)
	return result.LocalIfId, nil
}

func updatePeerConn(peerConnId, localIfId string) error {
	args := &vpc.UpdatePeerConnArgs{
		LocalIfId:   localIfId,
		LocalIfName: "test-update",
		Description: "test-description",
	}
	if err := client.UpdatePeerConn(peerConnId, args); err != nil {
		log.Errorf("update peer conn error: %v", err)
		return err
	}

	log.Infof("update peer conn %s success.", peerConnId)
	return nil
}

func acceptPeerConnApply(peerConnId string) error {
	if err := client.AcceptPeerConnApply(peerConnId, ""); err != nil {
		log.Errorf("accept peer conn error: %v", err)
		return err
	}

	log.Infof("accept peer conn %s success.", peerConnId)
	return nil
}

func rejectPeerConnApply(peerConnId string) error {
	if err := client.RejectPeerConnApply(peerConnId, ""); err != nil {
		log.Errorf("reject peer conn error: %v", err)
		return err
	}

	log.Infof("reject peer conn %s success.", peerConnId)
	return nil
}

func deletePeerConn(peerConnId string) error {
	if err := client.DeletePeerConn(peerConnId, ""); err != nil {
		log.Errorf("delete peer conn error: %v", err)
		return err
	}

	log.Infof("delete peer conn %s success.", peerConnId)
	return nil
}

func resizePeerConn(peerConnId string) error {
	args := &vpc.ResizePeerConnArgs{
		NewBandwidthInMbps: 20,
	}

	if err := client.ResizePeerConn(peerConnId, args); err != nil {
		log.Errorf("resize peer conn error: %v", err)
		return err
	}

	log.Infof("resize peer conn %s success.", peerConnId)
	return nil
}

func renewPeerConn(peerConnId string) error {
	args := &vpc.RenewPeerConnArgs{
		Billing: &vpc.Billing{
			Reservation: &vpc.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}

	if err := client.RenewPeerConn(peerConnId, args); err != nil {
		log.Errorf("renew peer conn error: %v", err)
		return err
	}

	log.Infof("renew peer conn %s success.", peerConnId)
	return nil
}

func openPeerConnSyncDNS(peerConnId string) error {
	args := &vpc.PeerConnSyncDNSArgs{
		Role: vpc.PEERCONN_ROLE_INITIATOR,
	}

	if err := client.OpenPeerConnSyncDNS(peerConnId, args); err != nil {
		log.Errorf("open peer conn sync dns error: %v", err)
		return err
	}

	log.Infof("open peer conn %s sync dns success.", peerConnId)
	return nil
}

func closePeerConnSyncDNS(peerConnId string) error {
	args := &vpc.PeerConnSyncDNSArgs{
		Role: vpc.PEERCONN_ROLE_INITIATOR,
	}

	if err := client.ClosePeerConnSyncDNS(peerConnId, args); err != nil {
		log.Errorf("close peer conn sync dns error: %v", err)
		return err
	}

	log.Infof("close peer conn %s sync dns success.", peerConnId)
	return nil
}
