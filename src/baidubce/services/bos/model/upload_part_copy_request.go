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

// upload_part_copy_request.go - the upload part copy request definition

package model

import (
    "fmt"

    "baidubce/http"
)

type UploadPartCopyRequest struct {
    ObjectRequest
    uploadId        string
    partNumber      int

    sourceBucket    string
    sourceObject    string
    sourceRange     string
    matchETag       string
    noneMatchETag   string
    modifiedSince   string
    unmodifiedSince string
}

func (req *UploadPartCopyRequest) SetUploadId(val string) { req.uploadId = val }
func (req *UploadPartCopyRequest) SetPartNumber(val int) { req.partNumber = val }

func (req *UploadPartCopyRequest) SetSourceBucket(val string) { req.sourceBucket = val }
func (req *UploadPartCopyRequest) SetSourceObject(val string) { req.sourceObject = val }
func (req *UploadPartCopyRequest) SetSourceRange(val string) { req.sourceRange = val }
func (req *UploadPartCopyRequest) SetMatchETag(val string) { req.matchETag = val }
func (req *UploadPartCopyRequest) SetNoneMatchETag(val string) { req.noneMatchETag = val }
func (req *UploadPartCopyRequest) SetModifiedSince(val string) { req.modifiedSince = val }
func (req *UploadPartCopyRequest) SetUnmodifiedSince(val string) { req.unmodifiedSince = val }

func (req *UploadPartCopyRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)
    httpReq.AddParam("uploadId", req.uploadId)
    httpReq.AddParam("partNumber", fmt.Sprintf("%d", req.partNumber))

    httpReq.AddHeader(http.BCE_COPY_SOURCE,
                      fmt.Sprintf("/%s/%s", req.sourceBucket, req.sourceObject))
    if len(req.sourceRange) != 0 {
        httpReq.AddHeader(http.BCE_COPY_SOURCE_RANGE, req.sourceRange)
    }
    if len(req.matchETag) != 0 {
        httpReq.AddHeader(http.BCE_COPY_SOURCE_IF_MATCH, req.matchETag)
    }
    if len(req.noneMatchETag) != 0 {
        httpReq.AddHeader(http.BCE_COPY_SOURCE_IF_NONE_MATCH, req.noneMatchETag)
    }
    if len(req.modifiedSince) != 0 {
        httpReq.AddHeader(http.BCE_COPY_SOURCE_IF_MODIFIED_SINCE, req.modifiedSince)
    }
    if len(req.unmodifiedSince) != 0 {
        httpReq.AddHeader(http.BCE_COPY_SOURCE_IF_UNMODIFIED_SINCE, req.unmodifiedSince)
    }

    return httpReq
}

func NewUploadPartCopyRequest(bucket, object string) *UploadPartCopyRequest {
    return &UploadPartCopyRequest{
        ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}}
}

