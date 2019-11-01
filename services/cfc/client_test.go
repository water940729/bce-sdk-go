package cfc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/services/cfc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CfcClient       *Client
	FunctionName01  string
	FunctionName02  string
	AliasName01     string
	AliasName02     string
	FunctionBRN     string
	CodeSha256      string
	RelationId      string
	zipFilePython   string
	zipFileNodejs01 string
	zipFileNodejs02 string
)

const (
	invokeTestReturnPayload = "Hello World"
)

var (
	logSuccess = true
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

type PayloadDemo struct {
	A string
	B int
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		fmt.Printf("config json file of ak/sk not given:(%s) err(%v)\n", conf, err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)
	FunctionName01 = fmt.Sprintf("sdktest-function01-%s", time.Now().Format("2006-01-02T150405"))
	FunctionName02 = fmt.Sprintf("sdktest-function02-%s", time.Now().Format("2006-01-02T150405"))
	zipFilePython = `UEsDBBQACAAIAAAAAAAAAAAAAAAAAAAAAAAIAAAAaW5kZXgucHlKSU1TyEjMS8lJLdJILUvNK9FRSM
7PK0mtKNG0UuBSUFBQKCjKzCuByGmCBYpSS0qL8hSUPFJzcvIVwvOLclKUuAABAAD//1BLBwhwCJ
1tRwAAAEgAAABQSwECFAAUAAgACAAAAAAAcAidbUcAAABIAAAACAAAAAAAAAAAAAAAAAAAAAAAaW
5kZXgucHlQSwUGAAAAAAEAAQA2AAAAfQAAAAAA`

	zipFileNodejs01 = `UEsDBBQACAAIAAyjX00AAAAAAAAAAAAAAAAIABAAaW5kZXguanNVWAwAsJ/ZW/ie2Vv6Z7qeS60o
yC8qKdbLSMxLyUktUrBV0EgtS80r0VFIzs8rSa0AMRJzcpISk7M1FWztFKq5FIAAJqSRV5qTo6Og
5JGak5OvUJ5flJOiqKRpzVVrDQBQSwcILzRMjVAAAABYAAAAUEsDBAoAAAAAAHCjX00AAAAAAAAA
AAAAAAAJABAAX19NQUNPU1gvVVgMALSf2Vu0n9lb+me6nlBLAwQUAAgACAAMo19NAAAAAAAAAAAA
AAAAEwAQAF9fTUFDT1NYLy5faW5kZXguanNVWAwAsJ/ZW/ie2Vv6Z7qeY2AVY2dgYmDwTUxW8A9W
iFCAApAYAycQGwFxHRCD+BsYiAKOISFBUCZIxwIgFkBTwogQl0rOz9VLLCjISdXLSSwuKS1OTUlJ
LElVDggGKXw772Y0iO5J8tAH0QBQSwcIDgnJLFwAAACwAAAAUEsBAhUDFAAIAAgADKNfTS80TI1Q
AAAAWAAAAAgADAAAAAAAAAAAQKSBAAAAAGluZGV4LmpzVVgIALCf2Vv4ntlbUEsBAhUDCgAAAAAA
cKNfTQAAAAAAAAAAAAAAAAkADAAAAAAAAAAAQP1BlgAAAF9fTUFDT1NYL1VYCAC0n9lbtJ/ZW1BL
AQIVAxQACAAIAAyjX00OCcksXAAAALAAAAATAAwAAAAAAAAAAECkgc0AAABfX01BQ09TWC8uX2lu
ZGV4LmpzVVgIALCf2Vv4ntlbUEsFBgAAAAADAAMA0gAAAHoBAAAAAA==`

	zipFileNodejs02 = `UEsDBBQACAAAAHCjX00AAAAAAAAAAAAAAAAJABUAX19NQUNPU1gvVVgIALSf2Vu0n9lbVVQFAAG0
n9lbUEsHCAAAAAAAAAAAAAAAAFBLAwQUAAgACAAMo19NAAAAAAAAAAAAAAAAEwAVAF9fTUFDT1NY
Ly5faW5kZXguanNVWAgAsJ/ZW/ie2VtVVAUAAfie2VtiYBVjZ2BiYPBNTFbwD1aIUIACkBgDJwMD
gxEDA0MdAwOYv4GBKOAYEhIEZYJ0LGBgYBBAU8KIEJdKzs/VSywoyEnVy0ksLiktTk1JSSxJVQ4I
Bil8O+9mNIjuSfLQB9GAAAAA//9QSwcIDgnJLGYAAACwAAAAUEsDBBQACAAIAAAAAAAAAAAAAAAA
AAAAAAAIAAAAaW5kZXguanNKrSjILyop1stIzEvJSS1SsFXQSC1LzSvRUUjOzytJrQAxEnNykhKT
szUVbO0UqrkUFBTgQhp5pTk5OgpKHqk5OfkKzm7Oikqa1ly11oAAAAD//1BLBwg7znNyUwAAAFYA
AABQSwECFAMUAAgAAABwo19NAAAAAAAAAAAAAAAACQAVAAAAAAAAAABA/UEAAAAAX19NQUNPU1gv
VVgIALSf2Vu0n9lbVVQFAAG0n9lbUEsBAhQDFAAIAAgADKNfTQ4JySxmAAAAsAAAABMAFQAAAAAA
AAAAQKSBTAAAAF9fTUFDT1NYLy5faW5kZXguanNVWAgAsJ/ZW/ie2VtVVAUAAfie2VtQSwECFAAU
AAgACAAAAAAAO85zclMAAABWAAAACAAAAAAAAAAAAAAAAAAIAQAAaW5kZXguanNQSwUGAAAAAAMA
AwDYAAAAkQEAAAAA`

	AliasName01 = fmt.Sprintf("sdktest-alias01-%s", time.Now().Format("2006-01-02T150405"))
	AliasName02 = fmt.Sprintf("sdktest-alias02-%s", time.Now().Format("2006-01-02T150405"))

	CfcClient, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		panic(err)
	}
	log.SetLogHandler(log.FILE)
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
	log.SetLogLevel(log.DEBUG)
}

func TestCreateFunction(t *testing.T) {
	// This function return Hello World
	cases := []api.CreateFunctionArgs{
		{
			Code:         &api.CodeFile{ZipFile: zipFilePython},
			Publish:      true,
			FunctionName: FunctionName01,
			Handler:      "index.handler",
			Runtime:      "python2",
			MemorySize:   128,
			Timeout:      3,
			Description:  "Description",
		},
		{
			Code:         &api.CodeFile{ZipFile: zipFileNodejs01},
			Publish:      false,
			FunctionName: FunctionName02,
			Handler:      "index.handler",
			Runtime:      "nodejs8.5",
			MemorySize:   256,
			Timeout:      3,
			Description:  "Description",
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateFunction(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestListFunctions(t *testing.T) {
	args := &api.ListFunctionsArgs{
		FunctionVersion: "1",
		Marker:          1,
		MaxItems:        2,
	}
	res, err := CfcClient.ListFunctions(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestGetFunction(t *testing.T) {
	res, err := CfcClient.GetFunction(&api.GetFunctionArgs{
		FunctionName: FunctionName01,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	FunctionBRN = res.Configuration.FunctionBrn
	fmt.Printf(FunctionBRN)
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestInvocations(t *testing.T) {
	cases := []struct {
		args        *api.InvocationsArgs
		respPayload string
		err         error
	}{
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        nil,
			},
			respPayload: invokeTestReturnPayload,
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeEvent,
				Payload:        `[{"a":1},{"a":2}]`,
			},
			respPayload: "",
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        `[{"a":,{"a":2}]`,
			},
			respPayload: "",
			err:         fmt.Errorf("could not parse payload into json"),
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeEvent,
				Payload:        []byte(`{"a":1}`),
			},
			respPayload: "",
			err:         nil,
		},
		{
			args: &api.InvocationsArgs{
				FunctionName:   FunctionName01,
				InvocationType: api.InvocationTypeRequestResponse,
				Payload:        []*PayloadDemo{&PayloadDemo{A: "1", B: 2}, &PayloadDemo{A: "3", B: 4}},
			},
			respPayload: invokeTestReturnPayload,
			err:         nil,
		},
	}
	for _, tc := range cases {
		t.Run("invoke", func(t *testing.T) {
			res, err := CfcClient.Invocations(tc.args)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			} else if res != nil && res.Payload != tc.respPayload {
				t.Errorf("Expected Payload to be %s, but got %s", tc.respPayload, res.Payload)
			}
		})
	}
}

func TestUpdateFunctionCode(t *testing.T) {
	res, err := CfcClient.UpdateFunctionCode(&api.UpdateFunctionCodeArgs{
		FunctionName: FunctionName02,
		ZipFile:      zipFileNodejs02,
		Publish:      false,
		DryRun:       false,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateFunctionConfiguration(t *testing.T) {
	res, err := CfcClient.UpdateFunctionConfiguration(&api.UpdateFunctionConfigurationArgs{
		FunctionName: FunctionName02,
		Timeout:      5,
		Description:  "wooo cool",
		Handler:      "index.handler",
		Runtime:      "nodejs8.5",
		Environment: &api.Environment{
			Variables: map[string]string{
				"name": "Test",
			},
		},
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestGetFunctionConfiguration(t *testing.T) {
	res, err := CfcClient.GetFunctionConfiguration(&api.GetFunctionConfigurationArgs{
		FunctionName: FunctionName02,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestPublishVersion(t *testing.T) {
	res, err := CfcClient.GetFunction(&api.GetFunctionArgs{
		FunctionName: FunctionName02,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	CodeSha256 = res.Configuration.CodeSha256
	fmt.Printf(FunctionBRN)
	err = CfcClient.PublishVersion(&api.PublishVersionArgs{
		FunctionName: FunctionName02,
		Description:  "test",
		CodeSha256:   CodeSha256,
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestListVersionsByFunction(t *testing.T) {
	res, err := CfcClient.ListVersionsByFunction(&api.ListVersionsByFunctionArgs{
		FunctionName: "testHelloWorld",
	})
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestCreateAlias(t *testing.T) {
	cases := []api.CreateAliasArgs{
		{
			FunctionName:    FunctionName02,
			FunctionVersion: "$LATEST",
			Name:            AliasName01,
			Description:     "test alias",
		},
		{
			FunctionName:    FunctionName02,
			FunctionVersion: "$LATEST",
			Name:            AliasName02,
			Description:     "test alias",
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateAlias(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}
		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestGetAlias(t *testing.T) {
	args := &api.GetAliasArgs{
		FunctionName: FunctionName02,
		AliasName:    AliasName01,
	}
	res, err := CfcClient.GetAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateAlias(t *testing.T) {
	args := &api.UpdateAliasArgs{
		FunctionName:    FunctionName02,
		AliasName:       AliasName01,
		FunctionVersion: "$LATEST",
		Description:     "test alias " + AliasName01,
	}
	res, err := CfcClient.UpdateAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestListAliases(t *testing.T) {
	args := &api.ListAliasesArgs{
		FunctionName:    FunctionName02,
		FunctionVersion: "$LATEST",
		Marker:          0,
		MaxItems:        2,
	}
	res, err := CfcClient.ListAliases(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteAlias(t *testing.T) {
	args := &api.DeleteAliasArgs{
		FunctionName: FunctionName02,
		AliasName:    AliasName02,
	}
	err := CfcClient.DeleteAlias(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestCreateTrigger(t *testing.T) {
	cases := []api.CreateTriggerArgs{
		{
			Target: FunctionBRN,
			Source: api.SourceTypeHTTP,
			Data: struct {
				ResourcePath string
				Method       string
				AuthType     string
			}{
				ResourcePath: fmt.Sprintf("tr01-%s", time.Now().Format("2006-01-02T150405")),
				Method:       "GET",
				AuthType:     "anonymous",
			},
		}, {
			Target: FunctionBRN,
			Source: api.SourceTypeHTTP,
			Data: struct {
				ResourcePath string
				Method       string
				AuthType     string
			}{
				ResourcePath: fmt.Sprintf("tr02-%s", time.Now().Format("2006-01-02T150405")),
				Method:       "GET",
				AuthType:     "anonymous",
			},
		},
	}
	for _, args := range cases {
		res, err := CfcClient.CreateTrigger(&args)
		if err != nil {
			t.Fatalf("err (%v)", err)
		}

		resStr, err := json.MarshalIndent(res, "", "	")
		if logSuccess && err == nil {
			t.Logf("res %s ", resStr)
		}
	}
}

func TestListTriggers(t *testing.T) {
	args := &api.ListTriggersArgs{
		FunctionBrn: FunctionBRN,
	}
	res, err := CfcClient.ListTriggers(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if len(res.Relation) > 0 {
		RelationId = res.Relation[0].RelationId
	}
	t.Logf("res %v", res)
	resStr, err := json.Marshal(res)
	if err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestUpdateTrigger(t *testing.T) {
	args := &api.UpdateTriggerArgs{
		RelationId: RelationId,
		Target:     FunctionBRN,
		Source:     api.SourceTypeHTTP,
		Data: struct {
			ResourcePath string
			Method       string
			AuthType     string
		}{
			ResourcePath: fmt.Sprintf("tr99-%s", time.Now().Format("2006-01-02T150405")),
			Method:       "GET",
			AuthType:     "anonymous",
		},
	}
	res, err := CfcClient.UpdateTrigger(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if res.Relation != nil {
		RelationId = res.Relation.RelationId
	}
	resStr, err := json.MarshalIndent(res, "", "	")
	if logSuccess && err == nil {
		t.Logf("res %s ", resStr)
	}
}

func TestDeleteTrigger(t *testing.T) {
	listArgs := &api.ListTriggersArgs{
		FunctionBrn: FunctionBRN,
	}
	res, err := CfcClient.ListTriggers(listArgs)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	if len(res.Relation) > 0 {
		RelationId = res.Relation[0].RelationId
	}
	args := &api.DeleteTriggerArgs{
		RelationId: RelationId,
		Target:     FunctionBRN,
		Source:     api.SourceTypeHTTP,
	}
	t.Logf("args (%+v)", args)
	err = CfcClient.DeleteTrigger(args)
	if err != nil {
		t.Errorf("err (%v)", err)
	}
}

func TestDeleteFunction(t *testing.T) {
	args := &api.DeleteFunctionArgs{
		FunctionName: FunctionName01,
	}
	err := CfcClient.DeleteFunction(args)
	if err != nil {
		t.Logf("res (%v)", err)
	}
}
