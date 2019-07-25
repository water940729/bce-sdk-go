package main

import (
	"github.com/baidubce/bce-sdk-go/services/appblb"
	"github.com/baidubce/bce-sdk-go/util/log"
	"github.com/gofrs/uuid"
)

const (
	AK       = ""
	SK       = ""
	EndPoint = "blb.bj.baidubce.com"
	LogDir   = "./logs/"
	BLBID    = ""
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)

	client, err := appblb.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	TCPListenerExample(client, BLBID)
	UDPListenerExample(client, BLBID)
	HTTPListenerExample(client, BLBID)
	HTTPSListenerExample(client, BLBID)
	SSLListenerExample(client, BLBID)
	PolicyTest(client, BLBID)
}

func TCPListenerExample(client *appblb.Client, BLBID string) {
	createArgs := &appblb.CreateAppTCPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 100,
		Scheduler:    "RoundRobin",
	}
	err := client.CreateAppTCPListener(BLBID, createArgs)
	if err != nil {
		log.Errorf("create tcp listener error: %+v\n", err)
		return
	}
	log.Info("create tcp listener success\n")

	updateArgs := &appblb.UpdateAppTCPListenerArgs{
		UpdateAppListenerArgs: appblb.UpdateAppListenerArgs{
			ListenerPort: 90,
			Scheduler:    "Hash",
		},
	}
	err = client.UpdateAppTCPListener(BLBID, updateArgs)
	if err != nil {
		log.Errorf("update tcp listener error: %+v\n", err)
		return
	}
	log.Info("update tcp listener success\n")

	describeArgs := &appblb.DescribeAppListenerArgs{
		ListenerPort: 90,
	}
	describeResult, err := client.DescribeAppTCPListeners(BLBID, describeArgs)
	if err != nil {
		log.Errorf("describe tcp listener error: %+v\n", err)
		return
	}
	log.Info("describe tcp listener success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppListenersArgs{
		PortList:    []uint16{80, 100},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppListeners(BLBID, deleteArgs)
	if err != nil {
		log.Errorf("delete tcp listener error: %+v\n", err)
		return
	}
	log.Info("delete tcp listener success\n")
}

func UDPListenerExample(client *appblb.Client, BLBID string) {
	createArgs := &appblb.CreateAppUDPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "RoundRobin",
	}
	err := client.CreateAppUDPListener(BLBID, createArgs)
	if err != nil {
		log.Errorf("create udp listener error: %+v\n", err)
		return
	}
	log.Info("create udp listener success\n")

	updateArgs := &appblb.UpdateAppUDPListenerArgs{
		UpdateAppListenerArgs: appblb.UpdateAppListenerArgs{
			ListenerPort: 80,
			Scheduler:    "Hash",
		},
	}
	err = client.UpdateAppUDPListener(BLBID, updateArgs)
	if err != nil {
		log.Errorf("update udp listener error: %+v\n", err)
		return
	}
	log.Info("update udp listener success\n")

	describeArgs := &appblb.DescribeAppListenerArgs{
		ListenerPort: 80,
	}
	describeResult, err := client.DescribeAppUDPListeners(BLBID, describeArgs)
	if err != nil {
		log.Errorf("describe udp listener error: %+v\n", err)
		return
	}
	log.Info("describe udp listener success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppListenersArgs{
		PortList:    []uint16{80},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppListeners(BLBID, deleteArgs)
	if err != nil {
		log.Errorf("delete udp listener error: %+v\n", err)
		return
	}
	log.Info("delete udp listener success\n")
}

func HTTPListenerExample(client *appblb.Client, BLBID string) {
	createArgs := &appblb.CreateAppHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "RoundRobin",
	}
	err := client.CreateAppHTTPListener(BLBID, createArgs)
	if err != nil {
		log.Errorf("create http listener error: %+v\n", err)
		return
	}
	log.Info("create http listener success\n")

	updateArgs := &appblb.UpdateAppHTTPListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "LeastConnection",
		KeepSession:  true,
	}
	err = client.UpdateAppHTTPListener(BLBID, updateArgs)
	if err != nil {
		log.Errorf("update http listener error: %+v\n", err)
		return
	}
	log.Info("update http listener success\n")

	describeArgs := &appblb.DescribeAppListenerArgs{
		ListenerPort: 80,
	}
	describeResult, err := client.DescribeAppHTTPListeners(BLBID, describeArgs)
	if err != nil {
		log.Errorf("describe http listener error: %+v\n", err)
		return
	}
	log.Info("describe http listener success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppListenersArgs{
		PortList:    []uint16{80},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppListeners(BLBID, deleteArgs)
	if err != nil {
		log.Errorf("delete http listener error: %+v\n", err)
		return
	}
	log.Info("delete http listener success\n")
}

func HTTPSListenerExample(client *appblb.Client, BLBID string) {
	createArgs := &appblb.CreateAppHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "RoundRobin",
		CertIds:      []string{"cert-xvysj80uif1y"},
	}
	err := client.CreateAppHTTPSListener(BLBID, createArgs)
	if err != nil {
		log.Errorf("create https listener error: %+v\n", err)
		return
	}
	log.Info("create https listener success\n")

	updateArgs := &appblb.UpdateAppHTTPSListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "LeastConnection",
		KeepSession:  true,
		CertIds:      []string{"cert-xvysj80uif1y"},
	}
	err = client.UpdateAppHTTPSListener(BLBID, updateArgs)
	if err != nil {
		log.Errorf("update https listener error: %+v\n", err)
		return
	}
	log.Info("update https listener success\n")

	describeArgs := &appblb.DescribeAppListenerArgs{
		ListenerPort: 80,
	}
	describeResult, err := client.DescribeAppHTTPSListeners(BLBID, describeArgs)
	if err != nil {
		log.Errorf("describe https listener error: %+v\n", err)
		return
	}
	log.Info("describe https listener success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppListenersArgs{
		PortList:    []uint16{80},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppListeners(BLBID, deleteArgs)
	if err != nil {
		log.Errorf("delete https listener error: %+v\n", err)
		return
	}
	log.Info("delete https listener success\n")
}

func SSLListenerExample(client *appblb.Client, BLBID string) {
	createArgs := &appblb.CreateAppSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "RoundRobin",
		CertIds:      []string{"cert-xvysj80uif1y"},
	}
	err := client.CreateAppSSLListener(BLBID, createArgs)
	if err != nil {
		log.Errorf("create ssl listener error: %+v\n", err)
		return
	}
	log.Info("create ssl listener success\n")

	updateArgs := &appblb.UpdateAppSSLListenerArgs{
		ClientToken:  getClientToken(),
		ListenerPort: 80,
		Scheduler:    "LeastConnection",
		CertIds:      []string{"cert-xvysj80uif1y"},
	}
	err = client.UpdateAppSSLListener(BLBID, updateArgs)
	if err != nil {
		log.Errorf("update ssl listener error: %+v\n", err)
		return
	}
	log.Info("update ssl listener success\n")

	describeArgs := &appblb.DescribeAppListenerArgs{
		ListenerPort: 80,
	}
	describeResult, err := client.DescribeAppSSLListeners(BLBID, describeArgs)
	if err != nil {
		log.Errorf("describe ssl listener error: %+v\n", err)
		return
	}
	log.Info("describe ssl listener success: %+v\n", describeResult)

	deleteArgs := &appblb.DeleteAppListenersArgs{
		PortList:    []uint16{80},
		ClientToken: getClientToken(),
	}
	err = client.DeleteAppListeners(BLBID, deleteArgs)
	if err != nil {
		log.Errorf("delete ssl listener error: %+v\n", err)
		return
	}
	log.Info("delete ssl listener success\n")
}

func PolicyTest(client *appblb.Client, blbId string) {
	createArgs := &appblb.CreatePolicysArgs{
		ListenerPort: 80,
		ClientToken:  getClientToken(),
		AppPolicyVos: []appblb.AppPolicy{
			{
				Description:      "test policy",
				AppServerGroupId: "",
				BackendPort:      80,
				Priority:         301,
				RuleList: []appblb.AppRule{
					{
						Key:   "*",
						Value: "*",
					},
				},
			},
		},
	}
	err := client.CreatePolicys(blbId, createArgs)
	if err != nil {
		log.Errorf("create policy error: %+v\n", err)
		return
	}
	log.Info("create policy success\n")

	describeArgs := &appblb.DescribePolicysArgs{
		Port: 80,
	}
	describeResult, err := client.DescribePolicys(blbId, describeArgs)
	if err != nil {
		log.Errorf("describe policy error: %+v\n", err)
		return
	}
	log.Info("describe policy success: %+v\n", describeResult)

	deleteArgs := &appblb.DeletePolicysArgs{
		Port:         80,
		PolicyIdList: []string{describeResult.PolicyList[0].Id},
		ClientToken:  getClientToken(),
	}
	err = client.DeletePolicys(blbId, deleteArgs)
	if err != nil {
		log.Errorf("delete policy error: %+v\n", err)
		return
	}
	log.Info("delete policy success\n")
}

func getClientToken() string {
	u, _ := uuid.NewV4()
	return u.String()
}
