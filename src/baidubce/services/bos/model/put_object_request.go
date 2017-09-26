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

// put_object_request.go - the put object request definition

package model

import (
    "baidubce/bce"
    "baidubce/http"
)

type PutObjectRequest struct {
    ObjectRequest
    content interface{} // Input content: raw string or *BodyStream
}

func (req *PutObjectRequest) SetContent(val interface{}) {
    req.content = val
}

func (req *PutObjectRequest) BuildHttpRequest() *http.Request {
    httpReq := req.ObjectRequest.BuildHttpRequest()
    httpReq.SetMethod(http.PUT)

    switch val := req.content.(type) {
    case string:
        httpReq.SetBody(http.NewBodyStreamFromString(val))
    case *http.BodyStream:
        httpReq.SetBody(val)
    default:
        req.SetClientError(bce.NewBceClientError("invalid put object input type"))
        return nil
    }
    return httpReq
}

func NewPutObjectRequest(bucket, object string, input interface{}) *PutObjectRequest {
    return &PutObjectRequest{ObjectRequest{bucketName: bucket, objectName: object}, input}
}

