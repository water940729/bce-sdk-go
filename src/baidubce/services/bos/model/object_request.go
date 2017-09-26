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

// object_requst.go - define common object request structure for reuse

package model

import (
    "baidubce/bce"
    "baidubce/http"
    "baidubce/model"
)

type ObjectMetadata struct {
    ContentLength      int64
    AppendOffset       int64
    InstanceLength     int64
    LastModified       string
    ContentMD5         string
    ContentSha256      string
    ContentRange       string
    ContentType        string
    ContentEncoding    string
    ContentDisposition string
    ETag               string
    Expires            string
    CacheControl       string
    StorageClass       string
    ObjectType         string
    UserMetadata       map[string]string
}

func (obj *ObjectMetadata) SetUserMeta(key, val string) {
    if obj.UserMetadata == nil {
        obj.UserMetadata = make(map[string]string)
    }
    obj.UserMetadata[key] = val
}

type ObjectRequest struct {
    model.BceRequest // Derived from BceRequest
    bucketName string
    objectName string
    *ObjectMetadata
}

func (req *ObjectRequest) GetUri() string {
    return bce.URI_PREFIX  + req.bucketName + "/" + req.objectName
}

func (req *ObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BceRequest.BuildHttpRequest()
    httpReq.SetUri(req.GetUri())
    return httpReq
}

