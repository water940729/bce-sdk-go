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

	createArgs := &api.CreateSecurityGroupArgs{
		Name: "sdk-create",
		Rules: []api.SecurityGroupRuleModel{
			{
				Remark:        "备注",
				Protocol:      "tcp",
				PortRange:     "1-65535",
				Direction:     "ingress",
				SourceIp:      "",
				SourceGroupId: "",
			},
		},
	}
	result, err := client.CreateSecurityGroup(createArgs)
	if err != nil {
		log.Errorf("create security group error: %+v\n", err)
		return
	}
	log.Info("create sg success: %+v\n", result)

	listCdsArgs := &api.ListSecurityGroupArgs{}
	listCdsResult, err := client.ListSecurityGroup(listCdsArgs)
	if err != nil {
		log.Errorf("list security group error: %+v\n", err)
		return
	}
	log.Info("list sg success: %+v\n", listCdsResult)

	authorArgs := &api.AuthorizeSecurityGroupArgs{
		Rule: &api.SecurityGroupRuleModel{
			Remark:        "备注",
			Protocol:      "udp",
			PortRange:     "1-65535",
			Direction:     "ingress",
			SourceIp:      "",
			SourceGroupId: "",
		},
	}
	err = client.AuthorizeSecurityGroupRule(result.SecurityGroupId, authorArgs)
	if err != nil {
		log.Errorf("authorize security group error: %+v\n", err)
		return
	}
	log.Info("authorize sg success\n")

	revokeArgs := &api.RevokeSecurityGroupArgs{
		Rule: &api.SecurityGroupRuleModel{
			Remark:        "备注",
			Protocol:      "udp",
			PortRange:     "1-65535",
			Direction:     "ingress",
			SourceIp:      "",
			SourceGroupId: "",
		},
	}
	err = client.RevokeSecurityGroupRule(result.SecurityGroupId, revokeArgs)
	if err != nil {
		log.Errorf("revoke security group error: %+v\n", err)
		return
	}
	log.Info("revoke sg success\n")

	err = client.DeleteSecurityGroupRule(result.SecurityGroupId)
	if err != nil {
		log.Errorf("delete security group error: %+v\n", err)
		return
	}
	log.Info("delete sg success\n")

}
