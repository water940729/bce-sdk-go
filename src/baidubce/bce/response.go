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

func (r *BceResponse) IsFail() bool {
    return r.response.StatusCode() >= 400
}

func (r *BceResponse) StatusCode() int {
    if r.statusCode == 0 {
        r.ParseResponse()
    }
    return r.statusCode
}

func (r *BceResponse) StatusText() string {
    if r.statusText == "" {
        r.ParseResponse()
    }
    return r.statusText
}

func (r *BceResponse) RequestId() string {
    if r.requestId == "" {
        r.ParseResponse()
    }
    return r.requestId
}

func (r *BceResponse) DebugId() string {
    if r.debugId == "" {
        r.ParseResponse()
    }
    return r.debugId
}

func (r *BceResponse) Header(key string) string {
    return r.response.GetHeader(key)
}

func (r *BceResponse) Headers() map[string]string {
    return r.response.GetHeaders()
}

func (r *BceResponse) Body() io.ReadCloser {
    return r.response.Body()
}

func (r *BceResponse) SetHttpResponse(response *http.Response) {
    r.response = response
}

func (r *BceResponse) ElapsedTime() time.Duration {
    return r.response.ElapsedTime()
}

func (r *BceResponse) ServiceError() *BceServiceError {
    return r.serviceError
}

func (r *BceResponse) ParseResponse() {
    r.statusCode = r.response.StatusCode()
    r.statusText = r.response.StatusText()
    r.requestId = r.response.GetHeader(http.BCE_REQUEST_ID)
    r.debugId = r.response.GetHeader(http.BCE_DEBUG_ID)
    if r.IsFail() {
        r.serviceError = &BceServiceError{}
        err := r.ParseJsonBody(r.serviceError)
        if err != nil {
            r.serviceError = NewBceServiceError(
                EMALFORMED_JSON,
                "Service json error message decode failed",
                r.requestId,
                r.statusCode)
        }
        r.serviceError.StatusCode = r.statusCode
    }
}

func (r *BceResponse) ParseJsonBody(result interface{}) error {
    defer r.Body().Close()
    jsonDecoder := json.NewDecoder(r.Body())
    return jsonDecoder.Decode(result)
}

