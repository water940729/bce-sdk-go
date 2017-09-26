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

// copy_object_requset.go - the copy object request definition

package model

import (
    "fmt"

    "glog"

    "baidubce/http"
)

type CopyObjectRequest struct {
    ObjectRequest
    sourceBucket      string
    sourceObject      string
    matchETag         string
    noneMatchETag     string
    modifiedSince     string
    unmodifiedSince   string
    metadataDirective string
    storageClass      string
}

func (req *CopyObjectRequest) SetSourceBucket(val string) { req.sourceBucket = val }
func (req *CopyObjectRequest) SetSourceObject(val string) { req.sourceObject = val }
func (req *CopyObjectRequest) SetMatchETag(val string) { req.matchETag = val }
func (req *CopyObjectRequest) SetNoneMatchETag(val string) { req.noneMatchETag = val }
func (req *CopyObjectRequest) SetModifiedSince(val string) { req.modifiedSince = val }
func (req *CopyObjectRequest) SetUnmodifiedSince(val string) { req.unmodifiedSince = val }
func (req *CopyObjectRequest) SetMetadataDirective(val string) {
    if val != METADATA_DIRECTIVE_COPY && val != METADATA_DIRECTIVE_REPLACE {
        glog.Info(
            "invalid metadata directive value: %s, use default: %s\n", val, METADATA_DIRECTIVE_COPY)
        val = METADATA_DIRECTIVE_COPY
    }
    req.metadataDirective = val
}
func (req *CopyObjectRequest) SetStorageClass(val string) {
    if val != STORAGE_CLASS_STANDARD && val != STORAGE_CLASS_STANDARD_IA &&
            val != STORAGE_CLASS_COLD {
        glog.Info(
            "invalid storage class value: %s, use default: %s\n", val, STORAGE_CLASS_STANDARD)
        val = STORAGE_CLASS_STANDARD
    }
    req.storageClass = val
}

func (req *CopyObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)

    httpReq.AddHeader(http.BCE_COPY_SOURCE,
                      fmt.Sprintf("/%s/%s", req.sourceBucket, req.sourceObject))
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
    if req.metadataDirective == METADATA_DIRECTIVE_COPY ||
            req.metadataDirective == METADATA_DIRECTIVE_REPLACE {
        httpReq.AddHeader(http.BCE_COPY_METADATA_DIRECTIVE, req.metadataDirective)
    }
    if req.storageClass == STORAGE_CLASS_STANDARD_IA ||
            req.storageClass == STORAGE_CLASS_COLD {
        httpReq.AddHeader(http.BCE_STORAGE_CLASS, req.storageClass)
    }

    return httpReq
}

func NewCopyObjectRequest(bucket, object string) *CopyObjectRequest {
    return &CopyObjectRequest{ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}}
}

