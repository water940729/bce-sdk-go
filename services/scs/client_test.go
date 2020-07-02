package scs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	SCS_CLIENT  *Client
	SCS_TEST_ID  string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

const (
	SDK_NAME_PREFIX = "sdk_scs_"
)

func init() {
	_, f, _, _ := runtime.Caller(0)
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

	SCS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateInstance(t *testing.T) {
	id := strconv.FormatInt(time.Now().Unix(),10)
	args := &CreateInstanceArgs{
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation: &Reservation{
				ReservationLength: 1,
			},
		},
		ClientToken:   getClientToken(),
		PurchaseCount: 1,
		InstanceName:  SDK_NAME_PREFIX + id,
		Port:          6379,
		EngineVersion: "3.2",
		NodeType:      "cache.n1.micro",
		ClusterType:   "master_slave",
		ReplicationNum:  1,
		ShardNum:      1,
		ProxyNum:      0,
	}
	result, err := SCS_CLIENT.CreateInstance(args)
	ExpectEqual(t.Errorf, nil, err)

	if len(result.InstanceIds) > 0 {
		SCS_TEST_ID = result.InstanceIds[0]
	}
}

func TestClient_ListInstances(t *testing.T) {
	args := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if e.InstanceID == SCS_TEST_ID {
			ExpectEqual(t.Errorf, "Postpaid", e.PaymentTiming)
		}
	}
}

func TestClient_GetInstanceDetail(t *testing.T) {
	_, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInstanceName(t *testing.T) {
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &UpdateInstanceNameArgs{
				InstanceName:  e.InstanceName + "_new",
				ClientToken:   getClientToken(),
			}
			err := SCS_CLIENT.UpdateInstanceName(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ResizeInstance(t *testing.T) {
	args := &ResizeInstanceArgs{
		NodeType:"cache.n1.small",
		ShardNum:2,
		ClientToken:  getClientToken(),
	}
	result, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
	if result.InstanceStatus == "Running" {
		err := SCS_CLIENT.ResizeInstance(SCS_TEST_ID, args)
		ExpectEqual(t.Errorf, nil, err)
	}
}

func TestClient_DeleteInstance(t *testing.T) {
	args := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus && "Postpaid" == e.PaymentTiming {
			err := SCS_CLIENT.DeleteInstance(e.InstanceID, getClientToken())
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetNodeTypeList(t *testing.T) {
	_, err := SCS_CLIENT.GetNodeTypeList()
	ExpectEqual(t.Errorf, nil, err)
}


func getClientToken() string {
	return util.NewUUID()
}
