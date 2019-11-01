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
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
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

// 调用函数 invocations POST /v1/functions/{FunctionName}/invocations
func (c *Client) Invocations(args *api.InvocationsArgs) (*api.InvocationsResult, error) {
	return api.Invocations(c, args)
}

// 调用函数 Invoke POST /v1/functions/{FunctionName}/invocations
func (c *Client) Invoke(args *api.InvocationsArgs) (*api.InvocationsResult, error) {
	return api.Invocations(c, args)
}

// 查询用户函数 ListFunctions GET /v1/functions
func (c *Client) ListFunctions(args *api.ListFunctionsArgs) (*api.ListFunctionsResult, error) {
	return api.ListFunctions(c, args)
}

// 查询单个用户函数 GetFunction GET /v1/functions/{FunctionName}
func (c *Client) GetFunction(args *api.GetFunctionArgs) (*api.GetFunctionResult, error) {
	return api.GetFunction(c, args)
}

// 创建函数 CreateFunction POST /v1/functions
func (c *Client) CreateFunction(args *api.CreateFunctionArgs) (*api.CreateFunctionResult, error) {
	return api.CreateFunction(c, args)
}

// 删除函数 DeleteFunction DELETE /v1/functions/{FunctionName}
func (c *Client) DeleteFunction(args *api.DeleteFunctionArgs) error {
	return api.DeleteFunction(c, args)
}

// 更新指定函数的代码 UpdateFunctionCode PUT /v1/functions/{FunctionName}/code
func (c *Client) UpdateFunctionCode(args *api.UpdateFunctionCodeArgs) (*api.UpdateFunctionCodeResult, error) {
	return api.UpdateFunctionCode(c, args)
}

// 获取指定函数配置 GetFunctionConfiguration GET /v1/functions/{FunctionName}/configuration
func (c *Client) GetFunctionConfiguration(args *api.GetFunctionConfigurationArgs) (*api.GetFunctionConfigurationResult, error) {
	return api.GetFunctionConfiguration(c, args)
}

// 修改函数配置 UpdateFunctionConfiguration PUT /v1/functions/{FunctionName}/configuration
func (c *Client) UpdateFunctionConfiguration(args *api.UpdateFunctionConfigurationArgs) (*api.UpdateFunctionConfigurationResult, error) {
	return api.UpdateFunctionConfiguration(c, args)
}

// 查询函数所有版本 ListVersionsByFunction GET /v1/functions/{FunctionName}/versions
func (c *Client) ListVersionsByFunction(args *api.ListVersionsByFunctionArgs) (*api.ListVersionsByFunctionResult, error) {
	return api.ListVersionsByFunction(c, args)
}

// 发布函数版本 PublishVersion POST /v1/functions/{FunctionName}/versions
func (c *Client) PublishVersion(args *api.PublishVersionArgs) error {
	return api.PublishVersion(c, args)
}

// 查询函数所有别名 ListAliases GET /v1/functions/{FunctionName}/aliases
func (c *Client) ListAliases(args *api.ListAliasesArgs) (*api.ListAliasesResult, error) {
	return api.ListAliases(c, args)
}

// 创建别名 CreateAlias POST /v1/functions/{FunctionName}/aliases
func (c *Client) CreateAlias(args *api.CreateAliasArgs) (*api.CreateAliasResult, error) {
	return api.CreateAlias(c, args)
}

// 查询别名详情 GetAlias GET /v1/functions/{FunctionName}/aliases/{aliasName}/
func (c *Client) GetAlias(args *api.GetAliasArgs) (*api.GetAliasResult, error) {
	return api.GetAlias(c, args)
}

// 修改别名 UpdateAlias PUT /v1/functions/{FunctionName}/aliases/{aliaName}/
func (c *Client) UpdateAlias(args *api.UpdateAliasArgs) (*api.UpdateAliasResult, error) {
	return api.UpdateAlias(c, args)
}

// 删除别名 DeleteAlias DELETE /v1/functions/{FunctionName}/aliases/{aliaName}
func (c *Client) DeleteAlias(args *api.DeleteAliasArgs) error {
	return api.DeleteAlias(c, args)
}

// 获取触发器列表 ListTriggers GET /v1/console/relation
func (c *Client) ListTriggers(args *api.ListTriggersArgs) (*api.ListTriggersResult, error) {
	return api.ListTriggers(c, args)
}

// 创建触发器 CreateTrigger POST /v1/console/relation
func (c *Client) CreateTrigger(args *api.CreateTriggerArgs) (*api.CreateTriggerResult, error) {
	return api.CreateTrigger(c, args)
}

// 更新触发器 UpdateTrigger PUT /v1/console/relation
func (c *Client) UpdateTrigger(args *api.UpdateTriggerArgs) (*api.UpdateTriggerResult, error) {
	return api.UpdateTrigger(c, args)
}

// 删除触发器 DeleteTrigger DELETE /v1/console/relation
func (c *Client) DeleteTrigger(args *api.DeleteTriggerArgs) error {
	return api.DeleteTrigger(c, args)
}
