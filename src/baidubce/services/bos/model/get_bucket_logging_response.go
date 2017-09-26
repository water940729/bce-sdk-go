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

// get_bucket_logging_response.go - the get bucket logging response definition

package model

import (
    "glog"

    "baidubce/model"
)

type GetBucketLoggingResponse struct {
    model.BceResponse
    status       string
    targetBucket string
    targetPrefix string
}

func (resp *GetBucketLoggingResponse) Status() string { return resp.status }

func (resp *GetBucketLoggingResponse) TargetBucket() string { return resp.targetBucket }

func (resp *GetBucketLoggingResponse) TargetPrefix() string { return resp.targetPrefix }

func (resp *GetBucketLoggingResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &BucketLoggingOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse logging json response failed: %v\n", err)
        return
    }
    resp.status = jsonBody.Status
    resp.targetBucket = jsonBody.TargetBucket
    resp.targetPrefix = jsonBody.TargetPrefix
}

func NewGetBucketLoggingResponse() *GetBucketLoggingResponse {
    return &GetBucketLoggingResponse{}
}

