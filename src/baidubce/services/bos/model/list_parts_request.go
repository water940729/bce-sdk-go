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

// list_parts_request.go - the list the uploaded parts of a multipart upload

package model

import (
    "strconv"

    "baidubce/http"
)

type ListPartsRequest struct {
    ObjectRequest
    maxParts         int
    partNumberMarker string
    uploadId         string
}

func (req *ListPartsRequest) SetMaxParts(v int) {
    if v > 1000 || v < 0 {
        v = 1000
    }
    req.maxParts = v
}

func (req *ListPartsRequest) SetPartNumberMarker(v string) {
    req.partNumberMarker = v
}

func (req *ListPartsRequest) SetUploadId(v string) {
    req.uploadId = v
}

func (req *ListPartsRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.GET)
    httpReq.AddParam("maxParts", strconv.Itoa(req.maxParts))
    httpReq.AddParam("partNumberMarker", req.partNumberMarker)
    httpReq.AddParam("uploadId", req.uploadId)

    return httpReq
}

func NewListPartsRequest(bucket, obj string) *ListPartsRequest {
    return &ListPartsRequest{ObjectRequest{bucketName: bucket, objectName: obj}, 1000, "", ""}
}

