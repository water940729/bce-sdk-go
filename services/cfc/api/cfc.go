/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/asaskevich/govalidator"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func Invocations(cli bce.Client, functionName string, payload interface{}, args *InvocationsArgs) (*InvocationResult, error) {
	if err := validateFunctionName(functionName); err != nil {
		return nil, err
	}
	if len(args.InvocationType) == 0 {
		args.InvocationType = InvocationTypeRequestResponse
	}
	if len(args.LogType) == 0 {
		args.LogType = LogTypeNone
	}
	req := &bce.BceRequest{}
	req.SetUri(getInvocationsUri(functionName))
	req.SetMethod(http.POST)
	req.SetParam("invocationType", string(args.InvocationType))
	req.SetParam("logType", string(args.LogType))
	if len(args.Qualifier) > 0 {
		req.SetParam("Qualifier", args.Qualifier)
	}
	if payload != nil {
		var payloadBytes []byte
		var err error
		switch payload.(type) {
		case string:
			payloadBytes = []byte(payload.(string))
		case []byte:
			payloadBytes = payload.([]byte)
		default:
			payloadBytes, err = json.Marshal(payload)
			if err != nil {
				return nil, err
			}
		}
		if !json.Valid(payloadBytes) {
			return nil, errors.New(ParseJsonError)
		}

		requestBody, err := bce.NewBodyFromBytes(payloadBytes)
		if err != nil {
			return nil, err
		}
		req.SetBody(requestBody)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	defer resp.Body().Close()
	body, err := ioutil.ReadAll(resp.Body())
	if err != nil {
		return nil, err
	}
	result := &InvocationResult{
		Payload: string(body),
	}
	errorStr := resp.Header("x-bce-function-error")
	if len(errorStr) > 0 {
		result.FunctionError = errorStr
	}
	logResult := resp.Header("x-bce-log-result")
	if len(logResult) > 0 {
		decodeBytes, err := base64.StdEncoding.DecodeString(string(logResult))
		if err != nil {
			return nil, err
		}
		result.LogResult = string(decodeBytes)
	}
	return result, nil
}

func ListFunctions(cli bce.Client, args *ListFunctionsArgs) ([]*Function, error) {
	if ok, err := govalidator.ValidateStruct(args); !ok {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(getFunctionsUri())
	req.SetMethod(http.GET)
	req.SetParam("FunctionVersion", string(args.FunctionVersion))
	req.SetParam("Marker", string(args.Marker))
	req.SetParam("MaxItems", string(args.MaxItems))
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ListFunctionsResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Functions, nil
}

func GetFunction(cli bce.Client, functionName, qualifier string) (*GetFunctionResult, error) {
	if err := validateFunctionName(functionName); err != nil {
		return nil, err
	}
	if err := validateQualifier(qualifier); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(getFunctionUri(functionName))
	req.SetMethod(http.GET)
	req.SetParam("Qualifier", qualifier)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetFunctionResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateFunction(cli bce.Client, args *FunctionArgs) (*Function, error) {
	if err := validateAndInitCreateFunctionArgs(args); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(getFunctionsUri())
	req.SetMethod(http.POST)
	if args != nil {
		argsBytes, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		requestBody, err := bce.NewBodyFromBytes(argsBytes)
		if err != nil {
			return nil, err
		}
		req.SetBody(requestBody)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &Function{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
