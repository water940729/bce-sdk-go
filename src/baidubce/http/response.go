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

// response.go - the custom HTTP response for BCE

package http

import (
    "io"
    "net/http"
    "time"
)

// Response defines the general http response structure for accessing the BCE services.
type Response struct {
    httpResponse *http.Response // the standard libaray http.Response object
    elapsedTime  time.Duration  // elapsed time just from sending request to receiving response
}

func (resp *Response) StatusText() string {
    return resp.httpResponse.Status
}

func (resp *Response) StatusCode() int {
    return resp.httpResponse.StatusCode
}

func (resp *Response) Protocol() string {
    return resp.httpResponse.Proto
}

func (resp *Response) HttpResponse() *http.Response {
    return resp.httpResponse
}

func (resp *Response) SetHttpResponse(response *http.Response) {
    resp.httpResponse = response
}

func (resp *Response) ElapsedTime() time.Duration {
    return resp.elapsedTime
}

func (resp *Response) GetHeader(name string) string {
    return resp.httpResponse.Header.Get(name)
}

func (resp *Response) GetHeaders() map[string]string {
    header := resp.httpResponse.Header
    ret := make(map[string]string, len(header))
    for k, v := range header {
        ret[k] = v[0]
    }
    return ret
}

func (resp *Response) ContentLength() int64 {
    return resp.httpResponse.ContentLength
}

func (resp *Response) Body() io.ReadCloser {
    return resp.httpResponse.Body
}

