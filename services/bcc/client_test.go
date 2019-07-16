package bcc

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bcc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BCC_CLIENT              *Client
	BCC_TestCdsId           string
	BCC_TestBccId           string
	BCC_TestSecurityGroupId string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

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

	BCC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	//log.SetLogLevel(log.DEBUG)
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

//bcc sdk unit test
func TestCreateInstance(t *testing.T) {
	createInstanceArgs := &api.CreateInstanceArgs{
		ImageId: "m-DpgNg8lO",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		InstanceType:        api.InstanceTypeN1,
		CpuCount:            1,
		MemoryCapacityInGB:  1,
		RootDiskSizeInGb:    40,
		RootDiskStorageType: api.StorageTypeCloudHP1,
		CreateCdsList: []api.CreateCdsModel{
			{
				StorageType: api.StorageTypeSSD,
				CdsSizeInGB: 0,
			},
		},
		AdminPass: "123qaz!@#",
		Name:      "terraform_sdkTest",
	}
	createResult, err := BCC_CLIENT.CreateInstance(createInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestListInstances(t *testing.T) {
	listArgs := &api.ListInstanceArgs{}
	_, err := BCC_CLIENT.ListInstances(listArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetInstanceDetail(t *testing.T) {
	_, err := BCC_CLIENT.GetInstanceDetail(BCC_TestCdsId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestResizeInstance(t *testing.T) {
	resizeArgs := &api.ResizeInstanceArgs{
		CpuCount:           2,
		MemoryCapacityInGB: 4,
	}
	err := BCC_CLIENT.ResizeInstance(BCC_TestBccId, resizeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStopInstance(t *testing.T) {
	err := BCC_CLIENT.StopInstance(BCC_TestBccId, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStartInstance(t *testing.T) {
	err := BCC_CLIENT.StartInstance(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebootInstance(t *testing.T) {
	err := BCC_CLIENT.RebootInstance(BCC_TestBccId, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebuildInstance(t *testing.T) {
	rebuildArgs := &api.RebuildInstanceArgs{
		ImageId:   "m-DpgNg8lO",
		AdminPass: "123qaz!@#",
	}
	err := BCC_CLIENT.RebuildInstance(BCC_TestBccId, rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestChangeInstancePass(t *testing.T) {
	changeArgs := &api.ChangeInstancePassArgs{
		AdminPass: "321zaq#@!",
	}
	err := BCC_CLIENT.ChangeInstancePass(BCC_TestBccId, changeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceAttribute(t *testing.T) {
	modifyArgs := &api.ModifyInstanceAttributeArgs{
		Name: "test-modify",
	}
	err := BCC_CLIENT.ModifyInstanceAttribute(BCC_TestBccId, modifyArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetInstanceVNC(t *testing.T) {
	_, err := BCC_CLIENT.GetInstanceVNC(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindSecurityGroup(t *testing.T) {
	err := BCC_CLIENT.BindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnBindSecurityGroup(t *testing.T) {
	err := BCC_CLIENT.UnBindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
	ExpectEqual(t.Errorf, err, nil)
}

//cds sdk unit test
func TestCreateCDSVolume(t *testing.T) {
	args := &api.CreateCDSVolumeArgs{
		PurchaseCount: 1,
		CdsSizeInGB:   40,
		StorageType:   api.StorageTypeSSD,
		Billing: &api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
	}

	result, err := BCC_CLIENT.CreateCDSVolume(args)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestCdsId = result.VolumeIds[0]
}

func TestListCDSVolume(t *testing.T) {
	queryArgs := &api.ListCDSVolumeArgs{}
	_, err := BCC_CLIENT.ListCDSVolume(queryArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetCDSVolumeDetail(t *testing.T) {
	_, err := BCC_CLIENT.GetCDSVolumeDetail(BCC_TestCdsId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAttachCDSVolume(t *testing.T) {
	args := &api.AttachVolumeArgs{
		InstanceId: BCC_TestBccId,
	}

	_, err := BCC_CLIENT.AttachCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDetachCDSVolume(t *testing.T) {
	args := &api.DetachVolumeArgs{
		InstanceId: BCC_TestBccId,
	}

	err := BCC_CLIENT.DetachCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestResizeCDSVolume(t *testing.T) {
	args := &api.ResizeCSDVolumeArgs{
		NewCdsSizeInGB: 100,
	}

	err := BCC_CLIENT.ResizeCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPurchaseReservedCDSVolume(t *testing.T) {
	args := &api.PurchaseReservedCSDVolumeArgs{
		Billing: &api.Billing{
			Reservation: &api.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "Month",
			},
		},
	}

	err := BCC_CLIENT.PurchaseReservedCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRenameCDSVolume(t *testing.T) {
	args := &api.RenameCSDVolumeArgs{
		Name: "testRenamedName",
	}

	err := BCC_CLIENT.RenameCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyCDSVolume(t *testing.T) {
	args := &api.ModifyCSDVolumeArgs{
		CdsName: "modifiedName",
		Desc:    "modifiedDesc",
	}

	err := BCC_CLIENT.ModifyCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyChargeTypeCDSVolume(t *testing.T) {
	args := &api.ModifyChargeTypeCSDVolumeArgs{
		Billing: &api.Billing{
			Reservation: &api.Reservation{
				ReservationLength: 1,
			},
		},
	}

	err := BCC_CLIENT.ModifyChargeTypeCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteCDSVolumeNew(t *testing.T) {
	args := &api.DeleteCSDVolumeArgs{
		AutoSnapshot: "on",
	}

	err := BCC_CLIENT.DeleteCDSVolumeNew(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteCDSVolume(t *testing.T) {
	err := BCC_CLIENT.DeleteCDSVolume(BCC_TestCdsId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteInstance(t *testing.T) {
	err := BCC_CLIENT.DeleteInstance(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}
