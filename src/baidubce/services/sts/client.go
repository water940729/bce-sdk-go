/*
 * Copyright 2014 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Define client for STS service
package sts

import (
    "baidubce/auth"
    "baidubce/bce"
    "baidubce/services/sts/model"
    "baidubce/util"
)

type Client struct {
    // Client of STS service is a kind of BceClient, so derived from
    // the BceClient and no other fields needed.
    *bce.BceClient
}

func (cli *Client) GetSessionToken(req *model.GetSessionTokenRequest,
        resp *model.GetSessionTokenResponse) error {
    return cli.SendRequest(req, resp)
}

// Make the STS service client with default configuration
// Use `cli.Config.xxx` to access the config or change it to non-default value
func NewClient(ak, sk, endpoint string) (*Client, error) {
    credentials, err := auth.NewBceDefaultCredentials(ak, sk)
    if err != nil {
        return nil, err
    }
    defaultSignOptions := &auth.SignOptions{
        auth.DEFAULT_HEADERS_TO_SIGN,
        util.NowUTCSeconds(),
        auth.DEFAULT_EXPIRE_SECONDS}
    defaultConf := &bce.BceClientConfiguration{
        bce.DEFAULT_PROTOCOL,
        endpoint,
        bce.DEFAULT_REGION,
        credentials,
        defaultSignOptions,
        bce.DEFAULT_RETRY_POLICY,
        bce.DEFAULT_USER_AGENT,
        bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
    v1Signer := &auth.BceV1Signer{}

    client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
    return client, nil
}

