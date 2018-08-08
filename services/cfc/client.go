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

// client.go - define the client for CFC service

// Package cfc defines the CFC services of BCE. The supported APIs are all defined in sub-package

package cfc

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cfc/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "cfc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of CFC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the CFC service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func newClient(ak, sk, token, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 || len(sk) == 0 {
		return nil, errors.New("ak sk illegal")
	}
	if len(token) == 0 {
		credentials, err = auth.NewBceCredentials(ak, sk)
	} else {
		credentials, err = auth.NewSessionBceCredentials(ak, sk, token)
	}
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:    endpoint,
		Region:      bce.DEFAULT_REGION,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption:  defaultSignOptions,
		Retry:       bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}
	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

func NewClient(ak, sk, endpoint string) (*Client, error) {
	return newClient(ak, sk, "", endpoint)
}

func NewStsClient(ak, sk, token, endpoint string) (*Client, error) {
	return newClient(ak, sk, token, endpoint)
}

func (c *Client) Invocations(functionName string, payload interface{}, args *api.InvocationsArgs) (
	*api.InvocationResult, error) {
	return api.Invocations(c, functionName, payload, args)
}

func (c *Client) ListFunctions(args *api.ListFunctionsArgs) ([]*api.Function, error) {
	return api.ListFunctions(c, args)
}

func (c *Client) GetFunction(functionName, qualifier string) (*api.GetFunctionResult, error) {
	return api.GetFunction(c, functionName, qualifier)
}

func (c *Client) CreateFunction(args *api.FunctionArgs) (*api.Function, error) {
	return api.CreateFunction(c, args)
}
