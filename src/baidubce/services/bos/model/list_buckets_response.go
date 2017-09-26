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

// list_buckets_response.go - the list buckets response definition

package model

import (
    "glog"

    "baidubce/model"
)

type ListBucketsResponse struct {
    model.BceResponse
    owner   BucketOwnerType
    buckets []BucketSummary
}

func (resp *ListBucketsResponse) Owner() BucketOwnerType { return resp.owner }
func (resp *ListBucketsResponse) Buckets() []BucketSummary { return resp.buckets }

func (resp *ListBucketsResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &ListBucketsOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse list buckets json response failed: %v\n", err)
        return
    }
    resp.owner = jsonBody.Owner
    resp.buckets = jsonBody.Buckets
}

func NewListBucketsResponse() *ListBucketsResponse {
    return &ListBucketsResponse{}
}

