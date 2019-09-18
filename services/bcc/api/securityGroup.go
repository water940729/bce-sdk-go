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

// securityGroup.go - the securityGroup APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateSecurityGroup(cli bce.Client, args *CreateSecurityGroupArgs) (*CreateSecurityGroupResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSecurityGroupUri())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateSecurityGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListSecurityGroup(cli bce.Client, queryArgs *ListSecurityGroupArgs) (*ListSecurityGroupResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSecurityGroupUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.InstanceId) != 0 {
			req.SetParam("instanceId", queryArgs.InstanceId)
		}
		if len(queryArgs.VpcId) != 0 {
			req.SetParam("vpcId", queryArgs.VpcId)
		}
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
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

	jsonBody := &ListSecurityGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func AuthorizeSecurityGroupRule(cli bce.Client, securityGroupId string, args *AuthorizeSecurityGroupArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSecurityGroupUriWithId(securityGroupId))
	req.SetMethod(http.PUT)

	req.SetParam("authorizeRule", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

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

func RevokeSecurityGroupRule(cli bce.Client, securityGroupId string, args *RevokeSecurityGroupArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSecurityGroupUriWithId(securityGroupId))
	req.SetMethod(http.PUT)

	req.SetParam("revokeRule", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

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

func DeleteSecurityGroup(cli bce.Client, securityGroupId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSecurityGroupUriWithId(securityGroupId))
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
