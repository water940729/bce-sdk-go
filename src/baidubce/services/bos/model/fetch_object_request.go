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

// fetch_object_request.go - the fetch object request definition

package model

import (
    "glog"

    "baidubce/http"
)

type FetchObjectRequest struct {
    ObjectRequest
    fetchSource  string
    fetchMode    string
    storageClass string
}

func (req *FetchObjectRequest) SetFetchSource(val string) {
    req.fetchSource = val
}

func (req *FetchObjectRequest) SetFetchMode(val string) {
    if val != FETCH_MODE_SYNC && val != FETCH_MODE_ASYNC {
        glog.Info(
            "invalid fetch mode value: %s, use default:%s\n", val, FETCH_MODE_SYNC)
        val = FETCH_MODE_SYNC
    }
    req.fetchMode = val
}

func (req *FetchObjectRequest) SetStorageClass(val string) {
    if val != STORAGE_CLASS_STANDARD && val != STORAGE_CLASS_STANDARD_IA &&
            val != STORAGE_CLASS_COLD {
        glog.Info(
            "invalid storage class value: %s, use default: %s\n", val, STORAGE_CLASS_STANDARD)
        val = STORAGE_CLASS_STANDARD
    }
    req.storageClass = val
}

func (req *FetchObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("fetch", "")

    if len(req.storageClass) == 0 {
        req.storageClass = STORAGE_CLASS_STANDARD
    }
    if len(req.fetchMode) == 0 {
        req.storageClass = FETCH_MODE_SYNC
    }
    if len(req.fetchSource) == 0 {
        glog.Error("fetch source not set")
    }
    httpReq.AddHeader(http.BCE_STORAGE_CLASS, req.storageClass)
    httpReq.AddHeader(http.BCE_PREFIX + "fetch-mode", req.fetchMode)
    httpReq.AddHeader(http.BCE_PREFIX + "fetch-source", req.fetchSource)

    return httpReq
}

func NewFetchObjectRequest(bucket, object string) *FetchObjectRequest {
    return &FetchObjectRequest{
        ObjectRequest: ObjectRequest{bucketName: bucket, objectName: object}}
}

