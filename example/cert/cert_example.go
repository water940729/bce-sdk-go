package main

import (
	"github.com/baidubce/bce-sdk-go/services/cert"
	"github.com/baidubce/bce-sdk-go/util/log"
)

const (
	AK       = ""
	SK       = ""
	EndPoint = "certificate.baidubce.com"
	LogDir   = "./logs/"
)

var (
	testCertServerData  = ``
	testCertPrivateData = ``

	testUpdateCertServerData  = ``
	testUpdateCertPrivateData = ``
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.STDOUT)

	client, err := cert.NewClient(AK, SK, EndPoint)
	if err != nil {
		log.Errorf("create client failed: %+v\n", err)
		return
	}

	createArgs := &cert.CreateCertArgs{
		CertName:        "sdkcreateTest",
		CertServerData:  testCertServerData,
		CertPrivateData: testCertPrivateData,
	}
	createResult, err := client.CreateCert(createArgs)
	if err != nil {
		log.Errorf("create cert error: %+v\n", err)
		return
	}
	log.Info("create cert success: %+v\n", createResult)

	updateArgs := &cert.UpdateCertNameArgs{
		CertName: "test-sdk-cert",
	}
	err = client.UpdateCertName(createResult.CertId, updateArgs)
	if err != nil {
		log.Errorf("update cert error: %+v\n", err)
		return
	}
	log.Info("update cert success\n")

	listResult, err := client.ListCerts()
	if err != nil {
		log.Errorf("list certs error: %+v\n", err)
		return
	}
	log.Info("list certs success: %+v\n", listResult)

	updateDataArgs := &cert.UpdateCertDataArgs{
		CertName:"test-sdk-cert",
		CertServerData:testUpdateCertServerData,
		CertPrivateData:testUpdateCertPrivateData,
	}
	err = client.UpdateCertData(createResult.CertId, updateDataArgs)
	if err != nil {
		log.Errorf("update cert data error: %+v\n", err)
		return
	}
	log.Info("update cert data success\n")

	meta, err := client.GetCertMeta(createResult.CertId)
	if err != nil {
		log.Errorf("get certs meta error: %+v\n", err)
		return
	}
	log.Info("get certs meta success: %+v\n", meta)

	err = client.DeleteCert(createResult.CertId)
	if err != nil {
		log.Errorf("delete certs error: %+v\n", err)
		return
	}
	log.Info("delete certs success\n")
}
