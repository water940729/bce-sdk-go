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

// response.go - defines the common BCE services response

package bce

import (
    "io"
    "time"
    "encoding/json"

    "baidubce/http"
)

// BceResponse defines the response structure for receiving BCE services response.
type BceResponse struct {
    statusCode   int
    statusText   string
    requestId    string
    debugId      string
    response     *http.Response
    serviceError *BceServiceError
}

func (resp *BceResponse) IsFail() bool {
    return resp.response.StatusCode() >= 400
}

func (resp *BceResponse) StatusCode() int {
    if resp.statusCode == 0 {
        resp.ParseResponse()
    }
    return resp.statusCode
}

func (resp *BceResponse) StatusText() string {
    if resp.statusText == "" {
        resp.ParseResponse()
    }
    return resp.statusText
}

func (resp *BceResponse) RequestId() string {
    if resp.requestId == "" {
        resp.ParseResponse()
    }
    return resp.requestId
}

func (resp *BceResponse) DebugId() string {
    if resp.debugId == "" {
        resp.ParseResponse()
    }
    return resp.debugId
}

func (resp *BceResponse) Header(key string) string {
    return resp.response.GetHeader(key)
}

func (resp *BceResponse) Headers() map[string]string {
    return resp.response.GetHeaders()
}

func (resp *BceResponse) Body() io.ReadCloser {
    return resp.response.Body()
}

func (resp *BceResponse) SetHttpResponse(response *http.Response) {
    resp.response = response
}

func (resp *BceResponse) ElapsedTime() time.Duration {
    return resp.response.ElapsedTime()
}

func (resp *BceResponse) ServiceError() *BceServiceError {
    return resp.serviceError
}

func (resp *BceResponse) ParseResponse() {
    resp.statusCode = resp.response.StatusCode()
    resp.statusText = resp.response.StatusText()
    resp.requestId = resp.response.GetHeader(http.BCE_REQUEST_ID)
    resp.debugId = resp.response.GetHeader(http.BCE_DEBUG_ID)
    if resp.IsFail() {
        resp.serviceError = &BceServiceError{}
        err := resp.ParseJsonBody(resp.serviceError)
        if err != nil {
            resp.serviceError = NewBceServiceError(
                EMALFORMED_JSON,
                "Service json error message decode failed",
                resp.requestId,
                resp.statusCode)
        }
        resp.serviceError.StatusCode = resp.statusCode
    }
}

func (resp *BceResponse) ParseJsonBody(result interface{}) error {
    defer resp.Body().Close()
    jsonDecoder := json.NewDecoder(resp.Body())
    return jsonDecoder.Decode(result)
}

