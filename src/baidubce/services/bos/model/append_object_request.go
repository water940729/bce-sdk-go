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

// append_object_request.go - the append object request definition

package model

import (
    "fmt"

    "glog"

    "baidubce/bce"
    "baidubce/http"
)

type AppendObjectRequest struct {
    ObjectRequest
    offset  int64
    content interface{} // Input content: raw string or *BodyStream

    cacheControl       string
    contentDisposition string
    contentMD5         string
    expires            string
    userMeta           map[string]string
    contentSha256      string
    storageClass       string
}

func (req *AppendObjectRequest) SetOffset(val int64) {
    if val < 0 {
        glog.Error("append object offset should be non-negative number:", val)
    }
    req.offset = val
}

func (req *AppendObjectRequest) SetCacheControl(val string) { req.cacheControl = val }

func (req *AppendObjectRequest) SetContentDisposition(val string) { req.contentDisposition = val }

func (req *AppendObjectRequest) SetContentMD5(val string) { req.contentMD5 = val }

func (req *AppendObjectRequest) SetExpires(val string) { req.expires = val }

func (req *AppendObjectRequest) SetUserMeta(key, val string) {
    if req.userMeta == nil {
        req.userMeta = make(map[string]string)
    }
    req.userMeta[key] = val
}

func (req *AppendObjectRequest) SetContentSha256(val string) { req.contentSha256 = val }

func (req *AppendObjectRequest) SetStorageClass(val string) {
    if val != STORAGE_CLASS_STANDARD && val != STORAGE_CLASS_STANDARD_IA &&
            val != STORAGE_CLASS_COLD {
        glog.Error(
            "invalid storage class value: %s, use default: %s\n", val, STORAGE_CLASS_STANDARD)
        val = STORAGE_CLASS_STANDARD
    }
    req.storageClass = val
}

func (req *AppendObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("append", "")

    if req.offset != 0 {
        httpReq.AddParam("offset", fmt.Sprintf("%d", req.offset))
    }

    if len(req.cacheControl) != 0 {
        httpReq.AddHeader(http.CACHE_CONTROL, req.cacheControl)
    }

    if len(req.contentDisposition) != 0 {
        httpReq.AddHeader(http.CONTENT_DISPOSITION, req.contentDisposition)
    }

    if len(req.contentMD5) != 0 {
        httpReq.AddHeader(http.CONTENT_MD5, req.contentMD5)
    }

    if len(req.expires) != 0 {
        httpReq.AddHeader(http.EXPIRES, req.expires)
    }

    if req.userMeta != nil {
        for k, v := range req.userMeta {
            httpReq.AddHeader(k, v)
        }
    }

    if len(req.contentSha256) != 0 {
        httpReq.AddHeader(http.BCE_CONTENT_SHA256, req.contentSha256)
    }

    if len(req.storageClass) != 0 {
        httpReq.AddHeader(http.BCE_STORAGE_CLASS, req.storageClass)
    }

    switch val := req.content.(type) {
    case string:
        httpReq.SetBody(http.NewBodyStreamFromString(val))
    case *http.BodyStream:
        httpReq.SetBody(val)
    default:
        req.SetClientError(bce.NewBceClientError("invalid append object input type"))
        return nil
    }

    return httpReq
}

func NewAppendObjectRequest(bucket, object string, input interface{}) *AppendObjectRequest {
    return &AppendObjectRequest{
            ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object},
            content: input}
}

