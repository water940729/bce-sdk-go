package cfc

import (
	"encoding/json"
	"errors"
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
	CfcClient    *Client
	FunctionName string
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
	FunctionName = fmt.Sprintf("sl%s", time.Now().Format("2006-01-02T150405"))
	CfcClient, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		panic(err)
	}
	log.SetLogLevel(log.WARN)
}

func TestCreateFunction(t *testing.T) {
	zipFile := "UEsDBBQACAAIAAAAAAAAAAAAAAAAAAAAAAAIAAAAaW5kZXgucHlKSU1TyEjMS8lJLdJILUvNK9FRSM7PK0mtKNG0UuBSUFBQKCjKzCuByGmCBYpSS0qL8hSUPFJzcvIVwvOLclKUuAABAAD//1BLBwhwCJ1tRwAAAEgAAABQSwECFAAUAAgACAAAAAAAcAidbUcAAABIAAAACAAAAAAAAAAAAAAAAAAAAAAAaW5kZXgucHlQSwUGAAAAAAEAAQA2AAAAfQAAAAAA"
	args := &api.FunctionArgs{
		Code:         &api.CodeFile{ZipFile: zipFile},
		Publish:      false,
		FunctionName: FunctionName,
		Handler:      "index.handler",
		Runtime:      "python2",
		MemorySize:   128,
		Timeout:      3,
		Description:  "Description",
	}
	res, err := CfcClient.CreateFunction(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v ", res)
}

func TestListFunctions(t *testing.T) {
	args := &api.ListFunctionsArgs{}
	args.FunctionVersion = "ALL"
	_, err := CfcClient.ListFunctions(args)
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
}

func TestGetFunction(t *testing.T) {
	res, err := CfcClient.GetFunction(FunctionName, "")
	if err != nil {
		t.Fatalf("err (%v)", err)
	}
	t.Logf("res %+v ", res)
}

func TestInvocations(t *testing.T) {
	cases := []struct {
		name    string
		payload interface{}
		err     error
	}{
		{
			payload: nil,
			err:     nil,
		},
		{
			payload: `[{"a":1},{"a":2}]`,
			err:     nil,
		},
		{
			payload: `[{"a":,{"a":2}]`,
			err:     errors.New(api.ParseJsonError),
		},
		{
			payload: []byte(`{"a":1}`),
			err:     nil,
		},
		{
			payload: &PayloadDemo{A: "1", B: 2},
			err:     nil,
		},
		{
			payload: []*PayloadDemo{&PayloadDemo{A: "1", B: 2}, &PayloadDemo{A: "3", B: 4}},
			err:     nil,
		},
	}
	for _, tc := range cases {
		t.Run("invoke", func(t *testing.T) {
			args := &api.InvocationsArgs{}
			args.InvocationType = api.InvocationTypeRequestResponse
			args.LogType = api.LogTypeTail
			res, err := CfcClient.Invocations(FunctionName, tc.payload, args)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			}
			if err == nil {
				t.Logf("res log(%s)\n", res.LogResult)
				t.Logf("res payload(%s)\n", res.Payload)
			}
		})
	}
}
