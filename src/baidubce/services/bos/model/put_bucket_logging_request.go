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

// put_bucket_logging_request.go - the put bucket logging request definition

package model

import (
    "encoding/json"

    "baidubce/bce"
    "baidubce/http"
)

type PutBucketLoggingRequest struct {
    BucketRequest
    logging interface{} // logging string | logging stream
}

func (req *PutBucketLoggingRequest) SetLogging(val interface{}) {
    req.logging = val
}

func (req *PutBucketLoggingRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)
    httpReq.AddParam("logging", "")
    httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")

    var body *http.BodyStream
    switch val := req.logging.(type) {
    case string:
        body = http.NewBodyStreamFromString(val)
    case *http.BodyStream:
        jsonBytes, err := json.Marshal(val)
        if err != nil {
            req.SetClientError(bce.NewBceClientError("loggin json encode fail: " + err.Error()))
            return nil
        }
        body = http.NewBodyStreamFromBytes(jsonBytes)
    default:
        req.SetClientError(bce.NewBceClientError("invalid lifecycle parameter"))
        return nil
    }
    httpReq.SetBody(body)
    return httpReq
}

func NewPutBucketLoggingRequest(name string, logging interface{}) *PutBucketLoggingRequest {
    return &PutBucketLoggingRequest{BucketRequest{bucketName: name}, logging}
}

