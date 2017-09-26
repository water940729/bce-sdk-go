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

// complete_multipart_upload_response.go - the complete multipart upload response definition

package model

import (
    "glog"

    "baidubce/model"
)

type CompleteMultipartUploadResponse struct {
    model.BceResponse
    bucket   string
    key      string
    eTag     string
    location string
}

func (resp *CompleteMultipartUploadResponse) Bucket() string { return resp.bucket }
func (resp *CompleteMultipartUploadResponse) Key() string { return resp.key }
func (resp *CompleteMultipartUploadResponse) ETag() string { return resp.eTag }
func (resp *CompleteMultipartUploadResponse) Location() string { return resp.location }

func (resp *CompleteMultipartUploadResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &CompleteMultipartUploadOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse complete multipart upload response failed: %v\n", err)
        return
    }
    resp.bucket = jsonBody.Bucket
    resp.key = jsonBody.Key
    resp.eTag = jsonBody.ETag
    resp.location = jsonBody.Location
}

func NewCompleteMultipartUploadResponse() *CompleteMultipartUploadResponse {
    return &CompleteMultipartUploadResponse{}
}

