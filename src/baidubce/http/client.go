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

// client.go - define the execute function to send http request and get response

// Package http defines the structure of request and response which used to access the BCE services
// as well as the http constant headers and and methods. And finally implement the `Execute` funct-
// ion to do the work.
package http

import (
    "time"
    "net/http"
    "net/url"
)

// The httpClient and httpRequest are the global variable to send the request and get response
// for reuse.
var (
    httpClient = &http.Client{}
    httpRequest = &http.Request{}
)

func Execute(request *Request) (*Response, error) {
    // Set the connection timeout for current request
    httpClient.Timeout = time.Duration(request.Timeout()) * time.Second

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

    // Setup the proxy setting if needed
    defaultTr := http.DefaultTransport
    tr, _ := defaultTr.(*http.Transport)
    if len(request.ProxyUrl()) != 0 {
        tr.Proxy = func(_ *http.Request) (*url.URL, error) {
            return url.Parse(request.ProxyUrl())
        }
    }
    httpClient.Transport = tr

    // Perform the http request and get response
    start := time.Now()
    internalResp, err := httpClient.Do(internalReq)
    end := time.Now()
    if err != nil {
        return nil, err
    }
    response := &Response{internalResp, end.Sub(start)}
    return response, nil
}

