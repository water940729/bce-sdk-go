package main

import (
	"fmt"
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
	if err := listAclEntrys(); err != nil {
		return
	}

	if err := createAclRules(); err != nil {
		return
	}

	aclRuleId, err := listAclRules()
	if err != nil {
		return
	}

	if err = updateAclRule(aclRuleId); err != nil {
		return
	}

	if err = deleteAclRule(aclRuleId); err != nil {
		return
	}
}

func listAclEntrys() error {
	vpcId := "vpc-4njbqurm0uag"

	result, err := client.ListAclEntrys(vpcId)
	if err != nil {
		log.Errorf("list acl entrys error: %v", err)
		return err
	}
	log.Infof("list acl entrys success, result: %v", result)

	return nil
}

func createAclRules() error {
	subnetId := "sbn-e4cg8e8zkizs"
	requests := []vpc.AclRuleRequest{
		{
			SubnetId:             subnetId,
			Protocol:             vpc.ACL_RULE_PROTOCOL_TCP,
			SourceIpAddress:      "192.168.2.0",
			DestinationIpAddress: "192.168.0.0/24",
			SourcePort:           "8888",
			DestinationPort:      "9999",
			Position:             12,
			Direction:            vpc.ACL_RULE_DIRECTION_INGRESS,
			Action:               vpc.ACL_RULE_ACTION_ALLOW,
		},
	}
	args := &vpc.CreateAclRuleArgs{
		AclRules: requests,
	}

	if err := client.CreateAclRule(args, ""); err != nil {
		fmt.Errorf("create acl rule error: %v", err)
		return err
	}

	log.Infof("create acl rule success.")
	return nil
}

func listAclRules() (string, error) {
	subnetId := "sbn-e4cg8e8zkizs"
	args := &vpc.ListAclRulesArgs{
		SubnetId: subnetId,
	}

	result, err := client.ListAclRules(args)
	if err != nil {
		log.Errorf("list acl rules error: %v", err)
		return "", err
	}

	log.Infof("list acl rules success, result: %v", result)
	return result.AclRules[0].Id, nil
}

func updateAclRule(aclRuleId string) error {
	args := &vpc.UpdateAclRuleArgs{
		Protocol:             vpc.ACL_RULE_PROTOCOL_TCP,
		SourceIpAddress:      "192.168.2.0",
		DestinationIpAddress: "192.168.0.0/24",
		SourcePort:           "3333",
		DestinationPort:      "4444",
		Position:             12,
		Action:               vpc.ACL_RULE_ACTION_ALLOW,
	}

	if err := client.UpdateAclRule(aclRuleId, args); err != nil {
		log.Errorf("update acl rule error: %v", err)
		return err
	}

	log.Infof("update acl rule %s success.", aclRuleId)
	return nil
}

func deleteAclRule(aclRuleId string) error {
	if err := client.DeleteAclRule(aclRuleId, ""); err != nil {
		log.Errorf("delete acl rule error: %v", err)
		return err
	}

	log.Infof("delete acl rule %s success.", aclRuleId)
	return nil
}
