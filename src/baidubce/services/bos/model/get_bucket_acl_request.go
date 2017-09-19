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

// Get bucket acl request definition
package model

import (
    "baidubce/http"
)

type GetBucketAclRequest struct {
    BucketRequest
}

func (req *GetBucketAclRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.GET)
    httpReq.AddParam("acl", "")
    return httpReq
}

func NewGetBucketAclRequest(name string) *GetBucketAclRequest {
    return &GetBucketAclRequest{BucketRequest{bucketName: name}}
}

