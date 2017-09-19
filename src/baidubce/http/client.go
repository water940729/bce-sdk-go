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

// Definine HTTP client facilities for BCE
package http

import (
    "time"
    "net/http"
    "net/url"
)

func Execute(request *Request) (*Response, error) {
    // Prepare the internal data structure for sending the http request
    internalClient := &http.Client{ Timeout: time.Duration(request.Timeout()) * time.Second }

    internalUrl := &url.URL{
        Scheme: request.Protocol(),
        Host: request.Host(),
        Path: request.Uri(),
        RawQuery: request.QueryString()}

    internalHeader := make(http.Header)
    for k, v := range request.Headers() {
        val := make([]string, 0, 1)
        val = append(val, v)
        internalHeader[k] = val
    }

    body := request.Body()
    internalReq := &http.Request{
        Method: request.Method(),
        URL: internalUrl,
        Header: internalHeader,
        Body: body}
    // Variable body's type is `*BodyStream`. If its value is nil, the `Body` field must be
    // explicitly assigned `nil` value, otherwise nil pointer dereference will arise.
    if body == nil {
        internalReq.Body = nil
        internalReq.ContentLength = 0
    } else {
        internalReq.ContentLength = body.Len()
    }

    // Perform the http request and get response
    start := time.Now()
    internalResp, err := internalClient.Do(internalReq)
    end := time.Now()
    if err != nil {
        return nil, err
    }
    response := &Response{internalResp, end.Sub(start)}
    return response, nil
}

