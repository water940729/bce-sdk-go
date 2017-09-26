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

// delete_bucket_lifecycle_request.go - the delete bucket lifecycle request definition

package model

import (
    "baidubce/http"
)

type DeleteBucketLifecycleRequest struct {
    BucketRequest
}

func (req *DeleteBucketLifecycleRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BucketRequest.BuildHttpRequest()
    httpReq.SetMethod(http.DELETE)
    httpReq.AddParam("lifecycle", "")
    return httpReq
}

func NewDeleteBucketLifecycleRequest(name string) *DeleteBucketLifecycleRequest {
    return &DeleteBucketLifecycleRequest{BucketRequest{bucketName: name}}
}
