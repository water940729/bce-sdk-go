package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/cfc"
	"github.com/baidubce/bce-sdk-go/services/cfc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"time"
)

func main() {
	log.SetLogLevel(log.DEBUG)
	log.SetLogHandler(log.FILE)
	log.SetLogDir(LogDir)

	client, err := cfc.NewClient(AK, SK, Endpoint)
	if err != nil {
		log.Errorf("create cfc client fail: %+v", err)
		return
	}

	FunctionName01 := fmt.Sprintf("sdktest-function01-%s", time.Now().Format("2006-01-02T150405"))

	zipFileNodejs01 := `UEsDBBQACAAIAAyjX00AAAAAAAAAAAAAAAAIABAAaW5kZXguanNVWAwAsJ/ZW/ie2Vv6Z7qeS60o
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

	zipFileNodejs02 := `UEsDBBQACAAAAHCjX00AAAAAAAAAAAAAAAAJABUAX19NQUNPU1gvVVgIALSf2Vu0n9lbVVQFAAG0
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

	AliasName01 := fmt.Sprintf("sdktest-alias01-%s", time.Now().Format("2006-01-02T150405"))

	// Create a function with the code and configuration
	createFunctionResult, err := client.CreateFunction(&api.CreateFunctionArgs{
		Code:         &api.CodeFile{ZipFile: zipFileNodejs01},
		Publish:      false,
		FunctionName: FunctionName01,
		Handler:      "index.handler",
		Runtime:      "nodejs8.5",
		MemorySize:   256,
		Timeout:      3,
		Description:  "Description",
	})
	log.Debugf("createFunctionResult %+v", *createFunctionResult)

	// Updates a function's code
	updateFunctionCodeResult, err := client.UpdateFunctionCode(&api.UpdateFunctionCodeArgs{
		FunctionName: FunctionName01,
		ZipFile:      zipFileNodejs02,
		Publish:      true,
		DryRun:       false,
	})
	log.Debugf("updateFunctionCodeResult %+v", *updateFunctionCodeResult)

	// Invokes a function synchronously or asynchronously
	invocationsResult, err := client.Invocations(&api.InvocationsArgs{
		FunctionName:   FunctionName01,
		InvocationType: api.InvocationTypeRequestResponse,
		Payload:        nil,
	})
	log.Debugf("invocationsResult %+v", *invocationsResult)

	// Returns configuration of the function
	getFunctionConfigurationResult, err := client.GetFunctionConfiguration(&api.GetFunctionConfigurationArgs{
		FunctionName: FunctionName01,
	})
	log.Debugf("getFunctionConfigurationResult %+v", *getFunctionConfigurationResult)

	// Updates a function's configuration
	updateFunctionConfigurationResult, err := client.UpdateFunctionConfiguration(&api.UpdateFunctionConfigurationArgs{
		FunctionName: FunctionName01,
		Timeout:      5,
		Description:  "hello cfc",
		Handler:      "index.handler",
		Runtime:      "nodejs8.5",
		Environment: &api.Environment{
			Variables: map[string]string{
				"name": "Test",
			},
		},
	})
	log.Debugf("updateFunctionConfigurationResult %+v", *updateFunctionConfigurationResult)

	// Returns information about the function or function version
	getFunctionResult, err := client.GetFunction(&api.GetFunctionArgs{
		FunctionName: FunctionName01,
	})
	functionBRN := getFunctionResult.Configuration.FunctionBrn
	codeSha256 := getFunctionResult.Configuration.CodeSha256
	log.Debugf("BRN %s", functionBRN)
	log.Debugf("codeSha256 %s", codeSha256)
	log.Debugf("getFunctionResult %+v", *getFunctionResult)

	// Returns a list of functions
	listFunctionResult, err := client.ListFunctions(&api.ListFunctionsArgs{
		Marker:   1,
		MaxItems: 2,
	})
	log.Debugf("listFunctionResult %+v", *listFunctionResult)

	// Creates a version from the current code and configuration of a function
	PublishVersionErr := client.PublishVersion(&api.PublishVersionArgs{
		FunctionName: FunctionName01,
		Description:  "go sdk cfc",
		CodeSha256:   codeSha256,
	})
	log.Debugf("PublishVersionErr %+v", PublishVersionErr)

	// Returns a list of function versions
	listVersionResult, err := client.ListVersionsByFunction(&api.ListVersionsByFunctionArgs{
		FunctionName: FunctionName01,
		Marker:       0,
		MaxItems:     2,
	})
	log.Debugf("listVersionResult %+v", *listVersionResult)

	// Create an alias for the function version
	createAliasResult, err := client.CreateAlias(&api.CreateAliasArgs{
		FunctionName:    FunctionName01,
		FunctionVersion: "$LATEST",
		Name:            AliasName01,
		Description:     "test alias",
	})
	log.Debugf("createAliasResult %+v", *createAliasResult)

	// Updates the configuration of a function alias
	updateAliasResult, err := client.UpdateAlias(&api.UpdateAliasArgs{
		FunctionName:    FunctionName01,
		AliasName:       AliasName01,
		FunctionVersion: "$LATEST",
		Description:     "test alias " + AliasName01,
	})
	log.Debugf("updateAliasResult %+v", *updateAliasResult)

	// Returns details about the function alias
	getAliasResult, err := client.GetAlias(&api.GetAliasArgs{
		FunctionName: FunctionName01,
		AliasName:    AliasName01,
	})
	log.Debugf("getAliasResult %+v", *getAliasResult)

	// Returns a list of aliases for the function
	listAliasesResult, err := client.ListAliases(&api.ListAliasesArgs{
		FunctionName:    FunctionName01,
		FunctionVersion: "$LATEST",
		Marker:          0,
		MaxItems:        2,
	})
	log.Debugf("listAliasesResult %+v", *listAliasesResult)

	// Create a trigger to the function
	createTriggerResult, err := client.CreateTrigger(&api.CreateTriggerArgs{
		Target: functionBRN,
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
	})
	var RelationId string
	if err == nil {
		RelationId = createTriggerResult.Relation.RelationId
	}
	log.Debugf("createTriggerResult %+v", *createTriggerResult)

	// Returns a list of trigger to the function
	listTriggerResult, err := client.ListTriggers(&api.ListTriggersArgs{
		FunctionBrn: functionBRN,
	})
	log.Debugf("listTriggerResult %+v", *listTriggerResult)

	// Updates the trigger to the function
	updateTriggerResult, err := client.UpdateTrigger(&api.UpdateTriggerArgs{
		RelationId: RelationId,
		Target:     functionBRN,
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
	})
	RelationId = updateTriggerResult.Relation.RelationId
	log.Debugf("updateTriggerResult %+v", updateTriggerResult)

	// Delete a trigger to the function
	DeleteTriggerErr := client.DeleteTrigger(&api.DeleteTriggerArgs{
		RelationId: RelationId,
		Target:     functionBRN,
		Source:     api.SourceTypeHTTP,
	})
	log.Debugf("DeleteTriggerErr %+v", DeleteTriggerErr)

	// Delete a function alias
	DeleteAliasErr := client.DeleteAlias(&api.DeleteAliasArgs{
		FunctionName: FunctionName01,
		AliasName:    AliasName01,
	})
	log.Debugf("DeleteAliasErr %+v", DeleteAliasErr)

	// Deletes a function
	DeleteFunctionErr := client.DeleteFunction(&api.DeleteFunctionArgs{
		FunctionName: FunctionName01,
	})
	log.Debugf("DeleteFunctionErr %+v", DeleteFunctionErr)

}
