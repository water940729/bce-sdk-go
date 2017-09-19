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

// Initiate multipart upload request definition
package model

import (
    "baidubce/http"
    "baidubce/util"
)

type InitiateMultipartUploadRequest struct {
    ObjectRequest
    storageClass string

    // Request params fields
    cacheControl       string
    contentDisposition string
    contentType        string
    expires            string
    uploads            string
}

func (req *InitiateMultipartUploadRequest) SetStorageClass(val string) {
    if val != STORAGE_CLASS_STANDARD && val != STORAGE_CLASS_STANDARD_IA &&
            val != STORAGE_CLASS_COLD {
        util.LOGGER.Info().Printf(
            "invalid storage class value: %s, use default: %s\n", val, STORAGE_CLASS_STANDARD)
        val = STORAGE_CLASS_STANDARD
    }
    req.storageClass = val
}

func (req *InitiateMultipartUploadRequest) SetCacheControl(val string) { req.cacheControl = val }
func (req *InitiateMultipartUploadRequest) SetContentDisposition(val string) {
    req.contentDisposition = val
}
func (req *InitiateMultipartUploadRequest) SetContentType(val string) { req.contentType = val }
func (req *InitiateMultipartUploadRequest) SetExpires(val string) { req.expires = val }
func (req *InitiateMultipartUploadRequest) SetUploads(val string) { req.uploads = val }

func (req *InitiateMultipartUploadRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("uploads", "")

    if req.storageClass == STORAGE_CLASS_STANDARD_IA || req.storageClass == STORAGE_CLASS_COLD {
        httpReq.AddHeader(http.BCE_STORAGE_CLASS, req.storageClass)
    }

    if len(req.cacheControl) != 0 {
        httpReq.AddParam(http.CACHE_CONTROL, req.cacheControl)
    }
    if len(req.contentDisposition) != 0 {
        httpReq.AddParam(http.CONTENT_DISPOSITION, req.contentDisposition)
    }
    if len(req.contentType) != 0 {
        httpReq.AddParam(http.CONTENT_TYPE, req.contentType)
    }
    if len(req.expires) != 0 {
        httpReq.AddParam(http.EXPIRES, req.expires)
    }
    if len(req.uploads) != 0 {
        httpReq.AddParam("uploads", req.uploads)
    }

    return httpReq
}

func NewInitiateMultipartUploadRequest(bucket, object string) *InitiateMultipartUploadRequest {
    return &InitiateMultipartUploadRequest{
                ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}}
}

