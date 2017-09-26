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

// get_bucket_storageclass_response.go - the get bucket storage class response definition

package model

import (
    "glog"

    "baidubce/model"
)

type GetBucketStorageClassResponse struct {
    model.BceResponse
    storageClass string
}

func (resp *GetBucketStorageClassResponse) StorageClass() string { return resp.storageClass }

func (resp *GetBucketStorageClassResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &BucketStorageClassInOut{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse storage class json response failed: %v\n", err)
        return
    }
    resp.storageClass = jsonBody.StorageClass
}

func NewGetBucketStorageClassResponse() *GetBucketStorageClassResponse {
    return &GetBucketStorageClassResponse{}
}

