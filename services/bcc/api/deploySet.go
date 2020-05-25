/*
 * Copyright 2020 Baidu, Inc.
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

// deploySet.go - the deploy set APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package api

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
//     - reqBody: http request body
// RETURNS:
//     - *CreateDeploySetResult: results of creating a deploy set
//     - error: nil if success otherwise the specific error
func CreateDeploySet(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateDeploySetResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetCreateUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateDeploySetResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListDeploySets - list all deploy sets
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
// RETURNS:
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func ListDeploySets(cli bce.Client) (*ListDeploySetsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetListUri())
	req.SetMethod(http.GET)

	fmt.Println(req)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDeploySetsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ModifyDeploySet - modify the deploy set atrribute
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyDeploySet(cli bce.Client, deploySetId string, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(deploySetId))
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)
	req.SetParam("modifyAttribute", "")

	// Send request and get response
	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

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

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteDeploySet(cli bce.Client, deploySetId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(deploySetId))
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

func getDeploySetCreateUri() string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + REQUEST_DEPLOYSET_URI + REQUEST_CREATE_URI
}

func getDeploySetListUri() string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + REQUEST_DEPLOYSET_URI + REQUEST_LIST_URI
}

func getDeploySetUriWithId(id string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + REQUEST_DEPLOYSET_URI + "/" + id
}
