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

// list_multipart_uploads_response.go - the list multipart uploads response definition

package model

import (
    "glog"

    "baidubce/model"
)

type ListMultipartUploadsResponse struct {
    model.BceResponse
    bucket         string
    commonPrefixes string
    delimiter      string
    prefix         string
    isTruncated    bool
    keyMarker      string
    maxUploads     int
    nextKeyMarker  string
    uploads        []ListMultipartUploadsType
}

func (resp *ListMultipartUploadsResponse) Bucket() string { return resp.bucket }
func (resp *ListMultipartUploadsResponse) CommonPrefixes() string { return resp.commonPrefixes }
func (resp *ListMultipartUploadsResponse) Prefix() string { return resp.prefix }
func (resp *ListMultipartUploadsResponse) Delimiter() string { return resp.delimiter }
func (resp *ListMultipartUploadsResponse) IsTruncated() bool { return resp.isTruncated }
func (resp *ListMultipartUploadsResponse) KeyMarker() string { return resp.keyMarker }
func (resp *ListMultipartUploadsResponse) MaxUploads() int { return resp.maxUploads }
func (resp *ListMultipartUploadsResponse) NextKeyMarker() string { return resp.nextKeyMarker }
func (resp *ListMultipartUploadsResponse) Uploads() []ListMultipartUploadsType {return resp.uploads}

func (resp *ListMultipartUploadsResponse) ParseResponse() {
    resp.BceResponse.ParseResponse()
    if resp.BceResponse.IsFail() { return }

    jsonBody := &ListMultipartUploadsOutput{}
    if err := resp.BceResponse.ParseJsonBody(jsonBody); err != nil {
        glog.Error("parse list multipart uploads json response failed: %v\n", err)
        return
    }

    resp.bucket = jsonBody.Bucket
    resp.commonPrefixes = jsonBody.CommonPrefixes
    resp.prefix = jsonBody.Prefix
    resp.delimiter = jsonBody.Delimiter
    resp.isTruncated = jsonBody.IsTruncated
    resp.keyMarker = jsonBody.KeyMarker
    resp.maxUploads = jsonBody.MaxUploads
    resp.nextKeyMarker = jsonBody.NextKeyMarker
    resp.uploads = jsonBody.Uploads
}

func NewListMultipartUploadsResponse() *ListMultipartUploadsResponse {
    return &ListMultipartUploadsResponse{}
}

