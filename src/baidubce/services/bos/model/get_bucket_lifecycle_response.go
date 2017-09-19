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

// Get bucket lifecycle response definition
package model

import (
    "baidubce/model"
    "baidubce/util"
)

type GetBucketLifecycleResponse struct {
    model.BceResponse
    rule []LifecycleRuleType
}

func (resp *GetBucketLifecycleResponse) Rule() []LifecycleRuleType { return resp.rule }

func (resp *GetBucketLifecycleResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &BucketLifecycleInOut{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        util.LOGGER.Error().Printf("parse bucket lifecycle json response failed: %v\n", err)
        return
    }
    resp.rule   = jsonBody.Rule
}

func NewGetBucketLifecycleResponse() *GetBucketLifecycleResponse {
    return &GetBucketLifecycleResponse{}
}

