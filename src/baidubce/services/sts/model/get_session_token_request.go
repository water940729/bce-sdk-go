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

// get_session_token.go - define the request of GetSessionToken for STS service

// Package model defines the concrete request and response of the STS service.
package model

import (
    "strconv"

    "glog"

    "baidubce/bce"
    "baidubce/model"
    "baidubce/http"
)

const (
    DEFAULT_DURATION_SECONDS = 43200 // default is 12 hours
    REQUEST_URI = "sessionToken"
)

type GetSessionTokenRequest struct {
    model.BceRequest
    durationSeconds int
    acl             string
}

func (req *GetSessionTokenRequest) SetDurationSeconds(val int) {
    if val <= 0 {
        glog.Error("duration seconds is not a positive number:", val)
        return
    }
    req.durationSeconds = val
}

func (req *GetSessionTokenRequest) SetAcl(val string) {
    req.acl = val
}

func (req *GetSessionTokenRequest) GetUri() string {
    return bce.URI_PREFIX + REQUEST_URI
}

func (req *GetSessionTokenRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BceRequest.BuildHttpRequest()
    httpReq.SetUri(req.GetUri())
    httpReq.SetMethod(http.POST)
    httpReq.AddParam("durationSeconds", strconv.Itoa(req.durationSeconds))

    if len(req.acl) > 0 {
        httpReq.AddHeader(http.CONTENT_TYPE, "application/json; charset=utf-8")
        httpReq.SetBody(http.NewBodyStreamFromString(req.acl))
    }

    return httpReq
}

func NewGetSessionTokenRequest(dur int, acl string) *GetSessionTokenRequest {
    // If the duration seconds is not a positive, use the default value
    if dur <= 0 {
        dur = DEFAULT_DURATION_SECONDS
    }
    return &GetSessionTokenRequest{durationSeconds: dur, acl: acl}
}

