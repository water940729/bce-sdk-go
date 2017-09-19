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

// Initiate multipart upload response definition
package model

import (
    "baidubce/model"
    "baidubce/util"
)

type InitiateMultipartUploadResponse struct {
    model.BceResponse
    bucket   string
    key      string
    uploadId string
}

func (resp *InitiateMultipartUploadResponse) Bucket() string { return resp.bucket }
func (resp *InitiateMultipartUploadResponse) Key() string { return resp.key }
func (resp *InitiateMultipartUploadResponse) UploadId() string { return resp.uploadId }

func (resp *InitiateMultipartUploadResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &InitiateMultipartUploadOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        util.LOGGER.Error().Printf("parse initiate multipart upload response failed: %v\n", err)
        return
    }
    resp.bucket = jsonBody.Bucket
    resp.key = jsonBody.Key
    resp.uploadId = jsonBody.UploadId
}

func NewInitiateMultipartUploadResponse() *InitiateMultipartUploadResponse {
    return &InitiateMultipartUploadResponse{}
}

