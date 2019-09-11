package appblb

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
	"github.com/gofrs/uuid"
)

var (
	APPBLB_CLIENT             *Client
	APPBLB_ID                 string
	APPBLB_SERVERGROUP_ID     string
	APPBLB_SERVERGROUPPORT_ID string

	// set these values before start test
	VPC_TEST_ID    = ""
	SUBNET_TEST_ID = ""
	INSTANCE_ID    = ""
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	// run 7 times is not necessary, just for make sure we can get work directory path
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	APPBLB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_CreateLoadBalancer(t *testing.T) {
	createArgs := &CreateLoadBalancerArgs{
		ClientToken: getClientToken(),
		Name:        "sdkBlb",
		VpcId:       VPC_TEST_ID,
		SubnetId:    SUBNET_TEST_ID,
	}

	createResult, err := APPBLB_CLIENT.CreateLoadBalancer(createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_ID = createResult.BlbId
}

func TestClient_UpdateLoadBalancer(t *testing.T) {
	updateArgs := &UpdateLoadBalancerArgs{
		Name:        "testSdk",
		Description: "test desc",
	}
	err := APPBLB_CLIENT.UpdateLoadBalancer(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLoadBalancers(t *testing.T) {
	describeArgs := &DescribeLoadBalancersArgs{}
	_, err := APPBLB_CLIENT.DescribeLoadBalancers(describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeLoadBalancerDetail(t *testing.T) {
	_, err := APPBLB_CLIENT.DescribeLoadBalancerDetail(APPBLB_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppServerGroup(t *testing.T) {
	createArgs := &CreateAppServerGroupArgs{
		ClientToken: getClientToken(),
		Name:        "sdkTest",
	}
	createResult, err := APPBLB_CLIENT.CreateAppServerGroup(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_SERVERGROUP_ID = createResult.Id
}

func TestClient_UpdateAppServerGroup(t *testing.T) {
	updateArgs := &UpdateAppServerGroupArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		Name:        "testSdk",
		Description: "test desc",
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.UpdateAppServerGroup(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeAppServerGroup(t *testing.T) {
	describeArgs := &DescribeAppServerGroupArgs{}
	_, err := APPBLB_CLIENT.DescribeAppServerGroup(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAppServerGroupPort(t *testing.T) {
	createArgs := &CreateAppServerGroupPortArgs{
		ClientToken: getClientToken(),
		SgId:        APPBLB_SERVERGROUP_ID,
		Port:        80,
		Type:        "TCP",
	}
	createResult, err := APPBLB_CLIENT.CreateAppServerGroupPort(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)

	APPBLB_SERVERGROUPPORT_ID = createResult.Id
}

func TestClient_UpdateAppServerGroupPort(t *testing.T) {
	updateArgs := &UpdateAppServerGroupPortArgs{
		ClientToken:                 getClientToken(),
		SgId:                        APPBLB_SERVERGROUP_ID,
		PortId:                      APPBLB_SERVERGROUPPORT_ID,
		HealthCheck:                 "TCP",
		HealthCheckPort:             30,
		HealthCheckIntervalInSecond: 10,
		HealthCheckTimeoutInSecond:  10,
	}
	err := APPBLB_CLIENT.UpdateAppServerGroupPort(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppServerGroupPort(t *testing.T) {
	deleteArgs := &DeleteAppServerGroupPortArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		PortIdList:  []string{APPBLB_SERVERGROUPPORT_ID},
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppServerGroupPort(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateBlbRs(t *testing.T) {
	createArgs := &CreateBlbRsArgs{
		BlbRsWriteOpArgs: BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        APPBLB_SERVERGROUP_ID,
			BackendServerList: []AppBackendServer{
				{InstanceId: INSTANCE_ID, Weight: 30},
			},
		},
	}
	err := APPBLB_CLIENT.CreateBlbRs(APPBLB_ID, createArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateBlbRs(t *testing.T) {
	updateArgs := &UpdateBlbRsArgs{
		BlbRsWriteOpArgs: BlbRsWriteOpArgs{
			ClientToken: getClientToken(),
			SgId:        APPBLB_SERVERGROUP_ID,
			BackendServerList: []AppBackendServer{
				{InstanceId: INSTANCE_ID, Weight: 50},
			},
		},
	}
	err := APPBLB_CLIENT.UpdateBlbRs(APPBLB_ID, updateArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeBlbRs(t *testing.T) {
	describeArgs := &DescribeBlbRsArgs{
		SgId: APPBLB_SERVERGROUP_ID,
	}
	_, err := APPBLB_CLIENT.DescribeBlbRs(APPBLB_ID, describeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteBlbRs(t *testing.T) {
	deleteArgs := &DeleteBlbRsArgs{
		SgId:                APPBLB_SERVERGROUP_ID,
		BackendServerIdList: []string{INSTANCE_ID},
		ClientToken:         getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteBlbRs(APPBLB_ID, deleteArgs)

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeRsMount(t *testing.T) {
	_, err := APPBLB_CLIENT.DescribeRsMount(APPBLB_ID, APPBLB_SERVERGROUP_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DescribeRsUnMount(t *testing.T) {
	_, err := APPBLB_CLIENT.DescribeRsUnMount(APPBLB_ID, APPBLB_SERVERGROUP_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAppServerGroup(t *testing.T) {
	deleteArgs := &DeleteAppServerGroupArgs{
		SgId:        APPBLB_SERVERGROUP_ID,
		ClientToken: getClientToken(),
	}
	err := APPBLB_CLIENT.DeleteAppServerGroup(APPBLB_ID, deleteArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteLoadBalancer(t *testing.T) {
	err := APPBLB_CLIENT.DeleteLoadBalancer(APPBLB_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	u, _ := uuid.NewV4()
	return u.String()
}
