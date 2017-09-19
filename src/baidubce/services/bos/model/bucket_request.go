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

// Define common bucket request structure for reuse
package model

import (
    "baidubce/common"
    "baidubce/model"
    "baidubce/http"
)

type BucketRequest struct {
    model.BceRequest // Derived from BceRequest
    bucketName string
}

func (req *BucketRequest) GetUri() string {
    return common.URI_PREFIX + req.bucketName
}

func (req *BucketRequest) BuildHttpRequest() *http.Request {
    httpReq := req.BceRequest.BuildHttpRequest()
    httpReq.SetUri(req.GetUri())
    return httpReq
}

