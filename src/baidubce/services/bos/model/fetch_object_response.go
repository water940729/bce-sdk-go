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

// Fetch object response definition
package model

import (
    "baidubce/model"
    "baidubce/util"
)

type FetchObjectResponse struct {
    model.BceResponse
    code      string
    message   string
    requestId string
    jobId     string
}

func (resp *FetchObjectResponse) Code() string { return resp.code }
func (resp *FetchObjectResponse) Message() string { return resp.message }
func (resp *FetchObjectResponse) RequestId() string { return resp.requestId }
func (resp *FetchObjectResponse) JobId() string { return resp.jobId }

func (resp *FetchObjectResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &FetchObjectOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        util.LOGGER.Error().Printf("parse fecth object json response failed: %v\n", err)
        return
    }
    resp.code = jsonBody.Code
    resp.message = jsonBody.Message
    resp.requestId = jsonBody.RequestId
    resp.jobId = jsonBody.JobId
}

func NewFetchObjectResponse() *FetchObjectResponse {
    return &FetchObjectResponse{}
}

