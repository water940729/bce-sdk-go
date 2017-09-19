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

// List objects in the given bucket request definition
package model

import (
    "strconv"

    "baidubce/http"
)

type ListObjectsRequest struct {
    BucketRequest
    delimiter string
    marker    string
    maxKeys   int
    prefix    string
}

func (req *ListObjectsRequest) SetDelimiter(v string) {
    req.delimiter = v
}

func (req *ListObjectsRequest) SetMarker(v string) {
    req.marker = v
}

func (req *ListObjectsRequest) SetMaxKeys(v int) {
    if v > 1000 || v < 0 {
        v = 1000
    }
    req.maxKeys = v
}

func (req *ListObjectsRequest) SetPrefix(v string) {
    req.prefix = v
}

func (req *ListObjectsRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.GET)

    if len(req.delimiter) != 0 {
        httpReq.AddParam("delimiter", req.delimiter)
    }

    if len(req.marker) != 0 {
        httpReq.AddParam("marker", req.marker)
    }

    httpReq.AddParam("maxKeys", strconv.Itoa(req.maxKeys))

    if len(req.prefix) != 0 {
        httpReq.AddParam("prefix", req.prefix)
    }

    return httpReq
}

func NewListObjectsRequest(name string) *ListObjectsRequest {
    return &ListObjectsRequest{BucketRequest{bucketName: name}, "", "", 1000, ""}
}

