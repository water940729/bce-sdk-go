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

// Define the common BCE request
package model

import (
    "baidubce/common"
    "baidubce/http"
    "baidubce/util"
)

type BceRequester interface {
    GetUri() string
    SetHeader(key, val string)
    SetBody(*http.BodyStream)
    BuildHttpRequest() *http.Request  // May be override by specific request
    ClientError() *common.BceClientError
}

// The common BceRequest structure
type BceRequest struct {
    requestId   string
    headers     map[string]string
    body        *http.BodyStream
    clientError *common.BceClientError
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

func (bceReq *BceRequest) ClientError() *common.BceClientError {
    return bceReq.clientError
}

func (bceReq *BceRequest) SetClientError(err *common.BceClientError) {
    bceReq.clientError = err
}

func (bceReq *BceRequest) BuildHttpRequest() *http.Request {
    request := &http.Request{}

    // Construct the request ID with UUID V1
    bceReq.requestId = util.NewV1().String()

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

