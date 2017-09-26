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

// Get object request definition
package model

import (
    "fmt"
    "strings"

    "baidubce/http"
)

var (
    ALLOWED_RESPONSE_HEADERS = []string{
        "ContetDisposition",
        "ContentType",
        "ContentEncoding",
        "CacheControl",
        "Expires"}
)

type GetObjectRequest struct {
    ObjectRequest
    rangeParam      [2]int64
    responseHeaders map[string]string
}

func (req *GetObjectRequest) SetRange(start, end int64) {
    req.rangeParam[0] = start
    req.rangeParam[1] = end
}

func (req *GetObjectRequest) AddResponseHeader(key, val string) error {
    found := false
    for _, allowed := range ALLOWED_RESPONSE_HEADERS {
        if strings.ToLower(key) == strings.ToLower(allowed) {
            found = true
            break
        }
    }
    if !found {
        return fmt.Errorf("%s", "not allowed response header")
    }
    if req.responseHeaders == nil {
        req.responseHeaders = make(map[string]string)
    }
    req.responseHeaders["response" + key] = val
    return nil
}

func (req *GetObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.GET)

    if req.responseHeaders != nil {
        for k, v := range req.responseHeaders {
            httpReq.AddParam(k, v)
        }
    }

    rangeStr := ""
    if req.rangeParam[0] == -1 && req.rangeParam[1] == -1 {
        return httpReq
    } else if req.rangeParam[0] == -1 && req.rangeParam[1] > 0 {
        rangeStr = fmt.Sprintf("bytes=-%d", req.rangeParam[1])
    } else if req.rangeParam[0] >= 0 && req.rangeParam[1] == -1 {
        rangeStr = fmt.Sprintf("bytes=%d-", req.rangeParam[0])
    } else {
        rangeStr = fmt.Sprintf("bytes=%d-%d", req.rangeParam[0], req.rangeParam[1])
    }
    httpReq.AddHeader("Range", rangeStr)

    return httpReq
}

func NewGetObjectRequest(bucket, object string) *GetObjectRequest {
    return &GetObjectRequest{ObjectRequest{bucketName: bucket, objectName: object},
                            [2]int64{-1, -1}, nil}
}

