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

// list_multipart_uploads_request.go - the list multipart uploads request definition

package model

import (
    "strconv"

    "baidubce/http"
)

type ListMultipartUploadsRequest struct {
    BucketRequest
    delimiter  string
    keyMarker  string
    maxUploads int
    prefix     string
}

func (req *ListMultipartUploadsRequest) SetDelimiter(v string) {
    req.delimiter = v
}

func (req *ListMultipartUploadsRequest) SetKeyMarker(v string) {
    req.keyMarker = v
}

func (req *ListMultipartUploadsRequest) SetMaxUploads(v int) {
    if v > 1000 || v < 0 {
        v = 1000
    }
    req.maxUploads = v
}

func (req *ListMultipartUploadsRequest) SetPrefix(v string) {
    req.prefix = v
}

func (req *ListMultipartUploadsRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.GET)
    httpReq.AddParam("uploads", "")

    if len(req.delimiter) != 0 {
        httpReq.AddParam("delimiter", req.delimiter)
    }
    if len(req.keyMarker) != 0 {
        httpReq.AddParam("keyMarker", req.keyMarker)
    }
    if req.maxUploads >= 0 && req.maxUploads <= 1000 {
        httpReq.AddParam("maxUploads", strconv.Itoa(req.maxUploads))
    }
    if len(req.prefix) != 0 {
        httpReq.AddParam("prefix", req.prefix)
    }

    return httpReq
}

func NewListMultipartUploadsRequest(name string) *ListMultipartUploadsRequest {
    return &ListMultipartUploadsRequest{
            BucketRequest: BucketRequest{bucketName: name}}
}

