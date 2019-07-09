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

// instance.go - the instance APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateInstance(cli bce.Client, reqBody *bce.Body) (*CreateInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListInstances(cli bce.Client, args *ListInstanceArgs) (*ListInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.DedicatedHostId) != 0 {
			req.SetParam("dedicateHostId", args.DedicatedHostId)
		}
		if len(args.InternalIp) != 0 {
			req.SetParam("internalIp", args.InternalIp)
		}
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func GetInstanceDetail(cli bce.Client, instanceId string) (*GetInstanceDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetInstanceDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func DeleteInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.DELETE)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func ResizeInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("resize", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func RebuildInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("rebuild", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func StartInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("start", "")

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func StopInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("stop", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func RebootInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("reboot", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func ChangeInstancePass(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("changePass", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func ModifyInstanceAttribute(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("modifyAttribute", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func BindSecurityGroup(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("bind", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func UnBindSecurityGroup(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("unbind", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func GetInstanceVNC(cli bce.Client, instanceId string) (*GetInstanceVNCResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceVNCUri(instanceId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetInstanceVNCResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func InstancePurchaseReserved(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("purchaseReserved", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}
