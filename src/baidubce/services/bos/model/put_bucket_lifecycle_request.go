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

// Put bucket lifecycle request definition
package model

import (
    "encoding/json"

    "baidubce/http"
    "baidubce/common"
)

type PutBucketLifecycleRequest struct {
    BucketRequest
    lifecycle interface{}  // `acl string` | acl file body stream
}

func (req *PutBucketLifecycleRequest) SetLifecycle(val interface{}) {
    req.lifecycle = val
}

func (req *PutBucketLifecycleRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)
    httpReq.AddParam("lifecycle", "")
    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")

    var body *http.BodyStream
    switch val := req.lifecycle.(type) {
    case string:
        body = http.NewBodyStreamFromString(val)
    case *http.BodyStream:
        body = val
    case *BucketLifecycleInOut:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(common.NewBceClientError("lifecycle json encode fail: " + err.Error()))
            return nil
        }
        body = http.NewBodyStreamFromBytes(jsonBytes)
    default:
        req.SetClientError(common.NewBceClientError("invalid lifecycle parameter"))
        return nil
    }
    httpReq.SetBody(body)
    return httpReq
}

func NewPutBucketLifecycleRequest(name string, lifecycle interface{}) *PutBucketLifecycleRequest {
    return &PutBucketLifecycleRequest{BucketRequest{bucketName: name}, lifecycle}
}

