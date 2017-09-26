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

// get_session_token.go - define the response of GetSessionToken for STS service

package model

import (
    "glog"

    "baidubce/model"
)

type GetSessionTokenJsonSchema struct {
    AccessKeyId     string
    SecretAccessKey string
    SessionToken    string
    CreateTime      string
    Expiration      string
    UserId          string
}

type GetSessionTokenResponse struct {
    model.BceResponse
    accessKeyId     string
    secretAccessKey string
    sessionToken    string
    createTime      string
    expiration      string
    userId          string
}

// Readonly response fields
func (resp *GetSessionTokenResponse) AccessKeyId() string { return resp.accessKeyId }
func (resp *GetSessionTokenResponse) SecretAccessKey() string { return resp.secretAccessKey }
func (resp *GetSessionTokenResponse) SessionToken() string { return resp.sessionToken }
func (resp *GetSessionTokenResponse) CreateTime() string { return resp.createTime }
func (resp *GetSessionTokenResponse) Expiration() string { return resp.expiration }
func (resp *GetSessionTokenResponse) UserId() string { return resp.userId }

func (resp *GetSessionTokenResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &GetSessionTokenJsonSchema{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse GetSessionToken json response failed: %v\n", err)
        return
    }
    resp.accessKeyId     = jsonBody.AccessKeyId
    resp.secretAccessKey = jsonBody.SecretAccessKey
    resp.sessionToken    = jsonBody.SessionToken
    resp.createTime      = jsonBody.CreateTime
    resp.expiration      = jsonBody.Expiration
    resp.userId          = jsonBody.UserId
}

func NewGetSessionTokenResponse() *GetSessionTokenResponse {
    return &GetSessionTokenResponse{}
}

