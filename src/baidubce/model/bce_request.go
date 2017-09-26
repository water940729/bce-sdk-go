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

// bce_request.go - defines the common BCE servies request

// Package model implements the BceRequester and BceResponser interfaces to define the request and
// response data structure of BCE services.
package model

import (
    "uuid"

    "baidubce/bce"
    "baidubce/http"
)

// BceRequest defines the concrete request structure for accessing BCE services, and it implements
// the BceRequester interface defined in the bce package.
type BceRequest struct {
    requestId   string
    headers     map[string]string
    body        *http.BodyStream
    clientError *bce.BceClientError
}

func (_ *BceRequest) GetUri() string {
    return "/"
}

func (bceReq *BceRequest) SetHeader(key, val string) {
    if bceReq.headers == nil {
        bceReq.headers = make(map[string]string)
    }
    bceReq.headers[key] = val
}

func (bceReq *BceRequest) SetBody(body *http.BodyStream) {
    bceReq.body = body
}

func (bceReq *BceRequest) ClientError() *bce.BceClientError {
    return bceReq.clientError
}

func (bceReq *BceRequest) SetClientError(err *bce.BceClientError) {
    bceReq.clientError = err
}

func (bceReq *BceRequest) BuildHttpRequest() *http.Request {
    request := &http.Request{}

    // Construct the request ID with UUID V1
    bceReq.requestId = uuid.NewV1().String()

    // Build the http request components
    request.AddHeader(http.BCE_REQUEST_ID, bceReq.requestId)
    request.SetUri(bceReq.GetUri())
    if bceReq.body != nil {
        request.SetBody(bceReq.body)
    }
    for k, v := range bceReq.headers {
        request.AddHeader(k, v)
    }
    return request
}

